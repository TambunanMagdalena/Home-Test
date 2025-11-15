package controllers

import (
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/pkg/customerror"

	"github.com/gofiber/fiber/v2"
)

type userController struct {
	Options Options
}

type UserInterface interface {
	GetProfile(ctx *fiber.Ctx) error
	GetUserByID(ctx *fiber.Ctx) error
}

// GetProfile godoc
// @Summary Get user profile
// @Description Get authenticated user's profile information
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/profile [get]
func (c *userController) GetProfile(ctx *fiber.Ctx) error {
	userID := helpers.GetUserIDFromContext(ctx)
	if userID == "" {
		return helpers.UnauthorizedResponse(ctx, constants.ErrMissingToken)
	}

	user, err := c.Options.UseCases.User.GetUserByID(ctx.Context(), userID)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, user)
}

// GetUserByID godoc
// @Summary Get user by ID
// @Description Get user details by ID (Admin only)
// @Tags Users
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "User ID"
// @Success 200 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /users/{id} [get]
func (c *userController) GetUserByID(ctx *fiber.Ctx) error {
	currentUserID := helpers.GetUserIDFromContext(ctx)
	if err := c.Options.UseCases.Validate.IsAdminUser(ctx.Context(), currentUserID); err != nil {
		return helpers.ForbiddenResponse(ctx, constants.ErrAdminAccessRequired) // ✅ Gunakan constant
	}

	id := ctx.Params("id")

	if !helpers.IsValidUUID(id) {
		return helpers.BadRequestResponse(ctx, constants.ErrInvalidUUID)
	}

	user, err := c.Options.UseCases.User.GetUserByID(ctx.Context(), id)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.SuccessResponse(ctx, user) // ✅ Gunakan helper convenience
}
