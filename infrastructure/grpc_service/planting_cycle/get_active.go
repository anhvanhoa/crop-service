package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlantingCycleService) GetActivePlantingCycles(ctx context.Context, req *emptypb.Empty) (*plantingCycleP.ListPlantingCyclesResponse, error) {
	plantingCycles, err := s.getActivePlantingCyclesUsecase.Execute(ctx)
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
