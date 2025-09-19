package plant_variety_service

import (
	"context"

	"farm-service/domain/entity"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) UpdatePlantVariety(ctx context.Context, req *plantVarietyP.UpdatePlantVarietyRequest) (*plantVarietyP.PlantVarietyResponse, error) {
	plantVarietyReq, err := s.createEntityUpdatePlantVarietyReq(req)
	if err != nil {
		return nil, err
	}
	plantVariety, err := s.updatePlantVarietyUsecase.Execute(ctx, plantVarietyReq)
	if err != nil {
		return nil, err
	}
	return &plantVarietyP.PlantVarietyResponse{
		PlantVariety: s.createProtoPlantVariety(plantVariety),
	}, nil
}

func (s *PlantVarietyService) createEntityUpdatePlantVarietyReq(req *plantVarietyP.UpdatePlantVarietyRequest) (*entity.UpdatePlantVarietyRequest, error) {
	plantVariety := &entity.UpdatePlantVarietyRequest{
		ID:                 req.Id,
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
		Status:             req.Status,
	}
	return plantVariety, nil
}
