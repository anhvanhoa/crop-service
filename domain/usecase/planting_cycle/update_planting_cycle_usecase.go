package planting_cycle

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// UpdatePlantingCycleUsecase defines the interface for updating planting cycles
type UpdatePlantingCycleUsecase interface {
	Execute(ctx context.Context, req *entity.UpdatePlantingCycleRequest) (*entity.PlantingCycle, error)
}

// updatePlantingCycleUsecase implements UpdatePlantingCycleUsecase
type updatePlantingCycleUsecase struct {
	plantingCycleRepo repository.PlantingCycleRepository
}

// NewUpdatePlantingCycleUsecase creates a new instance of UpdatePlantingCycleUsecase
func NewUpdatePlantingCycleUsecase(plantingCycleRepo repository.PlantingCycleRepository) UpdatePlantingCycleUsecase {
	return &updatePlantingCycleUsecase{
		plantingCycleRepo: plantingCycleRepo,
	}
}

// Execute updates an existing planting cycle
func (u *updatePlantingCycleUsecase) Execute(ctx context.Context, req *entity.UpdatePlantingCycleRequest) (*entity.PlantingCycle, error) {
	// Check if planting cycle exists
	exists, err := u.plantingCycleRepo.Exists(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, ErrPlantingCycleNotFound
	}

	data := &entity.PlantingCycle{
		ID:                  req.ID,
		CycleName:           req.CycleName,
		GrowingZoneID:       req.GrowingZoneID,
		PlantVarietyID:      req.PlantVarietyID,
		SeedDate:            req.SeedDate,
		TransplantDate:      req.TransplantDate,
		ExpectedHarvestDate: req.ExpectedHarvestDate,
		ActualHarvestDate:   req.ActualHarvestDate,
		PlantQuantity:       req.PlantQuantity,
		SeedBatch:           req.SeedBatch,
		Status:              req.Status,
		Notes:               req.Notes,
	}

	// Save updated planting cycle
	err = u.plantingCycleRepo.Update(ctx, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
