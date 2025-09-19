package planting_cycle

import (
	"context"
	"time"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetBySeedDateRangeUsecase defines the interface for getting planting cycles by seed date range
type GetBySeedDateRangeUsecase interface {
	Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)
}

// getBySeedDateRangeUsecase implements GetBySeedDateRangeUsecase
type getBySeedDateRangeUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetBySeedDateRangeUsecase creates a new instance of GetBySeedDateRangeUsecase
func NewGetBySeedDateRangeUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetBySeedDateRangeUsecase {
	return &getBySeedDateRangeUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles by seed date range
func (u *getBySeedDateRangeUsecase) Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by seed date range from repository
	plantingCycles, err := u.plantingCycleRepo.GetCyclesBySeedDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
