package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByHumidityRangeUsecase defines the interface for getting plant varieties by humidity range
type GetByHumidityRangeUsecase interface {
	Execute(ctx context.Context, minHumidity, maxHumidity float64) ([]*entity.PlantVariety, error)
}

// getByHumidityRangeUsecase implements GetByHumidityRangeUsecase
type getByHumidityRangeUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByHumidityRangeUsecase creates a new instance of GetByHumidityRangeUsecase
func NewGetByHumidityRangeUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByHumidityRangeUsecase {
	return &getByHumidityRangeUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties suitable for a humidity range
func (u *getByHumidityRangeUsecase) Execute(ctx context.Context, minHumidity, maxHumidity float64) ([]*entity.PlantVariety, error) {
	// Get plant varieties by humidity range from repository
	plantVarieties, err := u.plantVarietyRepo.GetVarietiesByHumidityRange(ctx, minHumidity, maxHumidity)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
