package models

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	ID            uuid.UUID  `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	BookingID     uuid.UUID  `json:"booking_id"`
	Amount        int        `json:"amount"`
	Status        string     `json:"status" gorm:"default:'pending'"`
	PaymentMethod string     `json:"payment_method"`
	PaidAt        *time.Time `json:"paid_at"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"` // ✅ Add this for better tracking

}

func (Payment) TableName() string {
	return "payments"
}

type PaymentResponse struct {
	ID            uuid.UUID  `json:"id"`
	BookingID     uuid.UUID  `json:"booking_id"`
	Amount        int        `json:"amount"`
	Status        string     `json:"status"`
	PaymentMethod string     `json:"payment_method"`
	PaidAt        *time.Time `json:"paid_at,omitempty"`
	CreatedAt     time.Time  `json:"created_at"`
}

type PaymentTransactionResponse struct {
	PaymentID     uuid.UUID `json:"payment_id"`
	Token         string    `json:"token"`
	RedirectURL   string    `json:"redirect_url"`
	TransactionID string    `json:"transaction_id"`
	Amount        int       `json:"amount"`
}

type CreatePaymentRequest struct {
	BookingID     uuid.UUID `json:"booking_id" validate:"required"`
	PaymentMethod string    `json:"payment_method" validate:"required"`
}
