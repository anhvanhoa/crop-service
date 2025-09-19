package plant_variety

import (
	"context"

	"farm-service/domain/repository"
)

// DeletePlantVarietyUsecase defines the interface for deleting plant varieties
type DeletePlantVarietyUsecase interface {
	Execute(ctx context.Context, id string) error
}

// deletePlantVarietyUsecase implements DeletePlantVarietyUsecase
type deletePlantVarietyUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewDeletePlantVarietyUsecase creates a new instance of DeletePlantVarietyUsecase
func NewDeletePlantVarietyUsecase(plantVarietyRepo repository.PlantVarietyRepository) DeletePlantVarietyUsecase {
	return &deletePlantVarietyUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute deletes a plant variety by ID
func (u *deletePlantVarietyUsecase) Execute(ctx context.Context, id string) error {
	// Check if plant variety exists
	exists, err := u.plantVarietyRepo.Exists(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrPlantVarietyNotFound
	}

	// Delete plant variety
	err = u.plantVarietyRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
