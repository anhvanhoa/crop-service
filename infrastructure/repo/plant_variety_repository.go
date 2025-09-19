package repo

import (
	"context"
	"fmt"

	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

// plantVarietyRepository implements repository.PlantVarietyRepository
type plantVarietyRepository struct {
	db *pg.DB
}

// NewPlantVarietyRepository creates a new instance of PlantVarietyRepository
func NewPlantVarietyRepository(db *pg.DB) repository.PlantVarietyRepository {
	return &plantVarietyRepository{
		db: db,
	}
}

// Create creates a new plant variety
func (r *plantVarietyRepository) Create(ctx context.Context, plantVariety *entity.PlantVariety) error {
	_, err := r.db.Model(plantVariety).Context(ctx).Insert()
	return err
}

// GetByID retrieves a plant variety by ID
func (r *plantVarietyRepository) GetByID(ctx context.Context, id string) (*entity.PlantVariety, error) {
	plantVariety := &entity.PlantVariety{}
	err := r.db.Model(plantVariety).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return plantVariety, nil
}

// Update updates an existing plant variety
func (r *plantVarietyRepository) Update(ctx context.Context, plantVariety *entity.PlantVariety) error {
	_, err := r.db.Model(plantVariety).Context(ctx).Where("id = ?", plantVariety.ID).UpdateNotZero()
	return err
}

// Delete removes a plant variety by ID
func (r *plantVarietyRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.PlantVariety)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

// List retrieves plant varieties with filtering and pagination
func (r *plantVarietyRepository) List(ctx context.Context, filter *entity.PlantVarietyFilter) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety

	query := r.db.Model(&plantVarieties).Context(ctx)

	// Apply filters
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.GrowingSeason != "" {
		query = query.Where("growing_season = ?", filter.GrowingSeason)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}

	// Apply sorting
	if filter.SortBy != "" {
		direction := "ASC"
		if filter.SortDirection == "desc" {
			direction = "DESC"
		}
		query = query.Order(fmt.Sprintf("%s %s", filter.SortBy, direction))
	} else {
		query = query.Order("created_at DESC")
	}

	// Apply pagination
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.Select()
	return plantVarieties, err
}

// GetByCategory retrieves plant varieties by category
func (r *plantVarietyRepository) GetByCategory(ctx context.Context, category string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("category = ?", category).Select()
	return plantVarieties, err
}

// GetByGrowingSeason retrieves plant varieties by growing season
func (r *plantVarietyRepository) GetByGrowingSeason(ctx context.Context, season string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("growing_season = ?", season).Select()
	return plantVarieties, err
}

// GetByStatus retrieves plant varieties by status
func (r *plantVarietyRepository) GetByStatus(ctx context.Context, status string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("status = ?", status).Select()
	return plantVarieties, err
}

// GetByCreatedBy retrieves plant varieties created by a specific user
func (r *plantVarietyRepository) GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("created_by = ?", createdBy).Select()
	return plantVarieties, err
}

// SearchByName searches plant varieties by name (partial match)
func (r *plantVarietyRepository) SearchByName(ctx context.Context, name string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("name ILIKE ?", "%"+name+"%").Select()
	return plantVarieties, err
}

// Count returns the total number of plant varieties matching the filter
func (r *plantVarietyRepository) Count(ctx context.Context, filter *entity.PlantVarietyFilter) (int64, error) {
	query := r.db.Model((*entity.PlantVariety)(nil)).Context(ctx)

	// Apply filters
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.Name != "" {
		query = query.Where("name ILIKE ?", "%"+filter.Name+"%")
	}
	if filter.Category != "" {
		query = query.Where("category = ?", filter.Category)
	}
	if filter.GrowingSeason != "" {
		query = query.Where("growing_season = ?", filter.GrowingSeason)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}

	count, err := query.Count()
	return int64(count), err
}

// Exists checks if a plant variety exists by ID
func (r *plantVarietyRepository) Exists(ctx context.Context, id string) (bool, error) {
	count, err := r.db.Model((*entity.PlantVariety)(nil)).Context(ctx).Where("id = ?", id).Count()
	return count > 0, err
}

// GetActiveVarieties retrieves all active plant varieties
func (r *plantVarietyRepository) GetActiveVarieties(ctx context.Context) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("status = ?", "active").Select()
	return plantVarieties, err
}

// GetVarietiesByTemperatureRange retrieves plant varieties suitable for a temperature range
func (r *plantVarietyRepository) GetVarietiesByTemperatureRange(ctx context.Context, minTemp, maxTemp float64) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).
		Where("optimal_temp_min >= ? AND optimal_temp_max <= ?", minTemp, maxTemp).
		Select()
	return plantVarieties, err
}

// GetVarietiesByHumidityRange retrieves plant varieties suitable for a humidity range
func (r *plantVarietyRepository) GetVarietiesByHumidityRange(ctx context.Context, minHumidity, maxHumidity float64) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).
		Where("optimal_humidity_min <= ? AND optimal_humidity_max >= ?", maxHumidity, minHumidity).
		Select()
	return plantVarieties, err
}

// GetVarietiesByWaterRequirement retrieves plant varieties by water requirement
func (r *plantVarietyRepository) GetVarietiesByWaterRequirement(ctx context.Context, waterRequirement string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("water_requirement = ?", waterRequirement).Select()
	return plantVarieties, err
}

// GetVarietiesByLightRequirement retrieves plant varieties by light requirement
func (r *plantVarietyRepository) GetVarietiesByLightRequirement(ctx context.Context, lightRequirement string) ([]*entity.PlantVariety, error) {
	var plantVarieties []*entity.PlantVariety
	err := r.db.Model(&plantVarieties).Context(ctx).Where("light_requirement = ?", lightRequirement).Select()
	return plantVarieties, err
}
