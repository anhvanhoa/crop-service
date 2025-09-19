package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetCyclesByVarietyUsecase defines the interface for getting planting cycles by plant variety
type GetCyclesByVarietyUsecase interface {
	Execute(ctx context.Context, plantVarietyID string) ([]*entity.PlantingCycle, error)
}

// getCyclesByVarietyUsecase implements GetCyclesByVarietyUsecase
type getCyclesByVarietyUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetCyclesByVarietyUsecase creates a new instance of GetCyclesByVarietyUsecase
func NewGetCyclesByVarietyUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetCyclesByVarietyUsecase {
	return &getCyclesByVarietyUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves all planting cycles for a specific plant variety
func (u *getCyclesByVarietyUsecase) Execute(ctx context.Context, plantVarietyID string) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by plant variety
	plantingCycles, err := u.plantingCycleRepo.GetByPlantVarietyID(ctx, plantVarietyID)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
