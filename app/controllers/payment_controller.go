package controllers

import (
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type paymentController struct {
	Options Options
}

type PaymentInterface interface {
	ProcessPayment(ctx *fiber.Ctx) error
	GetPaymentByBookingID(ctx *fiber.Ctx) error
	CreatePaymentTransaction(ctx *fiber.Ctx) error
	HandlePaymentNotification(ctx *fiber.Ctx) error
}

// CreatePaymentTransaction godoc
// @Summary Create real payment transaction (Midtrans)
// @Description Create Midtrans payment transaction for a booking
// @Tags Payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param booking_id path string true "Booking ID"
// @Success 200 {object} models.BasicResponse{data=models.PaymentTransactionResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 403 {object} models.BasicResponse
// @Router /payments/{booking_id}/transaction [post]
func (c *paymentController) CreatePaymentTransaction(ctx *fiber.Ctx) error {
	bookingID := ctx.Params("booking_id")

	if !helpers.IsValidUUID(bookingID) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	userID := helpers.GetUserIDFromContext(ctx)
	userRole := helpers.GetUserRoleFromContext(ctx)

	// Check if user owns the booking or is admin
	booking, err := c.Options.UseCases.Booking.GetBookingByID(ctx.Context(), bookingID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	if userRole != constants.ROLE_ADMIN && booking.UserID.String() != userID {
		return helpers.ForbiddenResponse(ctx, constants.ErrUnauthorizedAccess)
	}

	transaction, err := c.Options.UseCases.Payment.CreatePaymentTransaction(ctx.Context(), bookingID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, transaction)
}

// HandlePaymentNotification godoc
// @Summary Handle payment notification webhook
// @Description Webhook endpoint for Midtrans payment notifications
// @Tags Payments
// @Accept json
// @Produce json
// @Param payload body map[string]interface{} true "Midtrans notification payload"
// @Success 200 {object} models.BasicResponse
// @Router /payments/notification [post]
func (c *paymentController) HandlePaymentNotification(ctx *fiber.Ctx) error {
	var payload map[string]interface{}

	if err := ctx.BodyParser(&payload); err != nil {
		return helpers.BadRequestResponse(ctx, "Invalid notification payload")
	}

	err := c.Options.UseCases.Payment.HandlePaymentNotification(ctx.Context(), payload)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, nil)
}

// ProcessPayment godoc
// @Summary Process payment (Mock)
// @Description Process mock payment for a booking
// @Tags Payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreatePaymentRequest true "Payment data"
// @Success 200 {object} models.BasicResponse{data=models.PaymentResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 403 {object} models.BasicResponse
// @Router /payments [post]
func (c *paymentController) ProcessPayment(ctx *fiber.Ctx) error {
	var (
		reqBody models.CreatePaymentRequest
		resBody *models.PaymentResponse
		err     error
	)

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}

	if reqBody.BookingID == uuid.Nil {
		return helpers.BadRequestResponse(ctx, "Booking ID is required")
	}

	if reqBody.PaymentMethod == "" {
		return helpers.BadRequestResponse(ctx, "Payment method is required")
	}

	if !helpers.Contains(constants.ValidPaymentMethods, reqBody.PaymentMethod) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidPaymentMethod)
	}

	userID := helpers.GetUserIDFromContext(ctx)
	userRole := helpers.GetUserRoleFromContext(ctx)

	booking, err := c.Options.UseCases.Booking.GetBookingByID(ctx.Context(), reqBody.BookingID.String())
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	if userRole != constants.ROLE_ADMIN && booking.UserID.String() != userID {
		return helpers.ForbiddenResponse(ctx, constants.ErrUnauthorizedAccess)
	}

	resBody, err = c.Options.UseCases.Payment.ProcessPayment(ctx.Context(), reqBody.BookingID.String(), reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, resBody)
}

// GetPaymentByBookingID godoc
// @Summary Get payment by booking ID
// @Description Get payment details for a specific booking
// @Tags Payments
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param booking_id path string true "Booking ID"
// @Success 200 {object} models.BasicResponse{data=models.PaymentResponse}
// @Failure 403 {object} models.BasicResponse
// @Failure 404 {object} models.BasicResponse
// @Router /payments/{booking_id} [get]
func (c *paymentController) GetPaymentByBookingID(ctx *fiber.Ctx) error {
	bookingID := ctx.Params("booking_id")

	if !helpers.IsValidUUID(bookingID) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	userID := helpers.GetUserIDFromContext(ctx)
	userRole := helpers.GetUserRoleFromContext(ctx)

	booking, err := c.Options.UseCases.Booking.GetBookingByID(ctx.Context(), bookingID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	if userRole != constants.ROLE_ADMIN && booking.UserID.String() != userID {
		return helpers.ForbiddenResponse(ctx, constants.ErrUnauthorizedAccess)
	}

	payment, err := c.Options.UseCases.Payment.GetPaymentByBookingID(ctx.Context(), bookingID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, payment)
}
