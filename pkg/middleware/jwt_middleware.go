package middleware

import (
	"strings"
	"take-home-test/pkg/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

// JWTMiddleware validates JWT token
func JWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Authorization header is required",
		})
	}

	// Extract token from "Bearer <token>"
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Invalid authorization format",
		})
	}

	tokenString := parts[1]

	// Create config instance
	cfg := config.NewConfig()

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid signing method")
		}
		return []byte(cfg.GetJWTSecret()), nil
	})

	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status_code": fiber.StatusUnauthorized,
			"message":     "Invalid or expired token",
		})
	}

	// Extract claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// Set user info in context
		c.Locals("userID", claims["user_id"])
		c.Locals("email", claims["email"])
		c.Locals("role", claims["role"])
	}

	return c.Next()
}

// AdminMiddleware checks if user has admin role
func AdminMiddleware(c *fiber.Ctx) error {
	userRole, ok := c.Locals("role").(string)
	if !ok || userRole != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status_code": fiber.StatusForbidden,
			"message":     "Access denied. Admin role required",
		})
	}
	return c.Next()
}