package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetCyclesByZoneUsecase defines the interface for getting planting cycles by growing zone
type GetCyclesByZoneUsecase interface {
	Execute(ctx context.Context, growingZoneID string) ([]*entity.PlantingCycle, error)
}

// getCyclesByZoneUsecase implements GetCyclesByZoneUsecase
type getCyclesByZoneUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetCyclesByZoneUsecase creates a new instance of GetCyclesByZoneUsecase
func NewGetCyclesByZoneUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetCyclesByZoneUsecase {
	return &getCyclesByZoneUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves all planting cycles for a specific growing zone
func (u *getCyclesByZoneUsecase) Execute(ctx context.Context, growingZoneID string) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by growing zone
	plantingCycles, err := u.plantingCycleRepo.GetByGrowingZoneID(ctx, growingZoneID)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
