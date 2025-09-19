package repo

import (
	"context"
	"fmt"
	"time"

	"farm-service/domain/entity"
	"farm-service/domain/repository"

	"github.com/go-pg/pg/v10"
)

// plantingCycleRepository implements repository.PlantingCycleRepository
type plantingCycleRepository struct {
	db *pg.DB
}

// NewPlantingCycleRepository creates a new instance of PlantingCycleRepository
func NewPlantingCycleRepository(db *pg.DB) repository.PlantingCycleRepository {
	return &plantingCycleRepository{
		db: db,
	}
}

// Create creates a new planting cycle
func (r *plantingCycleRepository) Create(ctx context.Context, plantingCycle *entity.PlantingCycle) error {
	_, err := r.db.Model(plantingCycle).Context(ctx).Insert()
	return err
}

// GetByID retrieves a planting cycle by ID
func (r *plantingCycleRepository) GetByID(ctx context.Context, id string) (*entity.PlantingCycle, error) {
	plantingCycle := &entity.PlantingCycle{}
	err := r.db.Model(plantingCycle).Context(ctx).Where("id = ?", id).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return plantingCycle, nil
}

// Update updates an existing planting cycle
func (r *plantingCycleRepository) Update(ctx context.Context, plantingCycle *entity.PlantingCycle) error {
	_, err := r.db.Model(plantingCycle).Context(ctx).Where("id = ?", plantingCycle.ID).UpdateNotZero()
	return err
}

// Delete removes a planting cycle by ID
func (r *plantingCycleRepository) Delete(ctx context.Context, id string) error {
	_, err := r.db.Model((*entity.PlantingCycle)(nil)).Context(ctx).Where("id = ?", id).Delete()
	return err
}

// List retrieves planting cycles with filtering and pagination
func (r *plantingCycleRepository) List(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle

	query := r.db.Model(&plantingCycles).Context(ctx)

	// Apply filters
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.CycleName != "" {
		query = query.Where("cycle_name ILIKE ?", "%"+filter.CycleName+"%")
	}
	if filter.GrowingZoneID != "" {
		query = query.Where("growing_zone_id = ?", filter.GrowingZoneID)
	}
	if filter.PlantVarietyID != "" {
		query = query.Where("plant_variety_id = ?", filter.PlantVarietyID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.SeedDateFrom != nil {
		query = query.Where("seed_date >= ?", *filter.SeedDateFrom)
	}
	if filter.SeedDateTo != nil {
		query = query.Where("seed_date <= ?", *filter.SeedDateTo)
	}
	if filter.ExpectedHarvestFrom != nil {
		query = query.Where("expected_harvest_date >= ?", *filter.ExpectedHarvestFrom)
	}
	if filter.ExpectedHarvestTo != nil {
		query = query.Where("expected_harvest_date <= ?", *filter.ExpectedHarvestTo)
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
	return plantingCycles, err
}

// GetByGrowingZoneID retrieves all planting cycles for a specific growing zone
func (r *plantingCycleRepository) GetByGrowingZoneID(ctx context.Context, growingZoneID string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("growing_zone_id = ?", growingZoneID).Select()
	return plantingCycles, err
}

// GetByPlantVarietyID retrieves all planting cycles for a specific plant variety
func (r *plantingCycleRepository) GetByPlantVarietyID(ctx context.Context, plantVarietyID string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("plant_variety_id = ?", plantVarietyID).Select()
	return plantingCycles, err
}

// GetByStatus retrieves planting cycles by status
func (r *plantingCycleRepository) GetByStatus(ctx context.Context, status string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("status = ?", status).Select()
	return plantingCycles, err
}

// GetByCreatedBy retrieves planting cycles created by a specific user
func (r *plantingCycleRepository) GetByCreatedBy(ctx context.Context, createdBy string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("created_by = ?", createdBy).Select()
	return plantingCycles, err
}

// GetActiveCycles retrieves all active planting cycles
func (r *plantingCycleRepository) GetActiveCycles(ctx context.Context) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	activeStatuses := []string{"planning", "seeding", "transplanting", "growing", "flowering", "harvesting"}
	err := r.db.Model(&plantingCycles).Context(ctx).Where("status IN (?)", pg.In(activeStatuses)).Select()
	return plantingCycles, err
}

// GetCyclesByDateRange retrieves planting cycles within a date range
func (r *plantingCycleRepository) GetCyclesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("created_at >= ? AND created_at <= ?", startDate, endDate).
		Select()
	return plantingCycles, err
}

// GetCyclesBySeedDateRange retrieves planting cycles by seed date range
func (r *plantingCycleRepository) GetCyclesBySeedDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("seed_date >= ? AND seed_date <= ?", startDate, endDate).
		Select()
	return plantingCycles, err
}

