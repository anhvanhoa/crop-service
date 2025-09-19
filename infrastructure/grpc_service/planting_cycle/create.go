package planting_cycle_service

import (
	"context"

	"farm-service/domain/entity"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *PlantingCycleService) CreatePlantingCycle(ctx context.Context, req *plantingCycleP.CreatePlantingCycleRequest) (*plantingCycleP.PlantingCycleResponse, error) {
	plantingCycleReq, err := s.createEntityPlantingCycleReq(req)
	if err != nil {
		return nil, err
	}
	plantingCycle, err := s.createPlantingCycleUsecase.Execute(ctx, plantingCycleReq)
	if err != nil {
		return nil, err
	}
	return &plantingCycleP.PlantingCycleResponse{
		PlantingCycle: s.createProtoPlantingCycle(plantingCycle),
	}, nil
}

func (s *PlantingCycleService) createEntityPlantingCycleReq(req *plantingCycleP.CreatePlantingCycleRequest) (*entity.CreatePlantingCycleRequest, error) {
	plantingCycle := &entity.CreatePlantingCycleRequest{
		CycleName:      req.CycleName,
		GrowingZoneID:  req.GrowingZoneId,
		PlantVarietyID: req.PlantVarietyId,
		PlantQuantity:  int(req.PlantQuantity),
		SeedBatch:      req.SeedBatch,
		Status:         req.Status,
		Notes:          req.Notes,
		CreatedBy:      req.CreatedBy,
	}

	// Convert timestamps
	if req.SeedDate != nil {
		seedDate := req.SeedDate.AsTime()
		plantingCycle.SeedDate = &seedDate
	}
	if req.TransplantDate != nil {
		transplantDate := req.TransplantDate.AsTime()
		plantingCycle.TransplantDate = &transplantDate
	}
	if req.ExpectedHarvestDate != nil {
		expectedHarvestDate := req.ExpectedHarvestDate.AsTime()
		plantingCycle.ExpectedHarvestDate = &expectedHarvestDate
	}
	if req.ActualHarvestDate != nil {
		actualHarvestDate := req.ActualHarvestDate.AsTime()
		plantingCycle.ActualHarvestDate = &actualHarvestDate
	}

	return plantingCycle, nil
}

func (s *PlantingCycleService) createProtoPlantingCycle(plantingCycle *entity.PlantingCycle) *plantingCycleP.PlantingCycle {
	response := &plantingCycleP.PlantingCycle{
		Id:             plantingCycle.ID,
		CycleName:      plantingCycle.CycleName,
		GrowingZoneId:  plantingCycle.GrowingZoneID,
		PlantVarietyId: plantingCycle.PlantVarietyID,
		PlantQuantity:  int32(plantingCycle.PlantQuantity),
		SeedBatch:      plantingCycle.SeedBatch,
		Status:         plantingCycle.Status,
		Notes:          plantingCycle.Notes,
		CreatedBy:      plantingCycle.CreatedBy,
		CreatedAt:      timestamppb.New(plantingCycle.CreatedAt),
	}

	// Convert timestamps
	if plantingCycle.SeedDate != nil {
		response.SeedDate = timestamppb.New(*plantingCycle.SeedDate)
	}
	if plantingCycle.TransplantDate != nil {
		response.TransplantDate = timestamppb.New(*plantingCycle.TransplantDate)
	}
	if plantingCycle.ExpectedHarvestDate != nil {
		response.ExpectedHarvestDate = timestamppb.New(*plantingCycle.ExpectedHarvestDate)
	}
	if plantingCycle.ActualHarvestDate != nil {
		response.ActualHarvestDate = timestamppb.New(*plantingCycle.ActualHarvestDate)
	}
	if plantingCycle.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*plantingCycle.UpdatedAt)
	}

	return response
}

func (s *PlantingCycleService) createProtoPlantingCycleWithDetails(plantingCycle *entity.PlantingCycleWithDetails) *plantingCycleP.PlantingCycleWithDetails {
	// Create the basic planting cycle proto
	plantingCycleProto := s.createProtoPlantingCycle(&plantingCycle.PlantingCycle)

	// Create the response with details
	response := &plantingCycleP.PlantingCycleWithDetails{
		PlantingCycle: plantingCycleProto,
	}

	// Add plant variety details if available
	if plantingCycle.PlantVariety != nil {
		response.PlantVariety = s.createProtoPlantVariety(plantingCycle.PlantVariety)
	}

	return response
}

func (s *PlantingCycleService) createProtoPlantVariety(plantVariety *entity.PlantVariety) *plantingCycleP.PlantVariety {
	response := &plantingCycleP.PlantVariety{
		Id:                 plantVariety.ID,
		Name:               plantVariety.Name,
		ScientificName:     plantVariety.ScientificName,
		Category:           plantVariety.Category,
		GrowingSeason:      plantVariety.GrowingSeason,
		GrowthDurationDays: int32(plantVariety.GrowthDurationDays),
		OptimalTempMin:     plantVariety.OptimalTempMin,
		OptimalTempMax:     plantVariety.OptimalTempMax,
		OptimalHumidityMin: plantVariety.OptimalHumidityMin,
		OptimalHumidityMax: plantVariety.OptimalHumidityMax,
		PhMin:              plantVariety.PHMin,
		PhMax:              plantVariety.PHMax,
		WaterRequirement:   plantVariety.WaterRequirement,
		LightRequirement:   plantVariety.LightRequirement,
		Description:        plantVariety.Description,
		MediaId:            plantVariety.MediaID,
		Status:             plantVariety.Status,
		CreatedBy:          plantVariety.CreatedBy,
		CreatedAt:          timestamppb.New(plantVariety.CreatedAt),
	}

	// Convert optional timestamp
	if plantVariety.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*plantVariety.UpdatedAt)
	}

	return response
}
