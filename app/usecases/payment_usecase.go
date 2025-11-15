package usecase

import (
	"context"
	"fmt"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"take-home-test/pkg/payment"

	"github.com/google/uuid"
)

type paymentUsecase usecase

type PaymentInterface interface {
	ProcessPayment(ctx context.Context, bookingID string, req models.CreatePaymentRequest) (*models.PaymentResponse, error)
	GetPaymentByBookingID(ctx context.Context, bookingID string) (*models.PaymentResponse, error)
	CreatePaymentTransaction(ctx context.Context, bookingID string) (*models.PaymentTransactionResponse, error)
	HandlePaymentNotification(ctx context.Context, payload map[string]interface{}) error
}

func (u *paymentUsecase) CreatePaymentTransaction(ctx context.Context, bookingID string) (*models.PaymentTransactionResponse, error) {
	fmt.Printf("🔧 Creating REAL payment transaction for booking: %s\n", bookingID)

	booking, err := u.Options.Repository.Booking.GetBookingByID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	field, err := u.Options.Repository.Field.GetFieldByID(ctx, booking.FieldID.String())
	if err != nil {
		return nil, err
	}

	user, err := u.Options.Repository.User.FindByID(ctx, booking.UserID.String())
	if err != nil {
		return nil, err
	}

	existingPayment, err := u.Options.Repository.Payment.GetPaymentByBookingID(ctx, bookingID)
	if err == nil && existingPayment.ID != uuid.Nil {
		if existingPayment.Status == constants.PAYMENT_STATUS_SUCCESS {
			return nil, fmt.Errorf(constants.ErrPaymentAlreadyProcessed)
		}
	}

	duration := booking.EndTime.Sub(booking.StartTime)
	hours := int(duration.Hours())
	if hours == 0 {
		hours = 1
	}
	amount := hours * field.PricePerHour

	paymentRecord := models.Payment{
		BookingID:     booking.ID,
		Amount:        amount,
		Status:        constants.PAYMENT_STATUS_PENDING,
		PaymentMethod: "",
	}

	createdPayment, err := u.Options.Repository.Payment.CreatePayment(ctx, paymentRecord)
	if err != nil {
		return nil, err
	}

	isProduction := u.Options.Config.ServiceEnvironment == "production"
	paymentService := payment.NewMidtransService(u.Options.Config.MidtransServerKey, isProduction)

	itemName := fmt.Sprintf("Booking %s - %s", field.Name, booking.StartTime.Format("02 Jan 2006 15:04"))
	snapResp, err := paymentService.CreateTransaction(
		booking.ID.String(),
		int64(amount),
		user.Name,
		user.Email,
		itemName,
	)
	if err != nil {
		u.Options.Repository.Payment.UpdatePaymentStatus(ctx, createdPayment.ID.String(), constants.PAYMENT_STATUS_FAILED)
		return nil, fmt.Errorf("payment gateway error: %v", err)
	}

	var transactionID string
	transactionDetails, err := paymentService.GetTransactionDetails(booking.ID.String())
	if err != nil {
		transactionID = "pending"
	} else {
		transactionID = transactionDetails.TransactionID
	}

	response := &models.PaymentTransactionResponse{
		PaymentID:     createdPayment.ID,
		Token:         snapResp.Token,
		RedirectURL:   snapResp.RedirectURL,
		TransactionID: transactionID,
		Amount:        amount,
	}

	return response, nil
}

