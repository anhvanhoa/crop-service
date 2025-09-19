package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) SearchPlantVarieties(ctx context.Context, req *plantVarietyP.SearchPlantVarietiesRequest) (*plantVarietyP.ListPlantVarietiesResponse, error) {
	plantVarieties, err := s.searchPlantVarietiesUsecase.Execute(ctx, req.Name)
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
