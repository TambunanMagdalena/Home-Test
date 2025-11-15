package constants

var (
	// User errors
	ErrUserNotFound     = `User with id '%s' not found`
	ErrUserNotFoundByID = `User with id '%s' not found`

	// Field errors
	ErrFieldNotFound     = `Field with id '%s' not found`
	ErrFieldNotFoundByID = `Field with id '%s' not found`

	// Booking errors
	ErrBookingNotFound     = `Booking with id '%s' not found`
	ErrBookingNotFoundByID = `Booking with id '%s' not found`

	// Payment errors
	ErrPaymentNotFound          = `Payment with id '%s' not found`
	ErrPaymentNotFoundByID      = `Payment with id '%s' not found`
	ErrPaymentNotFoundByBooking = `Payment for booking id '%s' not found`
)

const (
	// Auth errors
	ErrDuplicateEmail     = "Duplicate email! This email is already registered."
	ErrInvalidCredentials = "Invalid email or password"
	ErrInvalidToken       = "Invalid or expired token"
	ErrMissingToken       = "Authorization token is required"
	ErrInvalidTokenFormat = "Invalid authorization token format"

	// User errors
	ErrUnauthorizedAccess  = "Unauthorized access"
	ErrAdminAccessRequired = "Admin access required"

	// Field errors
	ErrDuplicateFieldName = "Field with this name already exists"
	ErrFieldRequired      = "Field %s is required"
	ErrInvalidPrice       = "Price per hour must be greater than 0"

	// Booking errors
	ErrTimeSlotOverlap   = "Time slot is already booked for this field"
	ErrInvalidTimeRange  = "End time must be after start time"
	ErrBookingInPast     = "Booking cannot be in the past"
	ErrMinimumDuration   = "Booking duration must be at least 1 hour"
	ErrFieldNotAvailable = "Field is not available for booking"

	// Payment errors
	ErrPaymentAlreadyProcessed = "Payment has already been processed"
	ErrInvalidPaymentMethod    = "Invalid payment method"
	ErrPaymentFailed           = "Payment processing failed"

	// Validation errors
	ErrInvalidUUID     = "Invalid UUID format"
	ErrInvalidEmail    = "Invalid email format"
	ErrInvalidPassword = "Password must be at least 6 characters"
	ErrInvalidRole     = "Role must be either 'user' or 'admin'"

	// General errors
	ErrInternalServer = "Internal server error"
	ErrBadRequest     = "Bad request"
)
