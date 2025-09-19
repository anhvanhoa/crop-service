package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetCycleWithDetailsUsecase defines the interface for getting a single planting cycle with details
type GetCycleWithDetailsUsecase interface {
	Execute(ctx context.Context, id string) (*entity.PlantingCycleWithDetails, error)
}

// getCycleWithDetailsUsecase implements GetCycleWithDetailsUsecase
type getCycleWithDetailsUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetCycleWithDetailsUsecase creates a new instance of GetCycleWithDetailsUsecase
func NewGetCycleWithDetailsUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetCycleWithDetailsUsecase {
	return &getCycleWithDetailsUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves a single planting cycle with details
func (u *getCycleWithDetailsUsecase) Execute(ctx context.Context, id string) (*entity.PlantingCycleWithDetails, error) {
	// Get planting cycle with details from repository
	plantingCycle, err := u.plantingCycleRepo.GetCycleWithDetails(ctx, id)
	if err != nil {
		return nil, err
	}

	return plantingCycle, nil
}
