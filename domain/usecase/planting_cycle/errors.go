package planting_cycle

import (
	"github.com/anhvanhoa/service-core/domain/oops"
)

// Common errors for planting cycle use cases
var (
	ErrPlantingCycleNotFound = oops.New("Planting cycle not found")
	ErrInvalidPlantingCycle  = oops.New("Invalid planting cycle data")
	ErrPlantingCycleExists   = oops.New("Planting cycle already exists")
	ErrInvalidStatus         = oops.New("Invalid status")
	ErrInvalidDateRange      = oops.New("Invalid date range")
)
