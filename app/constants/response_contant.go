package constants

const (
	// Success messages
	SUCCESS_RESPONSE_MESSAGE        = "Success"
	CREATED_RESPONSE_MESSAGE        = "Created successfully"
	UPDATED_RESPONSE_MESSAGE        = "Updated successfully"
	DELETED_RESPONSE_MESSAGE        = "Deleted successfully"
	REGISTER_SUCCESS_MESSAGE        = "User registered successfully"
	LOGIN_SUCCESS_MESSAGE           = "Login successful"
	LOGOUT_SUCCESS_MESSAGE          = "Logout successful"
	BOOKING_SUCCESS_MESSAGE         = "Booking created successfully"
	PAYMENT_SUCCESS_MESSAGE         = "Payment processed successfully"

	// Error messages
	ERR_INVALID_ID                  = "invalid id"
	ERR_INVALID_UUID                = "Invalid UUID"
	GORM_ERR_NOT_FOUND              = "record not found"
	ERR_VALIDATION_FAILED           = "Validation failed"
	ERR_UNAUTHORIZED_ACCESS         = "Unauthorized access"
	ERR_FORBIDDEN_ACCESS            = "Forbidden access"
	ERR_RESOURCE_NOT_FOUND          = "Resource not found"
	ERR_DUPLICATE_ENTRY             = "Duplicate entry"
	ERR_INTERNAL_SERVER             = "Internal server error"

	// HTTP Status related
	STATUS_SUCCESS      = "success"
	STATUS_ERROR        = "error"
	STATUS_UNAUTHORIZED = "unauthorized"
	STATUS_FORBIDDEN    = "forbidden"
	STATUS_NOT_FOUND    = "not found"
	STATUS_BAD_REQUEST  = "bad request"
)

// Response codes
const (
	CODE_SUCCESS        = 1000
	CODE_CREATED        = 1001
	CODE_UPDATED        = 1002
	CODE_DELETED        = 1003
	
	CODE_BAD_REQUEST    = 2000
	CODE_UNAUTHORIZED   = 2001
	CODE_FORBIDDEN      = 2002
	CODE_NOT_FOUND      = 2003
	CODE_VALIDATION_ERR = 2004
	CODE_DUPLICATE_ERR  = 2005
	
	CODE_INTERNAL_ERROR = 3000
	CODE_DATABASE_ERROR = 3001
)