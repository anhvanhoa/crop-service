package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetCyclesWithDetailsUsecase defines the interface for getting planting cycles with details
type GetCyclesWithDetailsUsecase interface {
	Execute(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycleWithDetails, int64, error)
}

// getCyclesWithDetailsUsecase implements GetCyclesWithDetailsUsecase
type getCyclesWithDetailsUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetCyclesWithDetailsUsecase creates a new instance of GetCyclesWithDetailsUsecase
func NewGetCyclesWithDetailsUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetCyclesWithDetailsUsecase {
	return &getCyclesWithDetailsUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles with details
func (u *getCyclesWithDetailsUsecase) Execute(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycleWithDetails, int64, error) {
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

	// Get planting cycles with details from repository
	plantingCycles, err := u.plantingCycleRepo.GetCyclesWithDetails(ctx, filter)
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
