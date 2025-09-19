package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetBySeasonUsecase defines the interface for getting plant varieties by season
type GetBySeasonUsecase interface {
	Execute(ctx context.Context, season string) ([]*entity.PlantVariety, error)
}

// getBySeasonUsecase implements GetBySeasonUsecase
type getBySeasonUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetBySeasonUsecase creates a new instance of GetBySeasonUsecase
func NewGetBySeasonUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetBySeasonUsecase {
	return &getBySeasonUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties by growing season
func (u *getBySeasonUsecase) Execute(ctx context.Context, season string) ([]*entity.PlantVariety, error) {
	// Get plant varieties by growing season from repository
	plantVarieties, err := u.plantVarietyRepo.GetByGrowingSeason(ctx, season)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
