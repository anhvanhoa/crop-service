package plant_variety

import (
	"context"

	"farm-service/domain/entity"
	"farm-service/domain/repository"
)

// GetByStatusUsecase defines the interface for getting plant varieties by status
type GetByStatusUsecase interface {
	Execute(ctx context.Context, status string) ([]*entity.PlantVariety, error)
}

// getByStatusUsecase implements GetByStatusUsecase
type getByStatusUsecase struct {
	plantVarietyRepo repository.PlantVarietyRepository
}

// NewGetByStatusUsecase creates a new instance of GetByStatusUsecase
func NewGetByStatusUsecase(plantVarietyRepo repository.PlantVarietyRepository) GetByStatusUsecase {
	return &getByStatusUsecase{
		plantVarietyRepo: plantVarietyRepo,
	}
}

// Execute retrieves plant varieties by status
func (u *getByStatusUsecase) Execute(ctx context.Context, status string) ([]*entity.PlantVariety, error) {
	// Get plant varieties by status from repository
	plantVarieties, err := u.plantVarietyRepo.GetByStatus(ctx, status)
	if err != nil {
		return nil, err
	}

	return plantVarieties, nil
}
