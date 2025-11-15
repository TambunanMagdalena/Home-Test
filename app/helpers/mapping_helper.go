package helpers

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// StringToInt converts a string to an integer.
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func ToSliceString(key string, source map[string]interface{}) map[string]interface{} {
	if val, ok := source[key]; ok && val != nil && reflect.TypeOf(val).Kind() != reflect.Slice {
		strVal := ""
		switch v := val.(type) {
		case string:
			strVal = v
		case int:
			strVal = strconv.Itoa(v)
		default:
			fmt.Println("Cannot convert to string.")
		}
		source[key] = []interface{}{strVal}
	} else {
		source[key] = []interface{}{}
	}

	return source
}

func ToNumericInt(key string, source map[string]interface{}) map[string]interface{} {
	if val, ok := source[key]; ok {
		if val != nil && val == "0" {
			intConv, _ := strconv.Atoi(val.(string))
			source[key] = intConv
		}
	}

	return source
}

// UUID Helper functions
// ParseUUID converts string to UUID, returns uuid.Nil if parsing fails
func ParseUUID(id string) uuid.UUID {
	parsedUUID, err := uuid.Parse(id)
	if err != nil {
		// Return zero UUID if parsing fails
		return uuid.Nil
	}
	return parsedUUID
}

// IsValidUUID checks if string is a valid UUID
func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// NewUUID generates a new UUID
func NewUUID() uuid.UUID {
	return uuid.New()
}

// UUIDToString converts UUID to string
func UUIDToString(id uuid.UUID) string {
	return id.String()
}

// Fiber Context Helper functions
// GetUserIDFromContext gets user ID from Fiber context
func GetUserIDFromContext(c *fiber.Ctx) string {
	if userID, ok := c.Locals("userID").(string); ok {
		return userID
	}
	return ""
}

// GetUserRoleFromContext gets user role from Fiber context
func GetUserRoleFromContext(c *fiber.Ctx) string {
	if role, ok := c.Locals("role").(string); ok {
		return role
	}
	return ""
}

// Parse query parameters with defaults
func ParseQueryInt(c *fiber.Ctx, key string, defaultValue int) int {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}

	if intValue, err := strconv.Atoi(value); err == nil {
		return intValue
	}
	return defaultValue
}

func ParseQueryString(c *fiber.Ctx, key string, defaultValue string) string {
	value := c.Query(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetUserUUIDFromContext gets user ID as UUID from Fiber context
func GetUserUUIDFromContext(c *fiber.Ctx) uuid.UUID {
	if userID, ok := c.Locals("userID").(string); ok {
		return ParseUUID(userID)
	}
	return uuid.Nil
}
