package repository

import (
	"context"

	"farm-service/domain/entity"
)

// PlantVarietyRepository defines the interface for plant variety data operations
type PlantVarietyRepository interface {
	// Create creates a new plant variety
	Create(ctx context.Context, plantVariety *entity.PlantVariety) error

	// GetByID retrieves a plant variety by ID
	GetByID(ctx context.Context, id string) (*entity.PlantVariety, error)

	// Update updates an existing plant variety
	Update(ctx context.Context, plantVariety *entity.PlantVariety) error

	// Delete removes a plant variety by ID
	Delete(ctx context.Context, id string) error

	// List retrieves plant varieties with filtering and pagination
	List(ctx context.Context, filter *entity.PlantVarietyFilter) ([]*entity.PlantVariety, error)

	// GetByCategory retrieves plant varieties by category
	GetByCategory(ctx context.Context, category string) ([]*entity.PlantVariety, error)

	// GetByGrowingSeason retrieves plant varieties by growing season
	GetByGrowingSeason(ctx context.Context, season string) ([]*entity.PlantVariety, error)

	// GetByStatus retrieves plant varieties by status
	GetByStatus(ctx context.Context, status string) ([]*entity.PlantVariety, error)

	// GetByCreatedBy retrieves plant varieties created by a specific user
	GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.PlantVariety, error)

	// SearchByName searches plant varieties by name (partial match)
	SearchByName(ctx context.Context, name string) ([]*entity.PlantVariety, error)

	// Count returns the total number of plant varieties matching the filter
	Count(ctx context.Context, filter *entity.PlantVarietyFilter) (int64, error)

	// Exists checks if a plant variety exists by ID
	Exists(ctx context.Context, id string) (bool, error)

	// GetActiveVarieties retrieves all active plant varieties
	GetActiveVarieties(ctx context.Context) ([]*entity.PlantVariety, error)

	// GetVarietiesByTemperatureRange retrieves plant varieties suitable for a temperature range
	GetVarietiesByTemperatureRange(ctx context.Context, minTemp, maxTemp float64) ([]*entity.PlantVariety, error)

	// GetVarietiesByHumidityRange retrieves plant varieties suitable for a humidity range
	GetVarietiesByHumidityRange(ctx context.Context, minHumidity, maxHumidity float64) ([]*entity.PlantVariety, error)

	// GetVarietiesByWaterRequirement retrieves plant varieties by water requirement
	GetVarietiesByWaterRequirement(ctx context.Context, waterRequirement string) ([]*entity.PlantVariety, error)

	// GetVarietiesByLightRequirement retrieves plant varieties by light requirement
	GetVarietiesByLightRequirement(ctx context.Context, lightRequirement string) ([]*entity.PlantVariety, error)
}
