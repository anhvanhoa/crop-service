package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) UpdatePlantingCycleStatus(ctx context.Context, req *plantingCycleP.UpdatePlantingCycleStatusRequest) (*plantingCycleP.PlantingCycleResponse, error) {
	plantingCycle, err := s.updateStatusUsecase.Execute(ctx, req.Id, req.Status)
	if err != nil {
		return nil, err
	}

	return &plantingCycleP.PlantingCycleResponse{
		PlantingCycle: s.createProtoPlantingCycle(plantingCycle),
	}, nil
}
