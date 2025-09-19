package entity

import (
	"time"
)

// PlantingCycle represents a planting cycle entity
type PlantingCycle struct {
	ID                  string
	CycleName           string
	GrowingZoneID       string
	PlantVarietyID      string
	SeedDate            *time.Time
	TransplantDate      *time.Time
	ExpectedHarvestDate *time.Time
	ActualHarvestDate   *time.Time
	PlantQuantity       int
	SeedBatch           string
	Status              string
	Notes               string
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

// CreatePlantingCycleRequest represents request to create a new planting cycle
type CreatePlantingCycleRequest struct {
	CycleName           string
	GrowingZoneID       string
	PlantVarietyID      string
	SeedDate            *time.Time
	TransplantDate      *time.Time
	ExpectedHarvestDate *time.Time
	ActualHarvestDate   *time.Time
	PlantQuantity       int
	SeedBatch           string
	Status              string
	Notes               string
	CreatedBy           string
}

// UpdatePlantingCycleRequest represents request to update a planting cycle
type UpdatePlantingCycleRequest struct {
	ID                  string
	CycleName           string
	GrowingZoneID       string
	PlantVarietyID      string
	SeedDate            *time.Time
	TransplantDate      *time.Time
	ExpectedHarvestDate *time.Time
	ActualHarvestDate   *time.Time
	PlantQuantity       int
	SeedBatch           string
	Status              string
	Notes               string
}

// PlantingCycleFilter represents filter for searching planting cycles
type PlantingCycleFilter struct {
	ID                  string
	CycleName           string
	GrowingZoneID       string
	PlantVarietyID      string
	Status              string
	CreatedBy           string
	SeedDateFrom        *time.Time
	SeedDateTo          *time.Time
	ExpectedHarvestFrom *time.Time
	ExpectedHarvestTo   *time.Time
	Limit               int
	Offset              int
	SortBy              string
	SortDirection       string
}

// PlantingCycleResponse represents response for planting cycle
type PlantingCycleResponse struct {
	ID                  string
	CycleName           string
	GrowingZoneID       string
	PlantVarietyID      string
	SeedDate            *time.Time
	TransplantDate      *time.Time
	ExpectedHarvestDate *time.Time
	ActualHarvestDate   *time.Time
	PlantQuantity       int
	SeedBatch           string
	Status              string
	Notes               string
	CreatedBy           string
	CreatedAt           time.Time
	UpdatedAt           *time.Time
}

// PlantingCycleWithDetails represents planting cycle with related data
type PlantingCycleWithDetails struct {
	PlantingCycle
	PlantVariety *PlantVariety `json:"plant_variety,omitempty"`
}

// PlantingCycleWithDetailsResponse represents response for planting cycle with details
type PlantingCycleWithDetailsResponse struct {
	PlantingCycleResponse
	PlantVariety *PlantVarietyResponse
}

// ToResponse converts PlantingCycle to PlantingCycleResponse
func (pc *PlantingCycle) ToResponse() *PlantingCycleResponse {
	return &PlantingCycleResponse{
		ID:                  pc.ID,
		CycleName:           pc.CycleName,
		GrowingZoneID:       pc.GrowingZoneID,
		PlantVarietyID:      pc.PlantVarietyID,
		SeedDate:            pc.SeedDate,
		TransplantDate:      pc.TransplantDate,
		ExpectedHarvestDate: pc.ExpectedHarvestDate,
		ActualHarvestDate:   pc.ActualHarvestDate,
		PlantQuantity:       pc.PlantQuantity,
		SeedBatch:           pc.SeedBatch,
		Status:              pc.Status,
		Notes:               pc.Notes,
		CreatedBy:           pc.CreatedBy,
		CreatedAt:           pc.CreatedAt,
		UpdatedAt:           pc.UpdatedAt,
	}
}

// ToResponseWithDetails converts PlantingCycleWithDetails to PlantingCycleWithDetailsResponse
func (pc *PlantingCycleWithDetails) ToResponseWithDetails() *PlantingCycleWithDetailsResponse {
	response := &PlantingCycleWithDetailsResponse{
		PlantingCycleResponse: *pc.PlantingCycle.ToResponse(),
	}

	if pc.PlantVariety != nil {
		response.PlantVariety = pc.PlantVariety.ToResponse()
	}

	return response
}

// GetStatusDisplayName returns display name for status
func (pc *PlantingCycle) GetStatusDisplayName() string {
	statusMap := map[string]string{
		"planning":      "Lập kế hoạch",
		"seeding":       "Gieo hạt",
		"transplanting": "Cấy ghép",
		"growing":       "Phát triển",
		"flowering":     "Ra hoa",
		"harvesting":    "Thu hoạch",
		"completed":     "Hoàn thành",
		"failed":        "Thất bại",
	}

	if displayName, exists := statusMap[pc.Status]; exists {
		return displayName
	}
	return pc.Status
}

// IsActive returns true if the planting cycle is still active
func (pc *PlantingCycle) IsActive() bool {
	activeStatuses := []string{"planning", "seeding", "transplanting", "growing", "flowering", "harvesting"}
	for _, status := range activeStatuses {
		if pc.Status == status {
			return true
		}
	}
	return false
}

// GetDurationDays calculates the duration in days from seed date to current date or harvest date
func (pc *PlantingCycle) GetDurationDays() *int {
	if pc.SeedDate == nil {
		return nil
	}

	var endDate time.Time
	if pc.ActualHarvestDate != nil {
		endDate = *pc.ActualHarvestDate
	} else {
		endDate = time.Now()
	}

	duration := int(endDate.Sub(*pc.SeedDate).Hours() / 24)
	return &duration
}
