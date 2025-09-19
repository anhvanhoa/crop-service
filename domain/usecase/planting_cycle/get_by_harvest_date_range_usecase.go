package planting_cycle

import (
	"context"
	"time"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByHarvestDateRangeUsecase defines the interface for getting planting cycles by harvest date range
type GetByHarvestDateRangeUsecase interface {
	Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)
}

// getByHarvestDateRangeUsecase implements GetByHarvestDateRangeUsecase
type getByHarvestDateRangeUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetByHarvestDateRangeUsecase creates a new instance of GetByHarvestDateRangeUsecase
func NewGetByHarvestDateRangeUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetByHarvestDateRangeUsecase {
	return &getByHarvestDateRangeUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles by harvest date range
func (u *getByHarvestDateRangeUsecase) Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by harvest date range from repository
	plantingCycles, err := u.plantingCycleRepo.GetCyclesByHarvestDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
