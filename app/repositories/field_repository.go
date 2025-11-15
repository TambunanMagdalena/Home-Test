package repositories

import (
	"context"
	"take-home-test/app/constants"
	"take-home-test/app/models"
	"take-home-test/pkg/customerror" // Ganti dari customerrors menjadi customerror

	"gorm.io/gorm"
)

type fieldRepository struct {
	Options Options
}

type FieldInterface interface {
	CreateField(ctx context.Context, field models.Field) (models.Field, error)
	GetFields(ctx context.Context) ([]models.Field, error)
	GetFieldByID(ctx context.Context, id string) (models.Field, error)
	UpdateField(ctx context.Context, field models.Field) (models.Field, error)
	DeleteField(ctx context.Context, id string) error
}

func (r *fieldRepository) CreateField(ctx context.Context, field models.Field) (models.Field, error) {
	err := r.Options.Postgres.WithContext(ctx).Create(&field).Error
	return field, err
}

func (r *fieldRepository) GetFields(ctx context.Context) ([]models.Field, error) {
	var fields []models.Field
	err := r.Options.Postgres.WithContext(ctx).Find(&fields).Error
	return fields, err
}

func (r *fieldRepository) GetFieldByID(ctx context.Context, id string) (models.Field, error) {
	var field models.Field
	err := r.Options.Postgres.WithContext(ctx).Where("id = ?", id).First(&field).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound { // Gunakan gorm.ErrRecordNotFound, bukan string comparison
			return field, customerror.NewNotFoundErrorf(constants.ErrFieldNotFound, id) // Ganti e. menjadi customerror.
		}
		return field, customerror.NewInternalServiceError(err.Error()) // Ganti e. menjadi customerror.
	}
	return field, nil
}

func (r *fieldRepository) UpdateField(ctx context.Context, field models.Field) (models.Field, error) {
	err := r.Options.Postgres.WithContext(ctx).Save(&field).Error
	return field, err
}

func (r *fieldRepository) DeleteField(ctx context.Context, id string) error {
	result := r.Options.Postgres.WithContext(ctx).Where("id = ?", id).Delete(&models.Field{})

	if result.Error != nil {
		return customerror.NewInternalServiceError(result.Error.Error()) // Ganti e. menjadi customerror.
	}

	if result.RowsAffected == 0 {
		return customerror.NewNotFoundErrorf(constants.ErrFieldNotFound, id) // Ganti e. menjadi customerror.
	}

	return nil
}
