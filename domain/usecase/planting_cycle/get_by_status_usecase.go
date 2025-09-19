package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByStatusUsecase defines the interface for getting planting cycles by status
type GetByStatusUsecase interface {
	Execute(ctx context.Context, status string) ([]*entity.PlantingCycle, error)
}

// getByStatusUsecase implements GetByStatusUsecase
type getByStatusUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewGetByStatusUsecase creates a new instance of GetByStatusUsecase
func NewGetByStatusUsecase(plantingCycleRepo repository.PlantingCycleRepository) GetByStatusUsecase {
	return &getByStatusUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute retrieves planting cycles by status
func (u *getByStatusUsecase) Execute(ctx context.Context, status string) ([]*entity.PlantingCycle, error) {
	// Get planting cycles by status from repository
	plantingCycles, err := u.plantingCycleRepo.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return plantingCycles, nil
}
