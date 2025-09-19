package plant_variety

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

var (
	ErrPlantVarietyNotFound = oops.New("Plant variety not found")
	ErrInvalidPlantVariety  = oops.New("Invalid plant variety data")
	ErrPlantVarietyExists   = oops.New("Plant variety already exists")
)
