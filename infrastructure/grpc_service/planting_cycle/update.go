package planting_cycle_service

import (
	"context"

	"farm-service/domain/entity"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) UpdatePlantingCycle(ctx context.Context, req *plantingCycleP.UpdatePlantingCycleRequest) (*plantingCycleP.PlantingCycleResponse, error) {
	plantingCycleReq, err := s.createEntityUpdatePlantingCycleReq(req)
	if err != nil {
		return nil, err
	}
	plantingCycle, err := s.updatePlantingCycleUsecase.Execute(ctx, plantingCycleReq)
	if err != nil {
		return nil, err
	}
	return &plantingCycleP.PlantingCycleResponse{
		PlantingCycle: s.createProtoPlantingCycle(plantingCycle),
	}, nil
}

func (s *PlantingCycleService) createEntityUpdatePlantingCycleReq(req *plantingCycleP.UpdatePlantingCycleRequest) (*entity.UpdatePlantingCycleRequest, error) {
	plantingCycle := &entity.UpdatePlantingCycleRequest{
		ID:             req.Id,
		CycleName:      req.CycleName,
		GrowingZoneID:  req.GrowingZoneId,
		PlantVarietyID: req.PlantVarietyId,
		PlantQuantity:  int(req.PlantQuantity),
		SeedBatch:      req.SeedBatch,
		Status:         req.Status,
		Notes:          req.Notes,
	}

	// Convert timestamps
	if req.SeedDate != nil {
		seedDate := req.SeedDate.AsTime()
		plantingCycle.SeedDate = &seedDate
	}
	if req.TransplantDate != nil {
		transplantDate := req.TransplantDate.AsTime()
		plantingCycle.TransplantDate = &transplantDate
	}
	if req.ExpectedHarvestDate != nil {
		expectedHarvestDate := req.ExpectedHarvestDate.AsTime()
		plantingCycle.ExpectedHarvestDate = &expectedHarvestDate
	}
	if req.ActualHarvestDate != nil {
		actualHarvestDate := req.ActualHarvestDate.AsTime()
		plantingCycle.ActualHarvestDate = &actualHarvestDate
	}

	return plantingCycle, nil
}
