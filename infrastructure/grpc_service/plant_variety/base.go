package plant_variety_service

import (
	"farm-service/domain/repository"
	"farm-service/domain/usecase/plant_variety"

	plantVarietyP "github.com/anhvanhoa/sf-proto/gen/plant_variety/v1"
)

type PlantVarietyService struct {
	plantVarietyP.UnsafePlantVarietyServiceServer
	createPlantVarietyUsecase      plant_variety.CreatePlantVarietyUsecase
	getPlantVarietyUsecase         plant_variety.GetPlantVarietyUsecase
	updatePlantVarietyUsecase      plant_variety.UpdatePlantVarietyUsecase
	deletePlantVarietyUsecase      plant_variety.DeletePlantVarietyUsecase
	listPlantVarietyUsecase        plant_variety.ListPlantVarietyUsecase
	getActivePlantVarietiesUsecase plant_variety.GetActivePlantVarietiesUsecase
	getByCategoryUsecase           plant_variety.GetByCategoryUsecase
	searchPlantVarietiesUsecase    plant_variety.SearchPlantVarietiesUsecase
	getByTemperatureRangeUsecase   plant_variety.GetByTemperatureRangeUsecase
	getByStatusUsecase             plant_variety.GetByStatusUsecase
	getBySeasonUsecase             plant_variety.GetBySeasonUsecase
	getByHumidityRangeUsecase      plant_variety.GetByHumidityRangeUsecase
	getByWaterRequirementUsecase   plant_variety.GetByWaterRequirementUsecase
	getByLightRequirementUsecase   plant_variety.GetByLightRequirementUsecase
}

func NewPlantVarietyService(plantVarietyRepository repository.PlantVarietyRepository) plantVarietyP.PlantVarietyServiceServer {
	return &PlantVarietyService{
		createPlantVarietyUsecase:      plant_variety.NewCreatePlantVarietyUsecase(plantVarietyRepository),
		getPlantVarietyUsecase:         plant_variety.NewGetPlantVarietyUsecase(plantVarietyRepository),
		updatePlantVarietyUsecase:      plant_variety.NewUpdatePlantVarietyUsecase(plantVarietyRepository),
		deletePlantVarietyUsecase:      plant_variety.NewDeletePlantVarietyUsecase(plantVarietyRepository),
		listPlantVarietyUsecase:        plant_variety.NewListPlantVarietyUsecase(plantVarietyRepository),
		getActivePlantVarietiesUsecase: plant_variety.NewGetActivePlantVarietiesUsecase(plantVarietyRepository),
		getByCategoryUsecase:           plant_variety.NewGetByCategoryUsecase(plantVarietyRepository),
		searchPlantVarietiesUsecase:    plant_variety.NewSearchPlantVarietiesUsecase(plantVarietyRepository),
		getByTemperatureRangeUsecase:   plant_variety.NewGetByTemperatureRangeUsecase(plantVarietyRepository),
		getByStatusUsecase:             plant_variety.NewGetByStatusUsecase(plantVarietyRepository),
		getBySeasonUsecase:             plant_variety.NewGetBySeasonUsecase(plantVarietyRepository),
		getByHumidityRangeUsecase:      plant_variety.NewGetByHumidityRangeUsecase(plantVarietyRepository),
		getByWaterRequirementUsecase:   plant_variety.NewGetByWaterRequirementUsecase(plantVarietyRepository),
		getByLightRequirementUsecase:   plant_variety.NewGetByLightRequirementUsecase(plantVarietyRepository),
	}
}
