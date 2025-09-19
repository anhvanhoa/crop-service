package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// ListPlantVarietyUsecase defines the interface for listing plant varieties
type ListPlantVarietyUsecase interface {
	Execute(ctx context.Context, filter *entity.PlantVarietyFilter) ([]*entity.PlantVariety, int64, error)
}

// listPlantVarietyUsecase implements ListPlantVarietyUsecase
type listPlantVarietyUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewListPlantVarietyUsecase creates a new instance of ListPlantVarietyUsecase
func NewListPlantVarietyUsecase(plantVarietyRepo repository.PlantVarietyRepository) ListPlantVarietyUsecase {
	return &listPlantVarietyUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves a list of plant varieties with filtering and pagination
func (u *listPlantVarietyUsecase) Execute(ctx context.Context, filter *entity.PlantVarietyFilter) ([]*entity.PlantVariety, int64, error) {
	// Set default values for pagination
	if filter.Limit <= 0 {
		filter.Limit = 10
	}
	if filter.Limit > 100 {
		filter.Limit = 100
	}
	if filter.Offset < 0 {
		filter.Offset = 0
	}

	// Get plant varieties from repository
	plantVarieties, err := u.plantVarietyRepo.List(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := u.plantVarietyRepo.Count(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return plantVarieties, total, nil
}
