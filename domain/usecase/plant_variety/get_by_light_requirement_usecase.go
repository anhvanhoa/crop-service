package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByLightRequirementUsecase defines the interface for getting plant varieties by light requirement
type GetByLightRequirementUsecase interface {
	Execute(ctx context.Context, lightRequirement string) ([]*entity.PlantVariety, error)
}

// getByLightRequirementUsecase implements GetByLightRequirementUsecase
type getByLightRequirementUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByLightRequirementUsecase creates a new instance of GetByLightRequirementUsecase
func NewGetByLightRequirementUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByLightRequirementUsecase {
	return &getByLightRequirementUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties by light requirement
func (u *getByLightRequirementUsecase) Execute(ctx context.Context, lightRequirement string) ([]*entity.PlantVariety, error) {
	// Get plant varieties by light requirement from repository
	plantVarieties, err := u.plantVarietyRepo.GetVarietiesByLightRequirement(ctx, lightRequirement)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
