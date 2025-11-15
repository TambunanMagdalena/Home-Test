package models

import (
	"time"

	"github.com/google/uuid"
)

type Field struct {
	ID           uuid.UUID `json:"id" gorm:"type:uuid;primary_key;default:gen_random_uuid()"`
	Name         string    `json:"name"`
	PricePerHour int       `json:"price_per_hour"`
	Location     string    `json:"location"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (Field) TableName() string {
	return "fields"
}

type FieldResponse struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	PricePerHour int       `json:"price_per_hour"`
	Location     string    `json:"location"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type CreateFieldRequest struct {
	Name         string `json:"name" validate:"required"`
	PricePerHour int    `json:"price_per_hour" validate:"required,min=0"`
	Location     string `json:"location" validate:"required"`
}

type UpdateFieldRequest struct {
	Name         string `json:"name" validate:"required"`
	PricePerHour int    `json:"price_per_hour" validate:"required,min=0"`
	Location     string `json:"location" validate:"required"`
}
