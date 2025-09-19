package plant_variety_service

import (
	"context"

	"farm-service/domain/entity"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *PlantVarietyService) CreatePlantVariety(ctx context.Context, req *plantVarietyP.CreatePlantVarietyRequest) (*plantVarietyP.PlantVarietyResponse, error) {
	plantVarietyReq, err := s.createEntityPlantVarietyReq(req)
	if err != nil {
		return nil, err
	}
	plantVariety, err := s.createPlantVarietyUsecase.Execute(ctx, plantVarietyReq)
	if err != nil {
		return nil, err
	}
	return &plantVarietyP.PlantVarietyResponse{
		PlantVariety: s.createProtoPlantVariety(plantVariety),
	}, nil
}

func (s *PlantVarietyService) createEntityPlantVarietyReq(req *plantVarietyP.CreatePlantVarietyRequest) (*entity.CreatePlantVarietyRequest, error) {
	plantVariety := &entity.CreatePlantVarietyRequest{
		Name:               req.Name,
		ScientificName:     req.ScientificName,
		Category:           req.Category,
		GrowingSeason:      req.GrowingSeason,
		GrowthDurationDays: int(req.GrowthDurationDays),
		OptimalTempMin:     req.OptimalTempMin,
		OptimalTempMax:     req.OptimalTempMax,
		OptimalHumidityMin: req.OptimalHumidityMin,
		OptimalHumidityMax: req.OptimalHumidityMax,
		PHMin:              req.PhMin,
		PHMax:              req.PhMax,
		WaterRequirement:   req.WaterRequirement,
		LightRequirement:   req.LightRequirement,
		Description:        req.Description,
		MediaID:            req.MediaId,
		CreatedBy:          req.CreatedBy,
	}
	return plantVariety, nil
}

func (s *PlantVarietyService) createProtoPlantVariety(plantVariety *entity.PlantVariety) *plantVarietyP.PlantVariety {
	response := &plantVarietyP.PlantVariety{
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

	if plantVariety.UpdatedAt != nil {
		response.UpdatedAt = timestamppb.New(*plantVariety.UpdatedAt)
	}

	return response
}
