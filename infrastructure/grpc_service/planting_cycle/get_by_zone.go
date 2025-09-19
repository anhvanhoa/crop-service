package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCyclesByZone(ctx context.Context, req *plantingCycleP.GetPlantingCyclesByZoneRequest) (*plantingCycleP.ListPlantingCyclesResponse, error) {
	plantingCycles, err := s.getCyclesByZoneUsecase.Execute(ctx, req.GrowingZoneId)
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
