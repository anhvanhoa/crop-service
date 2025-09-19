package planting_cycle

import (
	"context"
	"time"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByDateRangeUsecase defines the interface for getting planting cycles by date range
type GetByDateRangeUsecase interface {
	Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)
}

// getByDateRangeUsecase implements GetByDateRangeUsecase
type getByDateRangeUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetByDateRangeUsecase creates a new instance of GetByDateRangeUsecase
func NewGetByDateRangeUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetByDateRangeUsecase {
	return &getByDateRangeUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles within a date range
func (u *getByDateRangeUsecase) Execute(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by date range from repository
	plantingCycles, err := u.plantingCycleRepo.GetCyclesByDateRange(ctx, startDate, endDate)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
