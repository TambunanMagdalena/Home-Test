package models

import (
	"time"

	"github.com/google/uuid"
)

type Booking struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	UserID    uuid.UUID `json:"user_id"`
	FieldID   uuid.UUID `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status" gorm:"default:'pending'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Booking) TableName() string {
	return "bookings"
}

type BookingResponse struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	FieldID   uuid.UUID `json:"field_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateBookingRequest struct {
	FieldID   uuid.UUID `json:"field_id" validate:"required"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}
