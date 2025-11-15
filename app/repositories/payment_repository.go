package repositories

import (
	"context"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror"

	"gorm.io/gorm"
)

type paymentRepository struct {
	Options Options
}

// ✅ UPDATED INTERFACE - INCLUDES UpdatePaymentMethod
type PaymentInterface interface {
	CreatePayment(ctx context.Context, payment models.Payment) (models.Payment, error)
	GetPaymentByBookingID(ctx context.Context, bookingID string) (models.Payment, error)
	GetPaymentByID(ctx context.Context, id string) (models.Payment, error)
	UpdatePaymentStatus(ctx context.Context, id string, status string) error
	UpdatePaymentMethod(ctx context.Context, id string, paymentMethod string) error // ✅ ADDED
	ProcessPayment(ctx context.Context, id string) error
}

func (r *paymentRepository) CreatePayment(ctx context.Context, payment models.Payment) (models.Payment, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&payment).Error
	return payment, err
}

func (r *paymentRepository) GetPaymentByBookingID(ctx context.Context, bookingID string) (models.Payment, error) {
	var payment models.Payment
	err := r.Options.Postgres.WithContext(ctx).Where("booking_id = ?", bookingID).First(&payment).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return payment, customerror.NewNotFoundErrorf(constants.ErrPaymentNotFoundByBooking, bookingID)
		}
		return payment, customerror.NewInternalServiceError(err.Error())
	}
	return payment, nil
}

func (r *paymentRepository) GetPaymentByID(ctx context.Context, id string) (models.Payment, error) {
	var payment models.Payment
	err := r.Options.Postgres.WithContext(ctx).Where("id = ?", id).First(&payment).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return payment, customerror.NewNotFoundErrorf(constants.ErrPaymentNotFound, id)
		}
		return payment, customerror.NewInternalServiceError(err.Error())
	}
	return payment, nil
}

func (r *paymentRepository) UpdatePaymentStatus(ctx context.Context, id string, status string) error {
	result := r.Options.Postgres.WithContext(ctx).Model(&models.Payment{}).
		Where("id = ?", id).
		Update("status", status)

	if result.Error != nil {
		return customerror.NewInternalServiceError(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return customerror.NewNotFoundErrorf(constants.ErrPaymentNotFound, id)
	}

	return nil
}

func (r *paymentRepository) UpdatePaymentMethod(ctx context.Context, id string, paymentMethod string) error {
	result := r.Options.Postgres.WithContext(ctx).Model(&models.Payment{}).
		Where("id = ?", id).
		Update("payment_method", paymentMethod)

	if result.Error != nil {
		return customerror.NewInternalServiceError(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return customerror.NewNotFoundErrorf(constants.ErrPaymentNotFound, id)
	}

	return nil
}

func (r *paymentRepository) ProcessPayment(ctx context.Context, id string) error {
	result := r.Options.Postgres.WithContext(ctx).Model(&models.Payment{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"status":  constants.PAYMENT_STATUS_SUCCESS,
			"paid_at": gorm.Expr("CURRENT_TIMESTAMP"),
		})

	if result.Error != nil {
		return customerror.NewInternalServiceError(result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return customerror.NewNotFoundErrorf(constants.ErrPaymentNotFound, id)
	}

	return nil
}
