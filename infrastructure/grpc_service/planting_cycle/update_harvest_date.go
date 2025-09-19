package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) UpdatePlantingCycleHarvestDate(ctx context.Context, req *plantingCycleP.UpdatePlantingCycleHarvestDateRequest) (*plantingCycleP.PlantingCycleResponse, error) {
	harvestDate := req.HarvestDate.AsTime()

	plantingCycle, err := s.updateHarvestDateUsecase.Execute(ctx, req.Id, harvestDate)
	if err != nil {
		return nil, err
	}

	return &plantingCycleP.PlantingCycleResponse{
		PlantingCycle: s.createProtoPlantingCycle(plantingCycle),
	}, nil
}
