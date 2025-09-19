package planting_cycle

import (
	"context"

	"farm-service/domain/repository"
)

// DeletePlantingCycleUsecase defines the interface for deleting planting cycles
type DeletePlantingCycleUsecase interface {
	Execute(ctx context.Context, id string) error
}

// deletePlantingCycleUsecase implements DeletePlantingCycleUsecase
type deletePlantingCycleUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewDeletePlantingCycleUsecase creates a new instance of DeletePlantingCycleUsecase
func NewDeletePlantingCycleUsecase(plantingCycleRepo repository.PlantingCycleRepository) DeletePlantingCycleUsecase {
	return &deletePlantingCycleUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute deletes a planting cycle by ID
func (u *deletePlantingCycleUsecase) Execute(ctx context.Context, id string) error {
	// Check if planting cycle exists
	exists, err := u.plantingCycleRepo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrPlantingCycleNotFound
	}

	// Delete planting cycle
	err = u.plantingCycleRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
