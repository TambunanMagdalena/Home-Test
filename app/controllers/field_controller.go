package controllers

import (
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror"

	"github.com/gofiber/fiber/v2"
)

type fieldController struct {
	Options Options
}

type FieldInterface interface {
	CreateField(ctx *fiber.Ctx) error
	GetFields(ctx *fiber.Ctx) error
	GetFieldByID(ctx *fiber.Ctx) error
	UpdateField(ctx *fiber.Ctx) error
	DeleteField(ctx *fiber.Ctx) error
}

// CreateField godoc
// @Summary Create new field
// @Description Create a new sports field. ADMIN ACCESS ONLY - Regular users cannot create fields.
// @Tags Fields
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body models.CreateFieldRequest true "Field data including name, price_per_hour, and location"
// @Success 201 {object} models.BasicResponse{data=models.FieldResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 401 {object} models.BasicResponse
// @Failure 403 {object} models.BasicResponse
// @Router /fields [post]
func (ctrl *fieldController) CreateField(ctx *fiber.Ctx) error {
	var (
		reqBody models.CreateFieldRequest
		resBody *models.FieldResponse
		err     error
	)

	userID := helpers.GetUserIDFromContext(ctx)
	if err := ctrl.Options.UseCases.Validate.IsAdminUser(ctx.Context(), userID); err != nil {
		return helpers.ForbiddenResponse(ctx, constants.ErrAdminAccessRequired)
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}
	requestMap := map[string]any{
		"name": reqBody.Name,
	}
	if err := ctrl.Options.UseCases.Validate.IsValidRequestField(ctx.Context(), requestMap, "create"); err != nil {
		return helpers.BadRequestResponse(ctx, err.Error())
	}

	if reqBody.PricePerHour <= 0 {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidPrice)
	}

	resBody, err = ctrl.Options.UseCases.Field.CreateField(ctx.Context(), reqBody)
	if err != nil {

		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.CreatedResponse(ctx, resBody)
}

// GetFields godoc
// @Summary Get all fields
// @Description Get list of all available sports fields. PUBLIC ACCESS - No authentication required.
// @Tags Fields
// @Accept json
// @Produce json
// @Success 200 {object} models.BasicResponse{data=[]models.FieldResponse}
// @Router /fields [get]
func (ctrl *fieldController) GetFields(ctx *fiber.Ctx) error {
	fields, err := ctrl.Options.UseCases.Field.GetFields(ctx.Context())
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, fields)
}

// GetFieldByID godoc
// @Summary Get field by ID
// @Description Get details of a specific sports field. PUBLIC ACCESS - No authentication required.
// @Tags Fields
// @Accept json
// @Produce json
// @Param id path string true "Field ID (UUID format)"
// @Success 200 {object} models.BasicResponse{data=models.FieldResponse}
// @Failure 404 {object} models.BasicResponse
// @Router /fields/{id} [get]
func (ctrl *fieldController) GetFieldByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")

	if !helpers.IsValidUUID(id) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	field, err := ctrl.Options.UseCases.Field.GetFieldByID(ctx.Context(), id)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, field)
}

// UpdateField godoc
// @Summary Update field
// @Description Update sports field information. ADMIN ACCESS ONLY - Regular users cannot update fields.
// @Tags Fields
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Field ID (UUID format)"
// @Param request body models.UpdateFieldRequest true "Field update data"
// @Success 200 {object} models.BasicResponse{data=models.FieldResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 403 {object} models.BasicResponse
// @Failure 404 {object} models.BasicResponse
// @Router /fields/{id} [put]
func (ctrl *fieldController) UpdateField(ctx *fiber.Ctx) error {
	var (
		reqBody models.UpdateFieldRequest
		resBody *models.FieldResponse
		err     error
	)

	// Check if user is admin
	userID := helpers.GetUserIDFromContext(ctx)
	if err := ctrl.Options.UseCases.Validate.IsAdminUser(ctx.Context(), userID); err != nil {
		return helpers.ForbiddenResponse(ctx, constants.ErrAdminAccessRequired)
	}

	id := ctx.Params("id")

	if !helpers.IsValidUUID(id) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}

	requestMap := map[string]any{
		"id":   id,
		"name": reqBody.Name,
	}
	if err := ctrl.Options.UseCases.Validate.IsValidRequestField(ctx.Context(), requestMap, "update"); err != nil {
		return helpers.BadRequestResponse(ctx, err.Error())
	}

	if reqBody.PricePerHour <= 0 {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidPrice)
	}

	resBody, err = ctrl.Options.UseCases.Field.UpdateField(ctx.Context(), id, reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, resBody)
}

// DeleteField godoc
// @Summary Delete field
// @Description Delete a sports field. ADMIN ACCESS ONLY - Regular users cannot delete fields.
// @Tags Fields
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Field ID (UUID format)"
// @Success 200 {object} models.BasicResponse
// @Failure 403 {object} models.BasicResponse
// @Failure 404 {object} models.BasicResponse
// @Router /fields/{id} [delete]
func (ctrl *fieldController) DeleteField(ctx *fiber.Ctx) error {
	userID := helpers.GetUserIDFromContext(ctx)
	if err := ctrl.Options.UseCases.Validate.IsAdminUser(ctx.Context(), userID); err != nil {
		return helpers.ForbiddenResponse(ctx, constants.ErrAdminAccessRequired)
	}

	id := ctx.Params("id")

	if !helpers.IsValidUUID(id) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	err := ctrl.Options.UseCases.Field.DeleteField(ctx.Context(), id)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, nil)
}
