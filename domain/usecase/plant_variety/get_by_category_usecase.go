package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByCategoryUsecase defines the interface for getting plant varieties by category
type GetByCategoryUsecase interface {
	Execute(ctx context.Context, category string) ([]*entity.PlantVariety, error)
}

// getByCategoryUsecase implements GetByCategoryUsecase
type getByCategoryUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByCategoryUsecase creates a new instance of GetByCategoryUsecase
func NewGetByCategoryUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByCategoryUsecase {
	return &getByCategoryUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties by category
func (u *getByCategoryUsecase) Execute(ctx context.Context, category string) ([]*entity.PlantVariety, error) {
	// Get plant varieties by category from repository
	plantVarieties, err := u.plantVarietyRepo.GetByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
