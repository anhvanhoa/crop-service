package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) GetPlantVarietiesByTemperatureRange(ctx context.Context, req *plantVarietyP.GetPlantVarietiesByTemperatureRangeRequest) (*plantVarietyP.ListPlantVarietiesResponse, error) {
	plantVarieties, err := s.getByTemperatureRangeUsecase.Execute(ctx, req.MinTemp, req.MaxTemp)
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
