package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// SearchPlantVarietiesUsecase defines the interface for searching plant varieties
type SearchPlantVarietiesUsecase interface {
	Execute(ctx context.Context, name string) ([]*entity.PlantVariety, error)
}

// searchPlantVarietiesUsecase implements SearchPlantVarietiesUsecase
type searchPlantVarietiesUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewSearchPlantVarietiesUsecase creates a new instance of SearchPlantVarietiesUsecase
func NewSearchPlantVarietiesUsecase(plantVarietyRepo repository.PlantVarietyRepository) SearchPlantVarietiesUsecase {
	return &searchPlantVarietiesUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute searches plant varieties by name
func (u *searchPlantVarietiesUsecase) Execute(ctx context.Context, name string) ([]*entity.PlantVariety, error) {
	// Search plant varieties by name from repository
	plantVarieties, err := u.plantVarietyRepo.SearchByName(ctx, name)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
