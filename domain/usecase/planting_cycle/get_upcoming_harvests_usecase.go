package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetUpcomingHarvestsUsecase defines the interface for getting upcoming harvests
type GetUpcomingHarvestsUsecase interface {
	Execute(ctx context.Context, days int) ([]*entity.PlantingCycle, error)
}

// getUpcomingHarvestsUsecase implements GetUpcomingHarvestsUsecase
type getUpcomingHarvestsUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetUpcomingHarvestsUsecase creates a new instance of GetUpcomingHarvestsUsecase
func NewGetUpcomingHarvestsUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetUpcomingHarvestsUsecase {
	return &getUpcomingHarvestsUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles with upcoming harvests
func (u *getUpcomingHarvestsUsecase) Execute(ctx context.Context, days int) ([]*entity.PlantingCycle, error) {
	// Get upcoming harvests from repository
	plantingCycles, err := u.plantingCycleRepo.GetUpcomingHarvests(ctx, days)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
