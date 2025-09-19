package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) GetPlantVariety(ctx context.Context, req *plantVarietyP.GetPlantVarietyRequest) (*plantVarietyP.PlantVarietyResponse, error) {
	plantVariety, err := s.getPlantVarietyUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	if plantVariety == nil {
		return &plantVarietyP.PlantVarietyResponse{}, nil
	}
	return &plantVarietyP.PlantVarietyResponse{
		PlantVariety: s.createProtoPlantVariety(plantVariety),
	}, nil
}