// GetCyclesByHarvestDateRange retrieves planting cycles by expected harvest date range
func (r *plantingCycleRepository) GetCyclesByHarvestDateRange(ctx context.Context, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("expected_harvest_date >= ? AND expected_harvest_date <= ?", startDate, endDate).
		Select()
	return plantingCycles, err
}

// GetUpcomingHarvests retrieves planting cycles with upcoming harvests
func (r *plantingCycleRepository) GetUpcomingHarvests(ctx context.Context, days int) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	cutoffDate := time.Now().AddDate(0, 0, days)
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("expected_harvest_date <= ? AND expected_harvest_date >= ? AND actual_harvest_date IS NULL", cutoffDate, time.Now()).
		Select()
	return plantingCycles, err
}

// GetOverdueHarvests retrieves planting cycles with overdue harvests
func (r *plantingCycleRepository) GetOverdueHarvests(ctx context.Context) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("expected_harvest_date < ? AND actual_harvest_date IS NULL", time.Now()).
		Select()
	return plantingCycles, err
}

// GetCyclesByStatusAndDate retrieves planting cycles by status and date range
func (r *plantingCycleRepository) GetCyclesByStatusAndDate(ctx context.Context, status string, startDate, endDate time.Time) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("status = ? AND created_at >= ? AND created_at <= ?", status, startDate, endDate).
		Select()
	return plantingCycles, err
}

