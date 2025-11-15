package constants

const (
	// User roles
	ROLE_ADMIN = "admin"
	ROLE_USER  = "user"

	// Booking statuses
	BOOKING_STATUS_PENDING   = "pending"
	BOOKING_STATUS_PAID      = "paid"
	BOOKING_STATUS_CANCELED  = "canceled"
	BOOKING_STATUS_CONFIRMED = "confirmed"

	// Payment statuses
	PAYMENT_STATUS_PENDING = "pending"
	PAYMENT_STATUS_SUCCESS = "success"
	PAYMENT_STATUS_FAILED  = "failed"

	// Payment methods
	PAYMENT_METHOD_CASH        = "cash"
	PAYMENT_METHOD_TRANSFER    = "transfer"
	PAYMENT_METHOD_CREDIT_CARD = "credit_card"
	PAYMENT_METHOD_DEBIT_CARD  = "debit_card"

	// Time formats
	TIME_FORMAT_RFC3339 = "2006-01-02T15:04:05Z"
	TIME_FORMAT_ISO8601 = "2006-01-02T15:04:05-07:00"
	TIME_FORMAT_SIMPLE  = "2006-01-02 15:04:05"

	// Database constraints
	UNIQUE_CONSTRAINT_USER_EMAIL = "users_email_key"
	UNIQUE_CONSTRAINT_FIELD_NAME = "fields_name_key"

	// Field types (if needed for future expansion)
	FIELD_TYPE_FUTSAL     = "futsal"
	FIELD_TYPE_BASKETBALL = "basketball"
	FIELD_TYPE_TENNIS     = "tennis"
	FIELD_TYPE_BADMINTON  = "badminton"
)

var (
	// Valid user roles
	ValidUserRoles = []string{ROLE_USER, ROLE_ADMIN}

	// Valid booking statuses
	ValidBookingStatuses = []string{
		BOOKING_STATUS_PENDING,
		BOOKING_STATUS_PAID,
		BOOKING_STATUS_CANCELED,
		BOOKING_STATUS_CONFIRMED,
	}

	// Valid payment statuses
	ValidPaymentStatuses = []string{
		PAYMENT_STATUS_PENDING,
		PAYMENT_STATUS_SUCCESS,
		PAYMENT_STATUS_FAILED,
	}

	// Valid payment methods
	ValidPaymentMethods = []string{
		PAYMENT_METHOD_CASH,
		PAYMENT_METHOD_TRANSFER,
		PAYMENT_METHOD_CREDIT_CARD,
		PAYMENT_METHOD_DEBIT_CARD,
	}

	// Days of week (for potential scheduling features)
	ArrayDays = []string{
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}

	// Field types
	ValidFieldTypes = []string{
		FIELD_TYPE_FUTSAL,
		FIELD_TYPE_BASKETBALL,
		FIELD_TYPE_TENNIS,
		FIELD_TYPE_BADMINTON,
	}
)
