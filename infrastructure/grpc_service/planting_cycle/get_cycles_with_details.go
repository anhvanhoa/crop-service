package planting_cycle_service

import (
	"context"

	"farm-service/domain/entity"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

func (s *PlantingCycleService) GetPlantingCyclesWithDetails(ctx context.Context, req *plantingCycleP.GetPlantingCyclesWithDetailsRequest) (*plantingCycleP.ListPlantingCyclesWithDetailsResponse, error) {
	// Create filter from request
	filter := &entity.PlantingCycleFilter{
		Limit:         int(req.Limit),
		Offset:        int(req.Offset),
		SortBy:        req.SortBy,
		SortDirection: req.SortDirection,
	}

	// Set optional fields
	if req.Id != "" {
		filter.ID = req.Id
	}
	if req.CycleName != "" {
		filter.CycleName = req.CycleName
	}
	if req.GrowingZoneId != "" {
		filter.GrowingZoneID = req.GrowingZoneId
	}
	if req.PlantVarietyId != "" {
		filter.PlantVarietyID = req.PlantVarietyId
	}
	if req.Status != "" {
		filter.Status = req.Status
	}
	if req.CreatedBy != "" {
		filter.CreatedBy = req.CreatedBy
	}

	plantingCycles, total, err := s.getCyclesWithDetailsUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}

	var protoPlantingCycles []*plantingCycleP.PlantingCycleWithDetails
	for _, plantingCycle := range plantingCycles {
		protoPlantingCycles = append(protoPlantingCycles, s.createProtoPlantingCycleWithDetails(plantingCycle))
	}

	return &plantingCycleP.ListPlantingCyclesWithDetailsResponse{
		PlantingCyclesWithDetails: protoPlantingCycles,
		Total:                     total,
	}, nil
}
