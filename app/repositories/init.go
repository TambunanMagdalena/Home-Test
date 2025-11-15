package repositories

import (
	"take-home-test/pkg/config"

	"gorm.io/gorm"
)

type Main struct {
	User    UserInterface
	Field   FieldInterface
	Booking BookingInterface
	Payment PaymentInterface
}

type repository struct {
	Options Options
}

type Options struct {
	Postgres *gorm.DB
	Config   *config.Config
}

func Init(opts Options) *Main {
	repo := &repository{opts}

	m := &Main{
		User:    (*userRepository)(repo),
		Field:   (*fieldRepository)(repo),
		Booking: (*bookingRepository)(repo),
		Payment: (*paymentRepository)(repo),
	}

	return m
}
