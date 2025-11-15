package repositories

import (
	"context"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror" // Ganti dari customerrors menjadi customerror
	"time"

	"gorm.io/gorm"
)

type bookingRepository struct {
	Options Options
}

type BookingInterface interface {
	CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error)
	GetBookingByID(ctx context.Context, id string) (models.Booking, error)
	GetBookingsByUserID(ctx context.Context, userID string) ([]models.Booking, error)
	CheckTimeOverlap(ctx context.Context, fieldID string, startTime, endTime time.Time) (bool, error)
	UpdateBookingStatus(ctx context.Context, id string, status string) error
}

func (r *bookingRepository) CreateBooking(ctx context.Context, booking models.Booking) (models.Booking, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&booking).Error
	return booking, err
}

func (r *bookingRepository) GetBookingByID(ctx context.Context, id string) (models.Booking, error) {
	var booking models.Booking
	err := r.Options.Postgres.WithContext(ctx).Where("id = ?", id).First(&booking).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return booking, customerror.NewNotFoundErrorf(constants.ErrBookingNotFound, id) // Ganti menjadi customerror.
		}
		return booking, customerror.NewInternalServiceError(err.Error()) // Ganti menjadi customerror.
	}
	return booking, nil
}

func (r *bookingRepository) GetBookingsByUserID(ctx context.Context, userID string) ([]models.Booking, error) {
	var bookings []models.Booking
	err := r.Options.Postgres.WithContext(ctx).Where("user_id = ?", userID).Find(&bookings).Error
	return bookings, err
}

func (r *bookingRepository) CheckTimeOverlap(ctx context.Context, fieldID string, startTime, endTime time.Time) (bool, error) {
	var count int64

	err := r.Options.Postgres.WithContext(ctx).Model(&models.Booking{}).
		Where("field_id = ? AND status != ?", fieldID, constants.BOOKING_STATUS_CANCELED).
		Where("(start_time < ? AND end_time > ?) OR (start_time < ? AND end_time > ?) OR (start_time >= ? AND end_time <= ?)",
			endTime, startTime,
			startTime, endTime,
			startTime, endTime).
		Count(&count).Error

	if err != nil {
		return false, customerror.NewInternalServiceError(err.Error()) // Ganti menjadi customerror.
	}

	return count > 0, nil
}

func (r *bookingRepository) UpdateBookingStatus(ctx context.Context, id string, status string) error {
	result := r.Options.Postgres.WithContext(ctx).Model(&models.Booking{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return customerror.NewInternalServiceError(result.Error.Error()) // Ganti menjadi customerror.
	}

	if result.RowsAffected == 0 {
		return customerror.NewNotFoundErrorf(constants.ErrBookingNotFound, id) // Ganti menjadi customerror.
	}

	return nil
}
