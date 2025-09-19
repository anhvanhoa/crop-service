package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlantVarietyService) GetActivePlantVarieties(ctx context.Context, req *emptypb.Empty) (*plantVarietyP.ListPlantVarietiesResponse, error) {
	plantVarieties, err := s.getActivePlantVarietiesUsecase.Execute(ctx)
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
