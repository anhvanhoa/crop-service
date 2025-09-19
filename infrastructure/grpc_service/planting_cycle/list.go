package planting_cycle_service

import (
	"context"

	"farm-service/domain/entity"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) ListPlantingCycles(ctx context.Context, req *plantingCycleP.ListPlantingCyclesRequest) (*plantingCycleP.ListPlantingCyclesResponse, error) {
	filter := s.createEntityPlantingCycleFilter(req)
	plantingCycles, total, err := s.listPlantingCycleUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}

	var protoPlantingCycles []*plantingCycleP.PlantingCycle
	for _, plantingCycle := range plantingCycles {
		protoPlantingCycles = append(protoPlantingCycles, s.createProtoPlantingCycle(plantingCycle))
	}

	return &plantingCycleP.ListPlantingCyclesResponse{
		PlantingCycles: protoPlantingCycles,
		Total:          total,
	}, nil
}

func (s *PlantingCycleService) createEntityPlantingCycleFilter(req *plantingCycleP.ListPlantingCyclesRequest) *entity.PlantingCycleFilter {
	filter := &entity.PlantingCycleFilter{
		ID:             req.Id,
		CycleName:      req.CycleName,
		GrowingZoneID:  req.GrowingZoneId,
		PlantVarietyID: req.PlantVarietyId,
		Status:         req.Status,
		CreatedBy:      req.CreatedBy,
		Limit:          int(req.Limit),
		Offset:         int(req.Offset),
		SortBy:         req.SortBy,
		SortDirection:  req.SortDirection,
	}

	if req.SeedDateFrom != nil {
		seedDateFrom := req.SeedDateFrom.AsTime()
		filter.SeedDateFrom = &seedDateFrom
	}
	if req.SeedDateTo != nil {
		seedDateTo := req.SeedDateTo.AsTime()
		filter.SeedDateTo = &seedDateTo
	}
	if req.ExpectedHarvestFrom != nil {
		expectedHarvestFrom := req.ExpectedHarvestFrom.AsTime()
		filter.ExpectedHarvestFrom = &expectedHarvestFrom
	}
	if req.ExpectedHarvestTo != nil {
		expectedHarvestTo := req.ExpectedHarvestTo.AsTime()
		filter.ExpectedHarvestTo = &expectedHarvestTo
	}

	return filter
}
