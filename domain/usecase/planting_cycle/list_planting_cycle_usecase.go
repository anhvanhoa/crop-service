package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// ListPlantingCycleUsecase defines the interface for listing planting cycles
type ListPlantingCycleUsecase interface {
	Execute(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycle, int64, error)
}

// listPlantingCycleUsecase implements ListPlantingCycleUsecase
type listPlantingCycleUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewListPlantingCycleUsecase creates a new instance of ListPlantingCycleUsecase
func NewListPlantingCycleUsecase(plantingCycleRepo repository.PlantingCycleRepository) ListPlantingCycleUsecase {
	return &listPlantingCycleUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves a list of planting cycles with filtering and pagination
func (u *listPlantingCycleUsecase) Execute(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycle, int64, error) {
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

	// Get planting cycles from repository
	plantingCycles, err := u.plantingCycleRepo.List(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	// Get total count
	total, err := u.plantingCycleRepo.Count(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	return plantingCycles, total, nil
}
