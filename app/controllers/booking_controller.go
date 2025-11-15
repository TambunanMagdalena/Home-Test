package controllers

import (
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror"

	"github.com/gofiber/fiber/v2"
)

type bookingController struct {
	Options Options
}

type BookingInterface interface {
	CreateBooking(ctx *fiber.Ctx) error
	GetBookingByID(ctx *fiber.Ctx) error
	GetUserBookings(ctx *fiber.Ctx) error
}

// CreateBooking godoc
// @Summary Create new booking
// @Description Create a new field booking
// @Tags Bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateBookingRequest true "Booking data"
// @Success 201 {object} models.BasicResponse{data=models.BookingResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 401 {object} models.BasicResponse
// @Failure 409 {object} models.BasicResponse
// @Router /bookings [post]
func (ctrl *bookingController) CreateBooking(ctx *fiber.Ctx) error {
	var (
		reqBody models.CreateBookingRequest
		resBody *models.BookingResponse
		err     error
	)

	userID := helpers.GetUserIDFromContext(ctx)
	if userID == "" {
		return helpers.UnauthorizedResponse(ctx, constants.ErrMissingToken)
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}

	if reqBody.FieldID.String() == "" {
		return helpers.BadRequestResponse(ctx, "Field ID is required")
	}

	if err := ctrl.Options.UseCases.Validate.IsValidBookingTime(
		ctx.Context(),
		reqBody.FieldID.String(),
		reqBody.StartTime.Format(constants.TIME_FORMAT_RFC3339), // ✅ GUNAKAN CONSTANTS
		reqBody.EndTime.Format(constants.TIME_FORMAT_RFC3339),
	); err != nil {
		return helpers.BadRequestResponse(ctx, err.Error())
	}

	resBody, err = ctrl.Options.UseCases.Booking.CreateBooking(ctx.Context(), userID, reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.CreatedResponse(ctx, resBody)
}

// GetBookingByID godoc
// @Summary Get booking by ID
// @Description Get booking details by ID
// @Tags Bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Booking ID"
// @Success 200 {object} models.BasicResponse{data=models.BookingResponse}
// @Failure 403 {object} models.BasicResponse
// @Failure 404 {object} models.BasicResponse
// @Router /bookings/{id} [get]
func (ctrl *bookingController) GetBookingByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if !helpers.IsValidUUID(id) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	userID := helpers.GetUserIDFromContext(ctx)
	userRole := helpers.GetUserRoleFromContext(ctx)

	booking, err := ctrl.Options.UseCases.Booking.GetBookingByID(ctx.Context(), id)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	if userRole != constants.ROLE_ADMIN && booking.UserID.String() != userID {
		return helpers.ForbiddenResponse(ctx, constants.ErrUnauthorizedAccess)
	}

	return helpers.SuccessResponse(ctx, booking)
}

// GetUserBookings godoc
// @Summary Get user bookings
// @Description Get all bookings for the authenticated user
// @Tags Bookings
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} models.BasicResponse{data=[]models.BookingResponse}
// @Router /bookings/user [get]
func (ctrl *bookingController) GetUserBookings(ctx *fiber.Ctx) error {
	userID := helpers.GetUserIDFromContext(ctx)
	if userID == "" {
		return helpers.UnauthorizedResponse(ctx, constants.ErrMissingToken)
	}

	bookings, err := ctrl.Options.UseCases.Booking.GetUserBookings(ctx.Context(), userID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, bookings)
}
