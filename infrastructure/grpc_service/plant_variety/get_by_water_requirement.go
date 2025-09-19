package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) GetPlantVarietiesByWaterRequirement(ctx context.Context, req *plantVarietyP.GetPlantVarietiesByWaterRequirementRequest) (*plantVarietyP.ListPlantVarietiesResponse, error) {
	plantVarieties, err := s.getByWaterRequirementUsecase.Execute(ctx, req.WaterRequirement)
	if err != nil {
		return nil, err
	}

	var protoPlantVarieties []*plantVarietyP.PlantVariety
	for _, plantVariety := range plantVarieties {
		protoPlantVarieties = append(protoPlantVarieties, s.createProtoPlantVariety(plantVariety))
	}

	return &plantVarietyP.ListPlantVarietiesResponse{
		PlantVarieties: protoPlantVarieties,
		Total:          int64(len(protoPlantVarieties)),
	}, nil
}
