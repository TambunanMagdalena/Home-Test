package usecase

import (
	"context"
	"take-home-test/app/models"
)

type fieldUsecase usecase

type FieldInterface interface {
	CreateField(ctx context.Context, req models.CreateFieldRequest) (*models.FieldResponse, error)
	GetFields(ctx context.Context) ([]models.FieldResponse, error)
	GetFieldByID(ctx context.Context, id string) (*models.FieldResponse, error)
	UpdateField(ctx context.Context, id string, req models.UpdateFieldRequest) (*models.FieldResponse, error)
	DeleteField(ctx context.Context, id string) error
}

func (u *fieldUsecase) CreateField(ctx context.Context, req models.CreateFieldRequest) (*models.FieldResponse, error) {
	field := models.Field{
		Name:         req.Name,
		PricePerHour: req.PricePerHour,
		Location:     req.Location,
	}

	createdField, err := u.Options.Repository.Field.CreateField(ctx, field)
	if err != nil {
		return nil, err
	}

	fieldResponse := &models.FieldResponse{
		ID:           createdField.ID,
		Name:         createdField.Name,
		PricePerHour: createdField.PricePerHour,
		Location:     createdField.Location,
		CreatedAt:    createdField.CreatedAt,
		UpdatedAt:    createdField.UpdatedAt,
	}

	return fieldResponse, nil
}

func (u *fieldUsecase) GetFields(ctx context.Context) ([]models.FieldResponse, error) {
	fields, err := u.Options.Repository.Field.GetFields(ctx)
	if err != nil {
		return nil, err
	}

	var fieldResponses []models.FieldResponse
	for _, field := range fields {
		fieldResponses = append(fieldResponses, models.FieldResponse{
			ID:           field.ID,
			Name:         field.Name,
			PricePerHour: field.PricePerHour,
			Location:     field.Location,
			CreatedAt:    field.CreatedAt,
			UpdatedAt:    field.UpdatedAt,
		})
	}

	return fieldResponses, nil
}

func (u *fieldUsecase) GetFieldByID(ctx context.Context, id string) (*models.FieldResponse, error) {
	field, err := u.Options.Repository.Field.GetFieldByID(ctx, id)
	if err != nil {
		return nil, err
	}

	fieldResponse := &models.FieldResponse{
		ID:           field.ID,
		Name:         field.Name,
		PricePerHour: field.PricePerHour,
		Location:     field.Location,
		CreatedAt:    field.CreatedAt,
		UpdatedAt:    field.UpdatedAt,
	}

	return fieldResponse, nil
}

func (u *fieldUsecase) UpdateField(ctx context.Context, id string, req models.UpdateFieldRequest) (*models.FieldResponse, error) {
	// Get existing field
	existingField, err := u.Options.Repository.Field.GetFieldByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Update field data
	existingField.Name = req.Name
	existingField.PricePerHour = req.PricePerHour
	existingField.Location = req.Location

	updatedField, err := u.Options.Repository.Field.UpdateField(ctx, existingField)
	if err != nil {
		return nil, err
	}

	fieldResponse := &models.FieldResponse{
		ID:           updatedField.ID,
		Name:         updatedField.Name,
		PricePerHour: updatedField.PricePerHour,
		Location:     updatedField.Location,
		CreatedAt:    updatedField.CreatedAt,
		UpdatedAt:    updatedField.UpdatedAt,
	}

	return fieldResponse, nil
}

func (u *fieldUsecase) DeleteField(ctx context.Context, id string) error {
	return u.Options.Repository.Field.DeleteField(ctx, id)
}
