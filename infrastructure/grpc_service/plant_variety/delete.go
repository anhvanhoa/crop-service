package plant_variety_service

import (
	"context"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlantVarietyService) DeletePlantVariety(ctx context.Context, req *plantVarietyP.DeletePlantVarietyRequest) (*emptypb.Empty, error) {
	err := s.deletePlantVarietyUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
