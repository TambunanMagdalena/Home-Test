package repositories

import (
	"context"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror" // Ganti dari customerrors menjadi customerror

	"gorm.io/gorm"
)

type userRepository struct {
	Options Options
}

type UserInterface interface {
	CreateUser(ctx context.Context, user models.User) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
	FindByID(ctx context.Context, id string) (models.User, error)
	IsEmailExist(ctx context.Context, email string) (bool, error)
}

func (r *userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&user).Error
	return user, err
}

func (r *userRepository) FindByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User
	err := r.Options.Postgres.WithContext(ctx).Where("email = ?", email).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { // Gunakan gorm.ErrRecordNotFound, bukan string comparison
			return user, customerror.NewNotFoundErrorf(constants.ErrUserNotFound, email) // Ganti e. menjadi customerror.
		}
		return user, customerror.NewInternalServiceError(err.Error()) // Ganti e. menjadi customerror.
	}
	return user, nil
}

func (r *userRepository) FindByID(ctx context.Context, id string) (models.User, error) {
	var user models.User
	err := r.Options.Postgres.WithContext(ctx).Where("id = ?", id).First(&user).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { // Gunakan gorm.ErrRecordNotFound, bukan string comparison
			return user, customerror.NewNotFoundErrorf(constants.ErrUserNotFoundByID, id) // Ganti e. menjadi customerror.
		}
		return user, customerror.NewInternalServiceError(err.Error()) // Ganti e. menjadi customerror.
	}
	return user, nil
}

func (r *userRepository) IsEmailExist(ctx context.Context, email string) (bool, error) {
	var count int64
	err := r.Options.Postgres.WithContext(ctx).Model(&models.User{}).Where("email = ?", email).Count(&count).Error
	return count > 0, err
}
