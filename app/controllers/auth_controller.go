package controllers

import (
	"take-home-test/app/constants"
	"take-home-test/app/helpers"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror"

	"github.com/gofiber/fiber/v2"
)

type authController struct {
	Options Options
}

type AuthInterface interface {
	Register(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
}

// Register godoc
// @Summary Register new user
// @Description Register a new user account
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.RegisterRequest true "Register request"
// @Success 201 {object} models.BasicResponse{data=models.UserResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 500 {object} models.BasicResponse
// @Router /auth/register [post]
func (ctrl *authController) Register(ctx *fiber.Ctx) error {
	var (
		reqBody models.RegisterRequest
		resBody *models.UserResponse
		err     error
	)

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}

	if err := helpers.ValidateRegisterRequest(reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, err.Error())
	}

	if reqBody.Role == "" {
		reqBody.Role = constants.ROLE_USER
	}

	resBody, err = ctrl.Options.UseCases.Auth.Register(ctx.Context(), reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.StandardResponse(ctx, fiber.StatusCreated, []string{constants.REGISTER_SUCCESS_MESSAGE}, resBody, nil)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and return JWT token
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body models.LoginRequest true "Login request"
// @Success 200 {object} models.BasicResponse{data=models.LoginResponse}
// @Failure 400 {object} models.BasicResponse
// @Failure 401 {object} models.BasicResponse
// @Router /auth/login [post]
func (ctrl *authController) Login(ctx *fiber.Ctx) error {
	var (
		reqBody models.LoginRequest
		resBody *models.LoginResponse
		err     error
	)

	if err := ctx.BodyParser(&reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, constants.ErrBadRequest)
	}

	if err := helpers.ValidateLoginRequest(reqBody); err != nil {
		return helpers.BadRequestResponse(ctx, err.Error())
	}

	resBody, err = ctrl.Options.UseCases.Auth.Login(ctx.Context(), reqBody)
	if err != nil {
		return helpers.StandardResponse(ctx, customerror.GetStatusCode(err), []string{err.Error()}, nil, nil)
	}

	return helpers.StandardResponse(ctx, fiber.StatusOK, []string{constants.LOGIN_SUCCESS_MESSAGE}, resBody, nil)
}
