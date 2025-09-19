package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdateStatusUsecase defines the interface for updating planting cycle status
type UpdateStatusUsecase interface {
	Execute(ctx context.Context, id, status string) (*entity.PlantingCycle, error)
}

// updateStatusUsecase implements UpdateStatusUsecase
type updateStatusUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewUpdateStatusUsecase creates a new instance of UpdateStatusUsecase
func NewUpdateStatusUsecase(plantingCycleRepo repository.PlantingCycleRepository) UpdateStatusUsecase {
	return &updateStatusUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute updates the status of a planting cycle
func (u *updateStatusUsecase) Execute(ctx context.Context, id, status string) (*entity.PlantingCycle, error) {
	// Check if planting cycle exists
	exists, err := u.plantingCycleRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantingCycleNotFound
	}

	// Update status
	err = u.plantingCycleRepo.UpdateStatus(ctx, id, status)
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