// Count returns the total number of planting cycles matching the filter
func (r *plantingCycleRepository) Count(ctx context.Context, filter *entity.PlantingCycleFilter) (int64, error) {
	query := r.db.Model((*entity.PlantingCycle)(nil)).Context(ctx)

	// Apply filters (same as List method)
	if filter.ID != "" {
		query = query.Where("id = ?", filter.ID)
	}
	if filter.CycleName != "" {
		query = query.Where("cycle_name ILIKE ?", "%"+filter.CycleName+"%")
	}
	if filter.GrowingZoneID != "" {
		query = query.Where("growing_zone_id = ?", filter.GrowingZoneID)
	}
	if filter.PlantVarietyID != "" {
		query = query.Where("plant_variety_id = ?", filter.PlantVarietyID)
	}
	if filter.Status != "" {
		query = query.Where("status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("created_by = ?", filter.CreatedBy)
	}
	if filter.SeedDateFrom != nil {
		query = query.Where("seed_date >= ?", *filter.SeedDateFrom)
	}
	if filter.SeedDateTo != nil {
		query = query.Where("seed_date <= ?", *filter.SeedDateTo)
	}
	if filter.ExpectedHarvestFrom != nil {
		query = query.Where("expected_harvest_date >= ?", *filter.ExpectedHarvestFrom)
	}
	if filter.ExpectedHarvestTo != nil {
		query = query.Where("expected_harvest_date <= ?", *filter.ExpectedHarvestTo)
	}

	count, err := query.Count()
	return int64(count), err
}

// Exists checks if a planting cycle exists by ID
func (r *plantingCycleRepository) Exists(ctx context.Context, id string) (bool, error) {
	count, err := r.db.Model((*entity.PlantingCycle)(nil)).Context(ctx).Where("id = ?", id).Count()
	return count > 0, err
}

// GetCyclesWithDetails retrieves planting cycles with related plant variety and growing zone information
func (r *plantingCycleRepository) GetCyclesWithDetails(ctx context.Context, filter *entity.PlantingCycleFilter) ([]*entity.PlantingCycleWithDetails, error) {
	var cyclesWithDetails []*entity.PlantingCycleWithDetails

	query := r.db.Model(&cyclesWithDetails).Context(ctx).
		Relation("PlantVariety")

	// Apply same filters as List method
	if filter.ID != "" {
		query = query.Where("planting_cycle.id = ?", filter.ID)
	}
	if filter.CycleName != "" {
		query = query.Where("planting_cycle.cycle_name ILIKE ?", "%"+filter.CycleName+"%")
	}
	if filter.GrowingZoneID != "" {
		query = query.Where("planting_cycle.growing_zone_id = ?", filter.GrowingZoneID)
	}
	if filter.PlantVarietyID != "" {
		query = query.Where("planting_cycle.plant_variety_id = ?", filter.PlantVarietyID)
	}
	if filter.Status != "" {
		query = query.Where("planting_cycle.status = ?", filter.Status)
	}
	if filter.CreatedBy != "" {
		query = query.Where("planting_cycle.created_by = ?", filter.CreatedBy)
	}

	// Apply sorting
	if filter.SortBy != "" {
		direction := "ASC"
		if filter.SortDirection == "desc" {
			direction = "DESC"
		}
		query = query.Order(fmt.Sprintf("planting_cycle.%s %s", filter.SortBy, direction))
	} else {
		query = query.Order("planting_cycle.created_at DESC")
	}

	// Apply pagination
	if filter.Limit > 0 {
		query = query.Limit(filter.Limit)
	}
	if filter.Offset > 0 {
		query = query.Offset(filter.Offset)
	}

	err := query.Select()
	return cyclesWithDetails, err
}

// GetCycleWithDetails retrieves a single planting cycle with related information
func (r *plantingCycleRepository) GetCycleWithDetails(ctx context.Context, id string) (*entity.PlantingCycleWithDetails, error) {
	cycleWithDetails := &entity.PlantingCycleWithDetails{}
	err := r.db.Model(cycleWithDetails).Context(ctx).
		Relation("PlantVariety").
		Where("planting_cycle.id = ?", id).
		Select()
	if err != nil {
		if err == pg.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return cycleWithDetails, nil
}

// GetCyclesByPlantVarietyAndZone retrieves planting cycles by both plant variety and growing zone
func (r *plantingCycleRepository) GetCyclesByPlantVarietyAndZone(ctx context.Context, plantVarietyID, growingZoneID string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("plant_variety_id = ? AND growing_zone_id = ?", plantVarietyID, growingZoneID).
		Select()
	return plantingCycles, err
}

// GetCyclesByBatch retrieves planting cycles by seed batch
func (r *plantingCycleRepository) GetCyclesByBatch(ctx context.Context, seedBatch string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("seed_batch = ?", seedBatch).Select()
	return plantingCycles, err
}

// GetCyclesByQuantityRange retrieves planting cycles by plant quantity range
func (r *plantingCycleRepository) GetCyclesByQuantityRange(ctx context.Context, minQuantity, maxQuantity int) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).
		Where("plant_quantity >= ? AND plant_quantity <= ?", minQuantity, maxQuantity).
		Select()
	return plantingCycles, err
}

// UpdateStatus updates the status of a planting cycle
func (r *plantingCycleRepository) UpdateStatus(ctx context.Context, id, status string) error {
	_, err := r.db.Model((*entity.PlantingCycle)(nil)).Context(ctx).
		Set("status = ?", status).
		Where("id = ?", id).
		Update()
	return err
}

// UpdateHarvestDate updates the actual harvest date of a planting cycle
func (r *plantingCycleRepository) UpdateHarvestDate(ctx context.Context, id string, harvestDate time.Time) error {
	_, err := r.db.Model((*entity.PlantingCycle)(nil)).Context(ctx).
		Set("actual_harvest_date = ?", harvestDate).
		Where("id = ?", id).
		Update()
	return err
}

// GetCyclesByStatuses retrieves planting cycles by multiple statuses
func (r *plantingCycleRepository) GetCyclesByStatuses(ctx context.Context, statuses []string) ([]*entity.PlantingCycle, error) {
	var plantingCycles []*entity.PlantingCycle
	err := r.db.Model(&plantingCycles).Context(ctx).Where("status IN (?)", pg.In(statuses)).Select()
	return plantingCycles, err
}
