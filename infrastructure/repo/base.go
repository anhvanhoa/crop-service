package repo

import (
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

type Repository struct {
	PlantVarietyRepository  repository.PlantVarietyRepository
	PlantingCycleRepository repository.PlantingCycleRepository
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{
		PlantVarietyRepository:  NewPlantVarietyRepository(db),
		PlantingCycleRepository: NewPlantingCycleRepository(db),
	}
}
