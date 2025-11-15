package usecase

import (
	"context"
	"fmt"
	"take-home-test/app/constants"
	"time"

)

type validateUsecase usecase

type ValidateInterface interface {
	IsValidFieldID(ctx context.Context, fieldID string) error
	IsValidUserID(ctx context.Context, userID string) error
	IsValidBookingID(ctx context.Context, bookingID string) error
	IsValidRequestField(ctx context.Context, request map[string]any, validateType string) error
	IsAdminUser(ctx context.Context, userID string) error
	IsValidBookingTime(ctx context.Context, fieldID string, startTime, endTime string) error
}

func (v *validateUsecase) IsValidFieldID(ctx context.Context, fieldID string) error {
	_, err := v.Options.Repository.Field.GetFieldByID(ctx, fieldID)
	if err != nil {
		return fmt.Errorf("field with id '%s' not found", fieldID)
	}
	return nil
}

func (v *validateUsecase) IsValidUserID(ctx context.Context, userID string) error {
	_, err := v.Options.Repository.User.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user with id '%s' not found", userID)
	}
	return nil
}

func (v *validateUsecase) IsValidBookingID(ctx context.Context, bookingID string) error {
	_, err := v.Options.Repository.Booking.GetBookingByID(ctx, bookingID)
	if err != nil {
		return fmt.Errorf("booking with id '%s' not found", bookingID)
	}
	return nil
}

func (v *validateUsecase) IsValidRequestField(ctx context.Context, request map[string]any, validateType string) error {
	// Check if field with same name already exists (for create)
	if validateType == "create" {
		fieldName := fmt.Sprint(request["name"])

		// Get all fields to check for duplicates
		fields, err := v.Options.Repository.Field.GetFields(ctx)
		if err != nil {
			return fmt.Errorf("failed to validate field: %v", err)
		}

		for _, field := range fields {
			if field.Name == fieldName {
				return fmt.Errorf("field with name '%s' already exists", fieldName)
			}
		}
	}

	// For update, we might want different validation logic
	if validateType == "update" {
		fieldID := fmt.Sprint(request["id"])
		fieldName := fmt.Sprint(request["name"])

		// Check if field exists
		if _, err := v.Options.Repository.Field.GetFieldByID(ctx, fieldID); err != nil {
			return fmt.Errorf("field with id '%s' not found", fieldID)
		}

		// Check if another field with same name exists (excluding current field)
		fields, err := v.Options.Repository.Field.GetFields(ctx)
		if err != nil {
			return fmt.Errorf("failed to validate field: %v", err)
		}

		for _, field := range fields {
			if field.Name == fieldName && field.ID.String() != fieldID {
				return fmt.Errorf("field with name '%s' already exists", fieldName)
			}
		}
	}

	return nil
}

func (v *validateUsecase) IsAdminUser(ctx context.Context, userID string) error {
	user, err := v.Options.Repository.User.FindByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user with id '%s' not found", userID)
	}

	if user.Role != constants.ROLE_ADMIN {
		return fmt.Errorf("user with id '%s' is not authorized to perform this action", userID)
	}

	return nil
}

func (v *validateUsecase) IsValidBookingTime(ctx context.Context, fieldID string, startTime, endTime string) error {
	// Check if field exists
	if _, err := v.Options.Repository.Field.GetFieldByID(ctx, fieldID); err != nil {
		return fmt.Errorf(constants.ErrFieldNotFound, fieldID)
	}

	// Parse times using constants
	start, err := parseTime(startTime)
	if err != nil {
		return fmt.Errorf("invalid start time format: %v", err)
	}

	end, err := parseTime(endTime)
	if err != nil {
		return fmt.Errorf("invalid end time format: %v", err)
	}

	// Check if end time is after start time
	if !end.After(start) {
		return fmt.Errorf(constants.ErrInvalidTimeRange)
	}

	// Check for time overlap
	hasOverlap, err := v.Options.Repository.Booking.CheckTimeOverlap(ctx, fieldID, start, end)
	if err != nil {
		return fmt.Errorf("failed to check booking availability: %v", err)
	}

	if hasOverlap {
		return fmt.Errorf(constants.ErrTimeSlotOverlap)
	}

	// Check if booking is in the past
	if start.Before(time.Now()) {
		return fmt.Errorf(constants.ErrBookingInPast)
	}

	// Check minimum duration
	minDuration := time.Hour
	if end.Sub(start) < minDuration {
		return fmt.Errorf(constants.ErrMinimumDuration)
	}

	return nil
}

// ❌ HAPUS FUNCTION INI - SUDAH ADA DI ATAS
/*
// Helper function to parse string to UUID
func ParseUUID(id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		// Return zero UUID if parsing fails
		return uuid.Nil
	}
	return parsedUUID
}
*/

// ✅ INI SATU-SATUNYA parseTime FUNCTION
// Helper function to parse time - GUNAKAN CONSTANTS
func parseTime(timeStr string) (time.Time, error) {
	formats := []string{
		constants.TIME_FORMAT_RFC3339,
		constants.TIME_FORMAT_ISO8601,
		constants.TIME_FORMAT_SIMPLE,
		time.RFC3339,
	}

	for _, format := range formats {
		if t, err := time.Parse(format, timeStr); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse time: %s", timeStr)
}