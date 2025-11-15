package helpers

import (
	"take-home-test/app/models"

	"github.com/gofiber/fiber/v2"
)

// HTTP status constants untuk backup
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusInternalServerError = 500
)

func ResponseWrapper(c *fiber.Ctx, statusCode int, response interface{}) error {
	// Set CORS headers
	c.Set("Access-Control-Allow-Origin", "*")
	c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
	c.Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
	
	return c.Status(statusCode).JSON(response)
}

func StandardResponse(c *fiber.Ctx, statusCode int, message interface{}, data interface{}, pagination *models.Pagination) error {
	switch {
	case pagination == nil:
		return ResponseWrapper(c, statusCode, models.Response{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
		})
	default:
		return ResponseWrapper(c, statusCode, models.ResponseWithPaginate{
			StatusCode: statusCode,
			Message:    message,
			Data:       data,
			Pagination: pagination,
		})
	}
}

func Response(c *fiber.Ctx, statusCode int, message []string) error {
	return ResponseWrapper(c, statusCode, models.BasicResponse{
		StatusCode: statusCode,
		Message:    message,
	})
}

// Convenience response functions - menggunakan numeric values
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return StandardResponse(c, StatusOK, "Success", data, nil)
}

func CreatedResponse(c *fiber.Ctx, data interface{}) error {
	return StandardResponse(c, StatusCreated, "Created successfully", data, nil)
}

func BadRequestResponse(c *fiber.Ctx, message string) error {
	return Response(c, StatusBadRequest, []string{message})
}

func UnauthorizedResponse(c *fiber.Ctx, message string) error {
	return Response(c, StatusUnauthorized, []string{message})
}

func ForbiddenResponse(c *fiber.Ctx, message string) error {
	return Response(c, StatusForbidden, []string{message})
}

func NotFoundResponse(c *fiber.Ctx, message string) error {
	return Response(c, StatusNotFound, []string{message})
}

func InternalServerErrorResponse(c *fiber.Ctx, message string) error {
	return Response(c, StatusInternalServerError, []string{message})
}

func ValidationErrorResponse(c *fiber.Ctx, messages []string) error {
	return Response(c, StatusBadRequest, messages)
}

// Response with pagination
func SuccessResponseWithPagination(c *fiber.Ctx, data interface{}, pagination *models.Pagination) error {
	return StandardResponse(c, StatusOK, "Success", data, pagination)
}