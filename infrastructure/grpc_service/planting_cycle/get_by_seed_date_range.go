package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCyclesBySeedDateRange(ctx context.Context, req *plantingCycleP.GetPlantingCyclesBySeedDateRangeRequest) (*plantingCycleP.ListPlantingCyclesResponse, error) {
	startDate := req.StartDate.AsTime()
	endDate := req.EndDate.AsTime()

	plantingCycles, err := s.getBySeedDateRangeUsecase.Execute(ctx, startDate, endDate)
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
