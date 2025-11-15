package usecase

import (
	"context"
	"errors"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type authUsecase usecase

type AuthInterface interface {
	Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, error)
	Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error)
}

func (u *authUsecase) Register(ctx context.Context, req models.RegisterRequest) (*models.UserResponse, error) {
	// Check if email already exists
	exists, err := u.Options.Repository.User.IsEmailExist(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.New(constants.ErrDuplicateEmail)
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// Create user
	user := models.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	createdUser, err := u.Options.Repository.User.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	// Return user response
	userResponse := &models.UserResponse{
		ID:        createdUser.ID,
		Name:      createdUser.Name,
		Email:     createdUser.Email,
		Role:      createdUser.Role,
		CreatedAt: createdUser.CreatedAt,
	}

	return userResponse, nil
}

func (u *authUsecase) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	// Find user by email
	user, err := u.Options.Repository.User.FindByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New(constants.ErrInvalidCredentials)
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New(constants.ErrInvalidCredentials)
	}

	// Generate JWT token
	token, err := u.generateJWT(user)
	if err != nil {
		return nil, err
	}

	// Create user response
	userResponse := models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	loginResponse := &models.LoginResponse{
		Token: token,
		User:  userResponse,
	}

	return loginResponse, nil
}

func (u *authUsecase) generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(u.Options.Config.GetJWTSecret()))
}