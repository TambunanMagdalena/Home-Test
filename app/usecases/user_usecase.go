package usecase

import (
	"context"
	"take-home-test/app/models"
)

type userUsecase usecase

type UserInterface interface {
	GetUserByID(ctx context.Context, id string) (*models.UserResponse, error)
}

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*models.UserResponse, error) {
	user, err := u.Options.Repository.User.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	userResponse := &models.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
	}

	return userResponse, nil
}