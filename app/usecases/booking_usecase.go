package usecase

import (
	"context"
	"fmt"
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/app/models"
	"time"
)

type bookingUsecase usecase

type BookingInterface interface {
	CreateBooking(ctx context.Context, userID string, req models.CreateBookingRequest) (*models.BookingResponse, error)
	GetBookingByID(ctx context.Context, id string) (*models.BookingResponse, error)
	GetUserBookings(ctx context.Context, userID string) ([]models.BookingResponse, error)
}

func (u *bookingUsecase) CreateBooking(ctx context.Context, userID string, req models.CreateBookingRequest) (*models.BookingResponse, error) {
	// Check if field exists
	field, err := u.Options.Repository.Field.GetFieldByID(ctx, req.FieldID.String())
	if err != nil {
		return nil, fmt.Errorf(constants.ErrFieldNotFound, req.FieldID.String())
	}

	hasOverlap, err := u.Options.Repository.Booking.CheckTimeOverlap(ctx, req.FieldID.String(), req.StartTime, req.EndTime)
	if err != nil {
		return nil, err
	}
	if hasOverlap {
		return nil, fmt.Errorf(constants.ErrTimeSlotOverlap)
	}

	if !req.EndTime.After(req.StartTime) {
		return nil, fmt.Errorf(constants.ErrInvalidTimeRange)
	}

	if req.StartTime.Before(time.Now()) {
		return nil, fmt.Errorf(constants.ErrBookingInPast)
	}

	minDuration := time.Hour
	if req.EndTime.Sub(req.StartTime) < minDuration {
		return nil, fmt.Errorf(constants.ErrMinimumDuration)
	}

	booking := models.Booking{
		UserID:    helpers.ParseUUID(userID),
		FieldID:   req.FieldID,
		StartTime: req.StartTime,
		EndTime:   req.EndTime,
		Status:    constants.BOOKING_STATUS_PENDING,
	}

	createdBooking, err := u.Options.Repository.Booking.CreateBooking(ctx, booking)
	if err != nil {
		return nil, err
	}

	duration := req.EndTime.Sub(req.StartTime)
	hours := int(duration.Hours())
	if hours == 0 {
		hours = 1
	}
	amount := hours * field.PricePerHour

	payment := models.Payment{
		BookingID:     createdBooking.ID,
		Amount:        amount,
		Status:        constants.PAYMENT_STATUS_PENDING,
		PaymentMethod: "",
	}

	if _, err := u.Options.Repository.Payment.CreatePayment(ctx, payment); err != nil {
	}

	bookingResponse := &models.BookingResponse{
		ID:        createdBooking.ID,
		UserID:    createdBooking.UserID,
		FieldID:   createdBooking.FieldID,
		StartTime: createdBooking.StartTime,
		EndTime:   createdBooking.EndTime,
		Status:    createdBooking.Status,
		CreatedAt: createdBooking.CreatedAt,
	}

	return bookingResponse, nil
}

func (u *bookingUsecase) GetBookingByID(ctx context.Context, id string) (*models.BookingResponse, error) {
	booking, err := u.Options.Repository.Booking.GetBookingByID(ctx, id)
	if err != nil {
		return nil, err
	}

	bookingResponse := &models.BookingResponse{
		ID:        booking.ID,
		UserID:    booking.UserID,
		FieldID:   booking.FieldID,
		StartTime: booking.StartTime,
		EndTime:   booking.EndTime,
		Status:    booking.Status,
		CreatedAt: booking.CreatedAt,
	}

	return bookingResponse, nil
}

func (u *bookingUsecase) GetUserBookings(ctx context.Context, userID string) ([]models.BookingResponse, error) {
	bookings, err := u.Options.Repository.Booking.GetBookingsByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var bookingResponses []models.BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, models.BookingResponse{
			ID:        booking.ID,
			UserID:    booking.UserID,
			FieldID:   booking.FieldID,
			StartTime: booking.StartTime,
			EndTime:   booking.EndTime,
			Status:    booking.Status,
			CreatedAt: booking.CreatedAt,
		})
	}

	return bookingResponses, nil
}
