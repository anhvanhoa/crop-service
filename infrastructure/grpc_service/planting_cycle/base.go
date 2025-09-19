package planting_cycle_service

import (
	"farm-service/domain/repository"
	"farm-service/domain/usecase/planting_cycle"

	plantingCycleP "github.com/anhvanhoa/sf-proto/gen/planting_cycle/v1"
)

type PlantingCycleService struct {
	plantingCycleP.UnsafePlantingCycleServiceServer
	createPlantingCycleUsecase     planting_cycle.CreatePlantingCycleUsecase
	getPlantingCycleUsecase        planting_cycle.GetPlantingCycleUsecase
	updatePlantingCycleUsecase     planting_cycle.UpdatePlantingCycleUsecase
	deletePlantingCycleUsecase     planting_cycle.DeletePlantingCycleUsecase
	listPlantingCycleUsecase       planting_cycle.ListPlantingCycleUsecase
	getCyclesByZoneUsecase         planting_cycle.GetCyclesByZoneUsecase
	getCyclesByVarietyUsecase      planting_cycle.GetCyclesByVarietyUsecase
	getActivePlantingCyclesUsecase planting_cycle.GetActivePlantingCyclesUsecase
	getOverdueHarvestsUsecase      planting_cycle.GetOverdueHarvestsUsecase
	getUpcomingHarvestsUsecase     planting_cycle.GetUpcomingHarvestsUsecase
	getByDateRangeUsecase          planting_cycle.GetByDateRangeUsecase
	getByHarvestDateRangeUsecase   planting_cycle.GetByHarvestDateRangeUsecase
	getBySeedDateRangeUsecase      planting_cycle.GetBySeedDateRangeUsecase
	getByStatusUsecase             planting_cycle.GetByStatusUsecase
	getCyclesWithDetailsUsecase    planting_cycle.GetCyclesWithDetailsUsecase
	getCycleWithDetailsUsecase     planting_cycle.GetCycleWithDetailsUsecase
	updateHarvestDateUsecase       planting_cycle.UpdateHarvestDateUsecase
	updateStatusUsecase            planting_cycle.UpdateStatusUsecase
}

func NewPlantingCycleService(plantingCycleRepository repository.PlantingCycleRepository) plantingCycleP.PlantingCycleServiceServer {
	return &PlantingCycleService{
		createPlantingCycleUsecase:     planting_cycle.NewCreatePlantingCycleUsecase(plantingCycleRepository),
		getPlantingCycleUsecase:        planting_cycle.NewGetPlantingCycleUsecase(plantingCycleRepository),
		updatePlantingCycleUsecase:     planting_cycle.NewUpdatePlantingCycleUsecase(plantingCycleRepository),
		deletePlantingCycleUsecase:     planting_cycle.NewDeletePlantingCycleUsecase(plantingCycleRepository),
		listPlantingCycleUsecase:       planting_cycle.NewListPlantingCycleUsecase(plantingCycleRepository),
		getCyclesByZoneUsecase:         planting_cycle.NewGetCyclesByZoneUsecase(plantingCycleRepository),
		getCyclesByVarietyUsecase:      planting_cycle.NewGetCyclesByVarietyUsecase(plantingCycleRepository),
		getActivePlantingCyclesUsecase: planting_cycle.NewGetActivePlantingCyclesUsecase(plantingCycleRepository),
		getOverdueHarvestsUsecase:      planting_cycle.NewGetOverdueHarvestsUsecase(plantingCycleRepository),
		getUpcomingHarvestsUsecase:     planting_cycle.NewGetUpcomingHarvestsUsecase(plantingCycleRepository),
		getByDateRangeUsecase:          planting_cycle.NewGetByDateRangeUsecase(plantingCycleRepository),
		getByHarvestDateRangeUsecase:   planting_cycle.NewGetByHarvestDateRangeUsecase(plantingCycleRepository),
		getBySeedDateRangeUsecase:      planting_cycle.NewGetBySeedDateRangeUsecase(plantingCycleRepository),
		getByStatusUsecase:             planting_cycle.NewGetByStatusUsecase(plantingCycleRepository),
		getCyclesWithDetailsUsecase:    planting_cycle.NewGetCyclesWithDetailsUsecase(plantingCycleRepository),
		getCycleWithDetailsUsecase:     planting_cycle.NewGetCycleWithDetailsUsecase(plantingCycleRepository),
		updateHarvestDateUsecase:       planting_cycle.NewUpdateHarvestDateUsecase(plantingCycleRepository),
		updateStatusUsecase:            planting_cycle.NewUpdateStatusUsecase(plantingCycleRepository),
	}
}
