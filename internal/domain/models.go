package domain

import "time"

// Order represents a shipment request.
type Order struct {
	ID           string    `json:"id"`
	PayoutCents  int64     `json:"payout_cents"`
	WeightLbs    int       `json:"weight_lbs"`
	VolumeCuFt   int       `json:"volume_cuft"`
	Origin       string    `json:"origin"`
	Destination  string    `json:"destination"`
	PickupDate   string    `json:"pickup_date"`   // YYYY-MM-DD
	DeliveryDate string    `json:"delivery_date"` // YYYY-MM-DD
	IsHazmat     bool      `json:"is_hazmat"`
	ParsedPickup time.Time `json:"-"`
	ParsedDelivery time.Time `json:"-"`
}

// Truck represents the carrier's vehicle constraints.
type Truck struct {
	ID            string `json:"id"`
	MaxWeightLbs  int    `json:"max_weight_lbs"`
	MaxVolumeCuFt int    `json:"max_volume_cuft"`
}

// OptimizationRequest is the incoming payload for the optimization endpoint.
type OptimizationRequest struct {
	Truck  Truck   `json:"truck"`
	Orders []Order `json:"orders"`
}

// OptimizationResponse is the result of the optimization process.
type OptimizationResponse struct {
	TruckID                   string   `json:"truck_id"`
	SelectedOrderIDs          []string `json:"selected_order_ids"`
	TotalPayoutCents          int64    `json:"total_payout_cents"`
	TotalWeightLbs            int      `json:"total_weight_lbs"`
	TotalVolumeCuFt           int      `json:"total_volume_cuft"`
	UtilizationWeightPercent  float64  `json:"utilization_weight_percent"`
	UtilizationVolumePercent  float64  `json:"utilization_volume_percent"`
}
