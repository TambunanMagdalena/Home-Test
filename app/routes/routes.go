package routes

import (
	"take-home-test/app/controllers"
	"take-home-test/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func ConfigureRouter(app *fiber.App, controller *controllers.Main) {
	// Swagger documentation
	app.Get("/swagger/*", swagger.New(swagger.Config{
		URL:          "/swagger/doc.json",
		DeepLinking:  false,
		DocExpansion: "list",
	}))

	// API v1 group
	api := app.Group("/api")
	{
		// Public routes (no auth required)
		public := api.Group("/auth")
		{
			public.Post("/register", controller.Auth.Register)
			public.Post("/login", controller.Auth.Login)
		}

		// Public Field routes (no auth required)
		api.Get("/fields", controller.Field.GetFields)                    // Public
		api.Get("/fields/:id", controller.Field.GetFieldByID)             // Public

		// ✅ PUBLIC Payment routes (no auth required)
		api.Get("/payments/:booking_id", controller.Payment.GetPaymentByBookingID) // Public - View payment

		// Protected routes (require JWT auth)
		protected := api.Group("", middleware.JWTMiddleware)
		{
			// User routes
			users := protected.Group("/users")
			{
				users.Get("/profile", controller.User.GetProfile)
				users.Get("/:id", controller.User.GetUserByID)
			}

			// Protected Field routes (admin only)
			fields := protected.Group("/fields")
			{
				fields.Post("", controller.Field.CreateField)       // Admin only
				fields.Put("/:id", controller.Field.UpdateField)    // Admin only
				fields.Delete("/:id", controller.Field.DeleteField) // Admin only
			}

			// Booking routes
			bookings := protected.Group("/bookings")
			{
				bookings.Post("", controller.Booking.CreateBooking)
				bookings.Get("/user", controller.Booking.GetUserBookings)
				bookings.Get("/:id", controller.Booking.GetBookingByID)
			}

			// ✅ PROTECTED Payment routes (butuh auth untuk action)
			payments := protected.Group("/payments")
			{
				payments.Post("", controller.Payment.ProcessPayment)                                   // Butuh auth - Process payment
				payments.Post("/:booking_id/transaction", controller.Payment.CreatePaymentTransaction) // Butuh auth - Create transaction
			}
		}

		// Public payment notification (no auth required)
		api.Post("/payments/notification", controller.Payment.HandlePaymentNotification)
	}

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "OK",
			"message": "Service is running",
		})
	})
}