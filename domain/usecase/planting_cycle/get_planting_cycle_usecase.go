package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetPlantingCycleUsecase defines the interface for getting planting cycles
type GetPlantingCycleUsecase interface {
	Execute(ctx context.Context, id string) (*entity.PlantingCycle, error)
}

// getPlantingCycleUsecase implements GetPlantingCycleUsecase
type getPlantingCycleUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetPlantingCycleUsecase creates a new instance of GetPlantingCycleUsecase
func NewGetPlantingCycleUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetPlantingCycleUsecase {
	return &getPlantingCycleUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves a planting cycle by ID
func (u *getPlantingCycleUsecase) Execute(ctx context.Context, id string) (*entity.PlantingCycle, error) {
	// Check if planting cycle exists
	exists, err := u.plantingCycleRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantingCycleNotFound
	}

	// Get planting cycle from repository
	plantingCycle, err := u.plantingCycleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return plantingCycle, nil
}
