package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdatePlantVarietyUsecase defines the interface for updating plant varieties
type UpdatePlantVarietyUsecase interface {
	Execute(ctx context.Context, req *entity.UpdatePlantVarietyRequest) (*entity.PlantVariety, error)
}

// updatePlantVarietyUsecase implements UpdatePlantVarietyUsecase
type updatePlantVarietyUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewUpdatePlantVarietyUsecase creates a new instance of UpdatePlantVarietyUsecase
func NewUpdatePlantVarietyUsecase(plantVarietyRepo repository.PlantVarietyRepository) UpdatePlantVarietyUsecase {
	return &updatePlantVarietyUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute updates an existing plant variety
func (u *updatePlantVarietyUsecase) Execute(ctx context.Context, req *entity.UpdatePlantVarietyRequest) (*entity.PlantVariety, error) {
	// Check if plant variety exists
	exists, err := u.plantVarietyRepo.Exists(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantVarietyNotFound
	}

	data := &entity.PlantVariety{
		ID:                 req.ID,
		Name:               req.Name,
		ScientificName:     req.ScientificName,
		Category:           req.Category,
		GrowingSeason:      req.GrowingSeason,
		GrowthDurationDays: req.GrowthDurationDays,
		OptimalTempMin:     req.OptimalTempMin,
		OptimalTempMax:     req.OptimalTempMax,
		OptimalHumidityMin: req.OptimalHumidityMin,
		OptimalHumidityMax: req.OptimalHumidityMax,
		PHMin:              req.PHMin,
		PHMax:              req.PHMax,
		WaterRequirement:   req.WaterRequirement,
		LightRequirement:   req.LightRequirement,
		Description:        req.Description,
		MediaID:            req.MediaID,
		Status:             req.Status,
	}

	// Save updated plant variety
	err = u.plantVarietyRepo.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
