package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetActivePlantVarietiesUsecase defines the interface for getting active plant varieties
type GetActivePlantVarietiesUsecase interface {
	Execute(ctx context.Context) ([]*entity.PlantVariety, error)
}

// getActivePlantVarietiesUsecase implements GetActivePlantVarietiesUsecase
type getActivePlantVarietiesUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetActivePlantVarietiesUsecase creates a new instance of GetActivePlantVarietiesUsecase
func NewGetActivePlantVarietiesUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetActivePlantVarietiesUsecase {
	return &getActivePlantVarietiesUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves all active plant varieties
func (u *getActivePlantVarietiesUsecase) Execute(ctx context.Context) ([]*entity.PlantVariety, error) {
	// Get active plant varieties from repository
	plantVarieties, err := u.plantVarietyRepo.GetActiveVarieties(ctx)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