func (u *paymentUsecase) HandlePaymentNotification(ctx context.Context, payload map[string]interface{}) error {
	orderID, ok := payload["order_id"].(string)
	if !ok {
		return fmt.Errorf("invalid notification payload: order_id missing")
	}

	isProduction := u.Options.Config.ServiceEnvironment == "production"
	paymentService := payment.NewMidtransService(u.Options.Config.MidtransServerKey, isProduction)

	notification, err := paymentService.CheckTransactionStatus(orderID)
	if err != nil {
		return err
	}

	fmt.Printf(" Notification details - OrderID: %s, Status: %s, Type: %s\n",
		notification.OrderID, notification.TransactionStatus, notification.PaymentType)

	// Determine statuses
	var paymentStatus string
	var bookingStatus string

	switch notification.TransactionStatus {
	case "capture", "settlement":
		paymentStatus = constants.PAYMENT_STATUS_SUCCESS
		bookingStatus = constants.BOOKING_STATUS_PAID
	case "deny", "cancel", "expire", "failure":
		paymentStatus = constants.PAYMENT_STATUS_FAILED
		bookingStatus = constants.BOOKING_STATUS_PENDING
	case "pending":
		paymentStatus = constants.PAYMENT_STATUS_PENDING
		bookingStatus = constants.BOOKING_STATUS_PENDING
	default:
		paymentStatus = constants.PAYMENT_STATUS_PENDING
		bookingStatus = constants.BOOKING_STATUS_PENDING
	}

	err = u.Options.Repository.Payment.UpdatePaymentStatus(ctx, notification.OrderID, paymentStatus)
	if err != nil {
		return err
	}

	if notification.PaymentType != "" {
		err = u.Options.Repository.Payment.UpdatePaymentMethod(ctx, notification.OrderID, notification.PaymentType)
		if err != nil {
		}
	}

	if paymentStatus == constants.PAYMENT_STATUS_SUCCESS {
		err = u.Options.Repository.Booking.UpdateBookingStatus(ctx, notification.OrderID, bookingStatus)
		if err != nil {
			return err
		}
	}

	return nil
}

func (u *paymentUsecase) ProcessPayment(ctx context.Context, bookingID string, req models.CreatePaymentRequest) (*models.PaymentResponse, error) {
	payment, err := u.Options.Repository.Payment.GetPaymentByBookingID(ctx, bookingID)
	if err != nil {
		return nil, err
	}
	if payment.Status == constants.PAYMENT_STATUS_SUCCESS {
		return nil, fmt.Errorf(constants.ErrPaymentAlreadyProcessed)
	}
	err = u.Options.Repository.Payment.ProcessPayment(ctx, payment.ID.String())
	if err != nil {
		return nil, err
	}

	err = u.Options.Repository.Payment.UpdatePaymentMethod(ctx, payment.ID.String(), req.PaymentMethod)
	if err != nil {
	}

	err = u.Options.Repository.Booking.UpdateBookingStatus(ctx, bookingID, constants.BOOKING_STATUS_PAID)
	if err != nil {
		return nil, err
	}

	updatedPayment, err := u.Options.Repository.Payment.GetPaymentByBookingID(ctx, bookingID)
	if err != nil {
		return nil, err
	}

	paymentResponse := &models.PaymentResponse{
		ID:            updatedPayment.ID,
		BookingID:     updatedPayment.BookingID,
		Amount:        updatedPayment.Amount,
		Status:        updatedPayment.Status,
		PaymentMethod: updatedPayment.PaymentMethod,
		PaidAt:        updatedPayment.PaidAt,
		CreatedAt:     updatedPayment.CreatedAt,
	}
	return paymentResponse, nil
}

func (u *paymentUsecase) GetPaymentByBookingID(ctx context.Context, bookingID string) (*models.PaymentResponse, error) {
	payment, err := u.Options.Repository.Payment.GetPaymentByBookingID(ctx, bookingID)
	if err != nil {
		return nil, err
	}
	paymentResponse := &models.PaymentResponse{
		ID:            payment.ID,
		BookingID:     payment.BookingID,
		Amount:        payment.Amount,
		Status:        payment.Status,
		PaymentMethod: payment.PaymentMethod,
		PaidAt:        payment.PaidAt,
		CreatedAt:     payment.CreatedAt,
	}

	return paymentResponse, nil
}
