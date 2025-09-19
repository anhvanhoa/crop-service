package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCycleWithDetails(ctx context.Context, req *plantingCycleP.GetPlantingCycleWithDetailsRequest) (*plantingCycleP.PlantingCycleWithDetailsResponse, error) {
	plantingCycle, err := s.getCycleWithDetailsUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if plantingCycle == nil {
		return &plantingCycleP.PlantingCycleWithDetailsResponse{}, nil
	}

	return &plantingCycleP.PlantingCycleWithDetailsResponse{
		PlantingCycleWithDetails: s.createProtoPlantingCycleWithDetails(plantingCycle),
	}, nil
}
