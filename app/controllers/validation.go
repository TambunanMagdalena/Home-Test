package controllers

import (
	"github.com/ezartsh/validet"
	"github.com/gofiber/fiber/v2"
)

// IsValidBookingTime validates booking time constraints

// IsValidPrice validates price field
func IsValidPrice(field string, path validet.PathKey, val string, look validet.Lookup) error {
	if field == "price_per_hour" {
		// You can add price validation logic here
		// For example, minimum price validation
		// if price < 0 {
		//     return fiber.NewError(fiber.StatusBadRequest, "Price must be positive")
		// }
	}
	return nil
}

// IsValidEmail validates email format and uniqueness
func IsValidEmail(field string, path validet.PathKey, val string, look validet.Lookup) error {
	if field == "email" {
		// Email format validation is already handled by validet
		// You can add additional business logic here if needed
	}
	return nil
}

// IsValidRole validates user role
func IsValidRole(field string, path validet.PathKey, val string, look validet.Lookup) error {
	if field == "role" {
		if val != "user" && val != "admin" {
			return fiber.NewError(fiber.StatusBadRequest, "Role must be either 'user' or 'admin'")
		}
	}
	return nil
}

// IsValidBookingStatus validates booking status
func IsValidBookingStatus(field string, path validet.PathKey, val string, look validet.Lookup) error {
	if field == "status" {
		validStatuses := []string{"pending", "paid", "canceled"}
		isValid := false
		for _, status := range validStatuses {
			if val == status {
				isValid = true
				break
			}
		}
		if !isValid {
			return fiber.NewError(fiber.StatusBadRequest, "Status must be one of: pending, paid, canceled")
		}
	}
	return nil
}

// IsValidPaymentStatus validates payment status
func IsValidPaymentStatus(field string, path validet.PathKey, val string, look validet.Lookup) error {
	if field == "status" {
		validStatuses := []string{"pending", "success", "failed"}
		isValid := false
		for _, status := range validStatuses {
			if val == status {
				isValid = true
				break
			}
		}
		if !isValid {
			return fiber.NewError(fiber.StatusBadRequest, "Status must be one of: pending, success, failed")
		}
	}
	return nil
}
