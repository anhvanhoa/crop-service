package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCycle(ctx context.Context, req *plantingCycleP.GetPlantingCycleRequest) (*plantingCycleP.PlantingCycleResponse, error) {
	plantingCycle, err := s.getPlantingCycleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if plantingCycle == nil {
		return &plantingCycleP.PlantingCycleResponse{}, nil
	}
	return &plantingCycleP.PlantingCycleResponse{
		PlantingCycle: s.createProtoPlantingCycle(plantingCycle),
	}, nil
}
