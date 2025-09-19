package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetOverdueHarvestsUsecase defines the interface for getting overdue harvests
type GetOverdueHarvestsUsecase interface {
	Execute(ctx context.Context) ([]*entity.PlantingCycle, error)
}

// getOverdueHarvestsUsecase implements GetOverdueHarvestsUsecase
type getOverdueHarvestsUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetOverdueHarvestsUsecase creates a new instance of GetOverdueHarvestsUsecase
func NewGetOverdueHarvestsUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetOverdueHarvestsUsecase {
	return &getOverdueHarvestsUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles with overdue harvests
func (u *getOverdueHarvestsUsecase) Execute(ctx context.Context) ([]*entity.PlantingCycle, error) {
	// Get overdue harvests from repository
	plantingCycles, err := u.plantingCycleRepo.GetOverdueHarvests(ctx)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
