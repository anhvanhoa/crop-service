package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetActivePlantingCyclesUsecase defines the interface for getting active planting cycles
type GetActivePlantingCyclesUsecase interface {
	Execute(ctx context.Context) ([]*entity.PlantingCycle, error)
}

// getActivePlantingCyclesUsecase implements GetActivePlantingCyclesUsecase
type getActivePlantingCyclesUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetActivePlantingCyclesUsecase creates a new instance of GetActivePlantingCyclesUsecase
func NewGetActivePlantingCyclesUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetActivePlantingCyclesUsecase {
	return &getActivePlantingCyclesUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves all active planting cycles
func (u *getActivePlantingCyclesUsecase) Execute(ctx context.Context) ([]*entity.PlantingCycle, error) {
	// Get active planting cycles from repository
	plantingCycles, err := u.plantingCycleRepo.GetActiveCycles(ctx)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
