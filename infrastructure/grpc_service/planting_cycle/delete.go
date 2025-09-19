package planting_cycle_service

import (
	"context"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *PlantingCycleService) DeletePlantingCycle(ctx context.Context, req *plantingCycleP.DeletePlantingCycleRequest) (*emptypb.Empty, error) {
	err := s.deletePlantingCycleUsecase.Execute(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
