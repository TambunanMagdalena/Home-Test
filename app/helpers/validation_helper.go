package helpers

import (
	"take-home-test/app/constants"
	"take-home-test/app/models"

	"github.com/gofiber/fiber/v2"
)

func ValidateRegisterRequest(req models.RegisterRequest) error {
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Name is required")
	}

	if req.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidEmail)
	}

	if req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidPassword)
	}

	if len(req.Password) < constants.MIN_PASSWORD_LENGTH {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidPassword)
	}

	if req.Role != "" && req.Role != constants.ROLE_USER && req.Role != constants.ROLE_ADMIN {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidRole)
	}

	return nil
}

func ValidateBookingRequest(req models.CreateBookingRequest) error {
	if req.FieldID.String() == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Field ID is required")
	}
	if req.StartTime.IsZero() {
		return fiber.NewError(fiber.StatusBadRequest, "Start time is required")
	}
	if req.EndTime.IsZero() {
		return fiber.NewError(fiber.StatusBadRequest, "End time is required")
	}
	if !req.EndTime.After(req.StartTime) {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidTimeRange)
	}
	return nil
}
func ValidateUUID(id string) error {
	if id == "" {
		return fiber.NewError(fiber.StatusBadRequest, "ID is required")
	}
	if !IsValidUUID(id) {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidUUID)
	}
	return nil
}

func ValidateFieldRequest(req models.CreateFieldRequest) error {
	if req.Name == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Field name is required")
	}
	if req.Location == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Field location is required")
	}
	if req.PricePerHour <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, constants.ErrInvalidPrice)
	}
	return nil
}

func ValidateLoginRequest(req models.LoginRequest) error {
	if req.Email == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Email is required")
	}

	if req.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Password is required")
	}

	return nil
}

func Contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
