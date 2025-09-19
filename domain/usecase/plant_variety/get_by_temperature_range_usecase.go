package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByTemperatureRangeUsecase defines the interface for getting plant varieties by temperature range
type GetByTemperatureRangeUsecase interface {
	Execute(ctx context.Context, minTemp, maxTemp float64) ([]*entity.PlantVariety, error)
}

// getByTemperatureRangeUsecase implements GetByTemperatureRangeUsecase
type getByTemperatureRangeUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByTemperatureRangeUsecase creates a new instance of GetByTemperatureRangeUsecase
func NewGetByTemperatureRangeUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByTemperatureRangeUsecase {
	return &getByTemperatureRangeUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties suitable for a temperature range
func (u *getByTemperatureRangeUsecase) Execute(ctx context.Context, minTemp, maxTemp float64) ([]*entity.PlantVariety, error) {
	// Get plant varieties by temperature range from repository
	plantVarieties, err := u.plantVarietyRepo.GetVarietiesByTemperatureRange(ctx, minTemp, maxTemp)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
