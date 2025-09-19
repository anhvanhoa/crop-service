package plant_variety_service

import (
	"context"

	"farm-service/domain/entity"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

func (s *PlantVarietyService) ListPlantVarieties(ctx context.Context, req *plantVarietyP.ListPlantVarietiesRequest) (*plantVarietyP.ListPlantVarietiesResponse, error) {
	filter := s.createEntityPlantVarietyFilter(req)
	plantVarieties, total, err := s.listPlantVarietyUsecase.Execute(ctx, filter)
	if err != nil {
		return nil, err
	}

	var protoPlantVarieties []*plantVarietyP.PlantVariety
	for _, plantVariety := range plantVarieties {
		protoPlantVarieties = append(protoPlantVarieties, s.createProtoPlantVariety(plantVariety))
	}

	return &plantVarietyP.ListPlantVarietiesResponse{
		PlantVarieties: protoPlantVarieties,
		Total:          total,
	}, nil
}

func (s *PlantVarietyService) createEntityPlantVarietyFilter(req *plantVarietyP.ListPlantVarietiesRequest) *entity.PlantVarietyFilter {
	return &entity.PlantVarietyFilter{
		ID:            req.Id,
		Name:          req.Name,
		Category:      req.Category,
		GrowingSeason: req.GrowingSeason,
		Status:        req.Status,
		CreatedBy:     req.CreatedBy,
		Limit:         int(req.Limit),
		Offset:        int(req.Offset),
		SortBy:        req.SortBy,
		SortDirection: req.SortDirection,
	}
}
