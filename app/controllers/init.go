package controllers

import (
	"take-home-test/pkg/config"
	"take-home-test/app/usecases"
)

type Main struct {
	Auth    AuthInterface
	User    UserInterface
	Field   FieldInterface
	Booking BookingInterface
	Payment PaymentInterface
}

type controller struct {
	Options Options
}

type Options struct {
	Config   *config.Config
	UseCases *usecase.Main
}

func Init(opts Options) *Main {
	ctrl := &controller{opts}

	m := &Main{
		Auth:    (*authController)(ctrl),
		User:    (*userController)(ctrl),
		Field:   (*fieldController)(ctrl),
		Booking: (*bookingController)(ctrl),
		Payment: (*paymentController)(ctrl),
	}

	return m
}