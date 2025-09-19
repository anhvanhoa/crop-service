package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCyclesByVariety(ctx context.Context, req *plantingCycleP.GetPlantingCyclesByVarietyRequest) (*plantingCycleP.ListPlantingCyclesResponse, error) {
	plantingCycles, err := s.getCyclesByVarietyUsecase.Execute(ctx, req.PlantVarietyId)
	if err != nil {
		return nil, err
	}

	var protoPlantingCycles []*plantingCycleP.PlantingCycle
	for _, plantingCycle := range plantingCycles {
		protoPlantingCycles = append(protoPlantingCycles, s.createProtoPlantingCycle(plantingCycle))
	}

	return &plantingCycleP.ListPlantingCyclesResponse{
		PlantingCycles: protoPlantingCycles,
		Total:          int64(len(protoPlantingCycles)),
	}, nil
}
