package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByWaterRequirementUsecase defines the interface for getting plant varieties by water requirement
type GetByWaterRequirementUsecase interface {
	Execute(ctx context.Context, waterRequirement string) ([]*entity.PlantVariety, error)
}

// getByWaterRequirementUsecase implements GetByWaterRequirementUsecase
type getByWaterRequirementUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByWaterRequirementUsecase creates a new instance of GetByWaterRequirementUsecase
func NewGetByWaterRequirementUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByWaterRequirementUsecase {
	return &getByWaterRequirementUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties by water requirement
func (u *getByWaterRequirementUsecase) Execute(ctx context.Context, waterRequirement string) ([]*entity.PlantVariety, error) {
	// Get plant varieties by water requirement from repository
	plantVarieties, err := u.plantVarietyRepo.GetVarietiesByWaterRequirement(ctx, waterRequirement)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
