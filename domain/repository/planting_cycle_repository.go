package repository

import (
	"context"
	"time"

	"farm-service/domain/entity"
)

// PlantingCycleRepository defines the interface for planting cycle data operations
type PlantingCycleRepository interface {
	// Create creates a new planting cycle
	Create(ctx context.Context, plantingCycle *entity.PlantingCycle) error

	// GetByID retrieves a planting cycle by ID
	GetByID(ctx context.Context, id string) (*entity.PlantingCycle, error)

	// Update updates an existing planting cycle
	Update(ctx context.Context, plantingCycle *entity.PlantingCycle) error

	// Delete removes a planting cycle by ID
	Delete(ctx context.Context, id string) error

	// List retrieves planting cycles with filtering and pagination
	List(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycle, error)

	// GetByGrowingZoneID retrieves all planting cycles for a specific growing zone
	GetByGrowingZoneID(ctx context.Context, growingZoneID string) ([]*entity.PlantingCycle, error)

	// GetByPlantVarietyID retrieves all planting cycles for a specific plant variety
	GetByPlantVarietyID(ctx context.Context, plantVarietyID string) ([]*entity.PlantingCycle, error)

	// GetByStatus retrieves planting cycles by status
	GetByStatus(ctx context.Context, status string) ([]*entity.PlantingCycle, error)

	// GetByCreatedBy retrieves planting cycles created by a specific user
	GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.PlantingCycle, error)

	// GetActiveCycles retrieves all active planting cycles
	GetActiveCycles(ctx context.Context) ([]*entity.PlantingCycle, error)

	// GetCyclesByDateRange retrieves planting cycles within a date range
	GetCyclesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)

	// GetCyclesBySeedDateRange retrieves planting cycles by seed date range
	GetCyclesBySeedDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)

	// GetCyclesByHarvestDateRange retrieves planting cycles by expected harvest date range
	GetCyclesByHarvestDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)

	// GetUpcomingHarvests retrieves planting cycles with upcoming harvests
	GetUpcomingHarvests(ctx context.Context, days int) ([]*entity.PlantingCycle, error)

	// GetOverdueHarvests retrieves planting cycles with overdue harvests
	GetOverdueHarvests(ctx context.Context) ([]*entity.PlantingCycle, error)

	// GetCyclesByStatusAndDate retrieves planting cycles by status and date range
	GetCyclesByStatusAndDate(ctx context.Context, status string, startDate, endDate time.Time) ([]*entity.PlantingCycle, error)

	// Count returns the total number of planting cycles matching the filter
	Count(ctx context.Context, filter *entity.PlantingCycleFilter) (int64, error)

	// Exists checks if a planting cycle exists by ID
	Exists(ctx context.Context, id string) (bool, error)

	// GetCyclesWithDetails retrieves planting cycles with related plant variety and growing zone information
	GetCyclesWithDetails(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycleWithDetails, error)

	// GetCycleWithDetails retrieves a single planting cycle with related information
	GetCycleWithDetails(ctx context.Context, id string) (*entity.PlantingCycleWithDetails, error)

	// GetCyclesByPlantVarietyAndZone retrieves planting cycles by both plant variety and growing zone
	GetCyclesByPlantVarietyAndZone(ctx context.Context, plantVarietyID, growingZoneID string) ([]*entity.PlantingCycle, error)

	// GetCyclesByBatch retrieves planting cycles by seed batch
	GetCyclesByBatch(ctx context.Context, seedBatch string) ([]*entity.PlantingCycle, error)

	// GetCyclesByQuantityRange retrieves planting cycles by plant quantity range
	GetCyclesByQuantityRange(ctx context.Context, minQuantity, maxQuantity int) ([]*entity.PlantingCycle, error)

	// UpdateStatus updates the status of a planting cycle
	UpdateStatus(ctx context.Context, id, status string) error

	// UpdateHarvestDate updates the actual harvest date of a planting cycle
	UpdateHarvestDate(ctx context.Context, id string, harvestDate time.Time) error

	// GetCyclesByStatus retrieves planting cycles by multiple statuses
	GetCyclesByStatuses(ctx context.Context, statuses []string) ([]*entity.PlantingCycle, error)
}
