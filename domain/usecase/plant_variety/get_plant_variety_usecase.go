package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetPlantVarietyUsecase defines the interface for getting plant varieties
type GetPlantVarietyUsecase interface {
	Execute(ctx context.Context, id string) (*entity.PlantVariety, error)
}

// getPlantVarietyUsecase implements GetPlantVarietyUsecase
type getPlantVarietyUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetPlantVarietyUsecase creates a new instance of GetPlantVarietyUsecase
func NewGetPlantVarietyUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetPlantVarietyUsecase {
	return &getPlantVarietyUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves a plant variety by ID
func (u *getPlantVarietyUsecase) Execute(ctx context.Context, id string) (*entity.PlantVariety, error) {
	// Check if plant variety exists
	exists, err := u.plantVarietyRepo.Exists(ctx, id)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantVarietyNotFound
	}

	// Get plant variety from repository
	plantVariety, err := u.plantVarietyRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return plantVariety, nil
}
