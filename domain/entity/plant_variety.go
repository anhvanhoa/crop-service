package entity

import (
	"time"
)

// PlantVariety represents a plant variety entity
type PlantVariety struct {
	ID                 string
	Name               string
	ScientificName     string
	Category           string
	GrowingSeason      string
	GrowthDurationDays int
	OptimalTempMin     float64
	OptimalTempMax     float64
	OptimalHumidityMin float64
	OptimalHumidityMax float64
	PHMin              float64
	PHMax              float64
	WaterRequirement   string
	LightRequirement   string
	Description        string
	MediaID            string
	Status             string
	CreatedBy          string
	CreatedAt          time.Time
	UpdatedAt          *time.Time
}

// CreatePlantVarietyRequest represents request to create a new plant variety
type CreatePlantVarietyRequest struct {
	Name               string
	ScientificName     string
	Category           string
	GrowingSeason      string
	GrowthDurationDays int
	OptimalTempMin     float64
	OptimalTempMax     float64
	OptimalHumidityMin float64
	OptimalHumidityMax float64
	PHMin              float64
	PHMax              float64
	WaterRequirement   string
	LightRequirement   string
	Description        string
	MediaID            string
	CreatedBy          string
}

// UpdatePlantVarietyRequest represents request to update a plant variety
type UpdatePlantVarietyRequest struct {
	ID                 string
	Name               string
	ScientificName     string
	Category           string
	GrowingSeason      string
	GrowthDurationDays int
	OptimalTempMin     float64
	OptimalTempMax     float64
	OptimalHumidityMin float64
	OptimalHumidityMax float64
	PHMin              float64
	PHMax              float64
	WaterRequirement   string
	LightRequirement   string
	Description        string
	MediaID            string
	Status             string
}

// PlantVarietyFilter represent
type PlantVarietyFilter struct {
	ID            string
	Name          string
	Category      string
	GrowingSeason string
	Status        string
	CreatedBy     string
	Limit         int
	Offset        int
	SortBy        string
	SortDirection string
}

// PlantVarietyResponse represents response for plant variety
type PlantVarietyResponse struct {
	ID                 string
	Name               string
	ScientificName     string
	Category           string
	GrowingSeason      string
	GrowthDurationDays int
	OptimalTempMin     float64
	OptimalTempMax     float64
	OptimalHumidityMin float64
	OptimalHumidityMax float64
	PHMin              float64
	PHMax              float64
	WaterRequirement   string
	LightRequirement   string
	Description        string
	MediaID            string
	Status             string
	CreatedBy          string
	CreatedAt          time.Time
	UpdatedAt          *time.Time
}

// ToResponse converts PlantVariety to PlantVarietyResponse
func (pv *PlantVariety) ToResponse() *PlantVarietyResponse {
	return &PlantVarietyResponse{
		ID:                 pv.ID,
		Name:               pv.Name,
		ScientificName:     pv.ScientificName,
		Category:           pv.Category,
		GrowingSeason:      pv.GrowingSeason,
		GrowthDurationDays: pv.GrowthDurationDays,
		OptimalTempMin:     pv.OptimalTempMin,
		OptimalTempMax:     pv.OptimalTempMax,
		OptimalHumidityMin: pv.OptimalHumidityMin,
		OptimalHumidityMax: pv.OptimalHumidityMax,
		PHMin:              pv.PHMin,
		PHMax:              pv.PHMax,
		WaterRequirement:   pv.WaterRequirement,
		LightRequirement:   pv.LightRequirement,
		Description:        pv.Description,
		MediaID:            pv.MediaID,
		Status:             pv.Status,
		CreatedBy:          pv.CreatedBy,
		CreatedAt:          pv.CreatedAt,
		UpdatedAt:          pv.UpdatedAt,
	}
}
