package usecase

import (
	"take-home-test/app/repositories"
	"take-home-test/pkg/config"
)

type Main struct {
	User     UserInterface
	Field    FieldInterface
	Booking  BookingInterface
	Payment  PaymentInterface
	Auth     AuthInterface
	Validate ValidateInterface
}

type usecase struct {
	Options Options
}

type Options struct {
	Repository *repositories.Main
	Config     *config.Config
}

func Init(opts Options) *Main {
	uc := &usecase{opts}

	m := &Main{
		User:     (*userUsecase)(uc),
		Field:    (*fieldUsecase)(uc),
		Booking:  (*bookingUsecase)(uc),
		Payment:  (*paymentUsecase)(uc),
		Auth:     (*authUsecase)(uc),
		Validate: (*validateUsecase)(uc),
	}

	return m
}
