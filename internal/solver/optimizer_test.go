package solver

import (
	"reflect"
	"sort"
	"testing"

	"github.com/poojanmishra/SmartLoad/internal/domain"
)

func TestOptimizer_Optimize_Basic(t *testing.T) {
	truck := domain.Truck{ID: "t1", MaxWeightLbs: 100, MaxVolumeCuFt: 100}
	orders := []domain.Order{
		{ID: "o1", PayoutCents: 10, WeightLbs: 50, VolumeCuFt: 50, Origin: "A", Destination: "B", PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"},
		{ID: "o2", PayoutCents: 20, WeightLbs: 40, VolumeCuFt: 40, Origin: "A", Destination: "B", PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"},
		{ID: "o3", PayoutCents: 100, WeightLbs: 200, VolumeCuFt: 200, Origin: "A", Destination: "B", PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"}, // Too big
	}

	opt := NewOptimizer()
	resp := opt.Optimize(domain.OptimizationRequest{Truck: truck, Orders: orders})

	expectedIDs := []string{"o1", "o2"}
	sort.Strings(resp.SelectedOrderIDs)
	sort.Strings(expectedIDs)

	if !reflect.DeepEqual(resp.SelectedOrderIDs, expectedIDs) {
		t.Errorf("Expected %v, got %v", expectedIDs, resp.SelectedOrderIDs)
	}
	if resp.TotalPayoutCents != 30 {
		t.Errorf("Expected payout 30, got %d", resp.TotalPayoutCents)
	}
}

func TestOptimizer_HazmatIsolation(t *testing.T) {
	truck := domain.Truck{ID: "t1", MaxWeightLbs: 100, MaxVolumeCuFt: 100}
	orders := []domain.Order{
		{ID: "h1", PayoutCents: 50, WeightLbs: 10, VolumeCuFt: 10, Origin: "A", Destination: "B", IsHazmat: true, PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"},
		{ID: "n1", PayoutCents: 40, WeightLbs: 10, VolumeCuFt: 10, Origin: "A", Destination: "B", IsHazmat: false, PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"},
		{ID: "h2", PayoutCents: 60, WeightLbs: 10, VolumeCuFt: 10, Origin: "A", Destination: "B", IsHazmat: true, PickupDate: "2023-01-01", DeliveryDate: "2023-01-02"},
	}

	opt := NewOptimizer()
	resp := opt.Optimize(domain.OptimizationRequest{Truck: truck, Orders: orders})

	// Should prioritize Hazmat group (50+60=110) over Non-Hazmat (40) if both fit,
	// OR pick Non-Hazmat if it was worth more. Here Hazmat wins.
	// But it CANNOT mix h1 + n1 (90) or h2 + n1 (100).
	// Max possible: h1+h2 = 110.

	expectedIDs := []string{"h1", "h2"}
	sort.Strings(resp.SelectedOrderIDs)
	sort.Strings(expectedIDs)

	if !reflect.DeepEqual(resp.SelectedOrderIDs, expectedIDs) {
		t.Errorf("Expected %v, got %v", expectedIDs, resp.SelectedOrderIDs)
	}
}

func TestOptimizer_RouteGrouping(t *testing.T) {
	truck := domain.Truck{ID: "t1", MaxWeightLbs: 100, MaxVolumeCuFt: 100}
	orders := []domain.Order{
		{ID: "r1", PayoutCents: 50, Origin: "A", Destination: "B", PickupDate: "2023-01-01", DeliveryDate: "2023-01-02", WeightLbs: 10, VolumeCuFt: 10},
		{ID: "r2", PayoutCents: 60, Origin: "A", Destination: "C", PickupDate: "2023-01-01", DeliveryDate: "2023-01-02", WeightLbs: 10, VolumeCuFt: 10},
	}

	opt := NewOptimizer()
	resp := opt.Optimize(domain.OptimizationRequest{Truck: truck, Orders: orders})

	// Different routes, cannot combine. Should pick the best single route.
	// r2 (60) > r1 (50)

	if len(resp.SelectedOrderIDs) != 1 || resp.SelectedOrderIDs[0] != "r2" {
		t.Errorf("Expected r2, got %v", resp.SelectedOrderIDs)
	}
}
