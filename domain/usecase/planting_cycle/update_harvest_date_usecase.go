package planting_cycle

import (
	"context"
	"time"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdateHarvestDateUsecase defines the interface for updating planting cycle harvest date
type UpdateHarvestDateUsecase interface {
	Execute(ctx context.Context, id string, harvestDate time.Time) (*entity.PlantingCycle, error)
}

// updateHarvestDateUsecase implements UpdateHarvestDateUsecase
type updateHarvestDateUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewUpdateHarvestDateUsecase creates a new instance of UpdateHarvestDateUsecase
func NewUpdateHarvestDateUsecase(plantingCycleRepo repository.PlantingCycleRepository) UpdateHarvestDateUsecase {
	return &updateHarvestDateUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute updates the harvest date of a planting cycle
func (u *updateHarvestDateUsecase) Execute(ctx context.Context, id string, harvestDate time.Time) (*entity.PlantingCycle, error) {
	// Check if planting cycle exists
	exists, err := u.plantingCycleRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantingCycleNotFound
	}

	// Update harvest date
	err = u.plantingCycleRepo.UpdateHarvestDate(ctx, id, harvestDate)
	if err != nil {
		return nil, err
	}

	// Get updated planting cycle
	plantingCycle, err := u.plantingCycleRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return plantingCycle, nil
}
