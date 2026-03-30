package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// ============================================================
// PRODUCT INTERFACE
// ============================================================

// ShippingCarrier interface defines the contract for all shipping carriers
type ShippingCarrier interface {
	// Ship sends a package and returns shipping result
	Ship(pkg Package) ShippingResult

	// CalculateCost calculates the shipping cost based on weight
	CalculateCost(weight float64) float64

	// GetCarrierName returns the name of the carrier
	GetCarrierName() string

	// GetEstimatedDelivery returns estimated delivery time
	GetEstimatedDelivery(destination string) string
}

// Package represents a package to be shipped
type Package struct {
	ID          string
	Destination string
	Weight      float64
}

// ShippingResult holds the result of a shipping operation
type ShippingResult struct {
	Carrier   string
	Tracking  string
	Status    string
	Timestamp string
}

// ============================================================
// CONCRETE PRODUCTS - Shipping Carriers
// ============================================================

// FedExCarrier implements ShippingCarrier
type FedExCarrier struct{}

func (f *FedExCarrier) Ship(pkg Package) ShippingResult {
	tracking := fmt.Sprintf("FX%d", rand.Intn(1000000000))
	fmt.Printf("📦 FedEx: Shipping package to %s\n", pkg.Destination)
	fmt.Printf("   Tracking: %s\n", tracking)
	fmt.Printf("   Status: ✅ Package picked up\n\n")
	return ShippingResult{
		Carrier:   f.GetCarrierName(),
		Tracking:  tracking,
		Status:    "Picked up",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (f *FedExCarrier) CalculateCost(weight float64) float64 {
	// FedEx: $1.50 per kg + $5 base fee
	baseFee := 5.0
	perKgCost := 1.5
	return baseFee + (weight * perKgCost)
}

func (f *FedExCarrier) GetCarrierName() string {
	return "FedEx"
}

func (f *FedExCarrier) GetEstimatedDelivery(destination string) string {
	return "2-3 business days"
}

// UPSCarrier implements ShippingCarrier
type UPSCarrier struct{}

func (u *UPSCarrier) Ship(pkg Package) ShippingResult {
	tracking := fmt.Sprintf("1Z%d", rand.Intn(1000000000))
	fmt.Printf("📦 UPS: Shipping package to %s\n", pkg.Destination)
	fmt.Printf("   Tracking: %s\n", tracking)
	fmt.Printf("   Status: ✅ In transit\n\n")
	return ShippingResult{
		Carrier:   u.GetCarrierName(),
		Tracking:  tracking,
		Status:    "In transit",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (u *UPSCarrier) CalculateCost(weight float64) float64 {
	// UPS: $1.25 per kg + $4 base fee
	baseFee := 4.0
	perKgCost := 1.25
	return baseFee + (weight * perKgCost)
}

func (u *UPSCarrier) GetCarrierName() string {
	return "UPS"
}

func (u *UPSCarrier) GetEstimatedDelivery(destination string) string {
	return "1-2 business days"
}

// DHLCarrier implements ShippingCarrier
type DHLCarrier struct{}

func (d *DHLCarrier) Ship(pkg Package) ShippingResult {
	tracking := fmt.Sprintf("DHL%d", rand.Intn(1000000000))
	fmt.Printf("📦 DHL: Shipping package to %s\n", pkg.Destination)
	fmt.Printf("   Tracking: %s\n", tracking)
	fmt.Printf("   Status: ✅ Label created\n\n")
	return ShippingResult{
		Carrier:   d.GetCarrierName(),
		Tracking:  tracking,
		Status:    "Label created",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (d *DHLCarrier) CalculateCost(weight float64) float64 {
	// DHL: $1.75 per kg + $6 base fee
	baseFee := 6.0
	perKgCost := 1.75
	return baseFee + (weight * perKgCost)
}

func (d *DHLCarrier) GetCarrierName() string {
	return "DHL"
}

func (d *DHLCarrier) GetEstimatedDelivery(destination string) string {
	return "3-5 business days"
}

// ExpressOvernight implements ShippingCarrier
type ExpressOvernight struct{}

func (e *ExpressOvernight) Ship(pkg Package) ShippingResult {
	tracking := fmt.Sprintf("EXP%d", rand.Intn(1000000000))
	fmt.Printf("🚀 Express Overnight: Shipping package to %s\n", pkg.Destination)
	fmt.Printf("   Tracking: %s\n", tracking)
	fmt.Printf("   Status: ✅ Priority handling\n\n")
	return ShippingResult{
		Carrier:   e.GetCarrierName(),
		Tracking:  tracking,
		Status:    "Priority handling",
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (e *ExpressOvernight) CalculateCost(weight float64) float64 {
	// Express: $5 per kg + $15 base fee + 50% surcharge
	baseFee := 15.0
	perKgCost := 5.0
	totalCost := baseFee + (weight * perKgCost)
	return totalCost * 1.5 // 50% premium for overnight
}

func (e *ExpressOvernight) GetCarrierName() string {
	return "Express Overnight"
}

func (e *ExpressOvernight) GetEstimatedDelivery(destination string) string {
	return "Next business day (guaranteed)"
}

// ============================================================
// FACTORY - Object Creator
// ============================================================

// ShippingCarrierFactory creates shipping carriers
type ShippingCarrierFactory struct{}

// CreateCarrier creates a shipping carrier based on type
func (f *ShippingCarrierFactory) CreateCarrier(carrierType string) (ShippingCarrier, error) {
	carrierType = strings.ToLower(strings.TrimSpace(carrierType))

	switch carrierType {
	case "fedex", "fedx":
		return &FedExCarrier{}, nil

	case "ups":
		return &UPSCarrier{}, nil

	case "dhl":
		return &DHLCarrier{}, nil

	case "express", "overnight", "express-overnight":
		return &ExpressOvernight{}, nil

	default:
		return nil, fmt.Errorf("unknown carrier: %s. Available: fedex, ups, dhl, express", carrierType)
	}
}

// GetAvailableCarriers returns list of available carriers
func (f *ShippingCarrierFactory) GetAvailableCarriers() []string {
	return []string{"fedex", "ups", "dhl", "express"}
}

// FindCheapestCarrier finds the cheapest carrier for given weight
func (f *ShippingCarrierFactory) FindCheapestCarrier(weight float64) (ShippingCarrier, float64, error) {
	carriers := f.GetAvailableCarriers()
	var cheapest ShippingCarrier
	minCost := float64(999999)

	for _, carrierType := range carriers {
		carrier, err := f.CreateCarrier(carrierType)
		if err != nil {
			continue
		}

		cost := carrier.CalculateCost(weight)
		if cost < minCost {
			minCost = cost
			cheapest = carrier
		}
	}

	if cheapest == nil {
		return nil, 0, fmt.Errorf("could not find any carrier")
	}

	return cheapest, minCost, nil
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

func main() {
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("FACTORY DESIGN PATTERN - Go Implementation")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	factory := &ShippingCarrierFactory{}

	// Example 1: Create Different Carriers Using Factory
	fmt.Println("📋 EXAMPLE 1: Create Different Carriers Using Factory")
	fmt.Println(strings.Repeat("-", 70))

	pkg1 := Package{
		ID:          "PKG001",
		Destination: "New York, NY",
		Weight:      2.5,
	}

	// Create FedEx carrier using factory
	fedexCarrier, _ := factory.CreateCarrier("fedex")
	fmt.Printf("Carrier: %s\n", fedexCarrier.GetCarrierName())
	fmt.Printf("Cost: $%.2f\n", fedexCarrier.CalculateCost(pkg1.Weight))
	fmt.Printf("Delivery: %s\n", fedexCarrier.GetEstimatedDelivery(pkg1.Destination))
	fedexCarrier.Ship(pkg1)

	// Example 2: Switch to Different Carrier
	fmt.Println("📋 EXAMPLE 2: Switch to Different Carrier")
	fmt.Println(strings.Repeat("-", 70))

	upsCarrier, _ := factory.CreateCarrier("ups")
	fmt.Printf("Carrier: %s\n", upsCarrier.GetCarrierName())
	fmt.Printf("Cost: $%.2f\n", upsCarrier.CalculateCost(pkg1.Weight))
	fmt.Printf("Delivery: %s\n", upsCarrier.GetEstimatedDelivery(pkg1.Destination))
	upsCarrier.Ship(pkg1)

	// Example 3: Process Multiple Packages with Different Carriers
	fmt.Println("📋 EXAMPLE 3: Multiple Packages with Different Carriers")
	fmt.Println(strings.Repeat("-", 70))

	packages := []struct {
		pkg       Package
		carrierType string
	}{
		{Package{"PKG002", "Los Angeles, CA", 1.2}, "dhl"},
		{Package{"PKG003", "Chicago, IL", 5.0}, "fedex"},
		{Package{"PKG004", "Miami, FL", 0.8}, "ups"},
	}

	for _, p := range packages {
		carrier, _ := factory.CreateCarrier(p.carrierType)
		cost := carrier.CalculateCost(p.pkg.Weight)
		fmt.Printf("📦 %s: %s - $%.2f - %s\n", p.pkg.ID, carrier.GetCarrierName(), cost, p.pkg.Destination)
		carrier.Ship(p.pkg)
	}

	// Example 4: Available Carriers List
	fmt.Println("📋 EXAMPLE 4: Available Carriers")
	fmt.Println(strings.Repeat("-", 70))
	availableCarriers := factory.GetAvailableCarriers()
	fmt.Println("Available carriers:")
	for _, carrierType := range availableCarriers {
		carrier, _ := factory.CreateCarrier(carrierType)
		fmt.Printf("  ✓ %s: $%.2f/kg (base)\n", carrier.GetCarrierName(), carrier.CalculateCost(1))
	}
	fmt.Println()

	// Example 5: Price Comparison
	fmt.Println("📋 EXAMPLE 5: Price Comparison - Find Cheapest Carrier")
	fmt.Println(strings.Repeat("-", 70))

	weights := []float64{1, 5, 10, 25}
	for _, weight := range weights {
		cheapest, cost, _ := factory.FindCheapestCarrier(weight)
		fmt.Printf("For %.0fkg: %s is cheapest at $%.2f\n", weight, cheapest.GetCarrierName(), cost)
	}
	fmt.Println()

	// ============================================================
	// ADVANCED EXAMPLE: Order Shipment Manager
	// ============================================================

	fmt.Println()
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println("ADVANCED EXAMPLE: Order Shipment Manager")
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()

	manager := NewShipmentManager()

	orders := []Order{
		{"ORD001", "3 books", 2.5, "New York, NY"},
		{"ORD002", "1 laptop", 3.2, "San Francisco, CA"},
		{"ORD003", "5 shoes", 5.5, "Miami, FL"},
		{"ORD004", "2 watches", 0.5, "Boston, MA"},
	}

	fmt.Println("Processing Orders:")
	fmt.Println(strings.Repeat("-", 70))

	// Process some orders with specific carriers
	manager.CreateShipment(orders[0], "fedex", factory)
	manager.CreateShipment(orders[1], "ups", factory)
	manager.CreateShipmentCheapest(orders[2], factory)
	manager.CreateShipmentCheapest(orders[3], factory)

	// Show history and costs
	manager.GetShipmentHistory()
	fmt.Printf("\nTotal Shipping Cost: $%.2f\n", manager.GetTotalShippingCost())

// ============================================================
// ADVANCED EXAMPLE: Order and Shipment Manager
// ============================================================

// Order represents an order to be shipped
type Order struct {
	ID          string
	Items       string
	Weight      float64
	Destination string
}

// Shipment represents a shipped order
type Shipment struct {
	OrderID  string
	Carrier  string
	Tracking string
	Status   string
	Cost     float64
}

// ShipmentManager manages order shipments
type ShipmentManager struct {
	shipments []Shipment
}

// NewShipmentManager creates a new shipment manager
func NewShipmentManager() *ShipmentManager {
	return &ShipmentManager{
		shipments: []Shipment{},
	}
}

// CreateShipment creates a shipment with specified carrier
func (sm *ShipmentManager) CreateShipment(order Order, carrierType string, factory *ShippingCarrierFactory) error {
	fmt.Printf("\n📦 Processing Order: %s\n", order.ID)
	fmt.Printf("   Items: %s\n", order.Items)
	fmt.Printf("   Weight: %.1fkg\n", order.Weight)
	fmt.Printf("   Destination: %s\n", order.Destination)

	// Use factory to create carrier
	carrier, err := factory.CreateCarrier(carrierType)
	if err != nil {
		return err
	}

	fmt.Printf("   Carrier: %s\n", carrier.GetCarrierName())
	cost := carrier.CalculateCost(order.Weight)
	fmt.Printf("   Shipping Cost: $%.2f\n", cost)
	fmt.Printf("   Delivery: %s\n", carrier.GetEstimatedDelivery(order.Destination))

	// Ship the package
	result := carrier.Ship(Package{
		ID:          order.ID,
		Destination: order.Destination,
		Weight:      order.Weight,
	})

	// Record shipment
	sm.shipments = append(sm.shipments, Shipment{
		OrderID:  order.ID,
		Carrier:  result.Carrier,
		Tracking: result.Tracking,
		Status:   result.Status,
		Cost:     cost,
	})

	return nil
}

// CreateShipmentCheapest auto-selects the cheapest carrier
func (sm *ShipmentManager) CreateShipmentCheapest(order Order, factory *ShippingCarrierFactory) error {
	fmt.Printf("\n📦 Processing Order: %s (Auto-select cheapest)\n", order.ID)
	fmt.Printf("   Items: %s\n", order.Items)
	fmt.Printf("   Weight: %.1fkg\n", order.Weight)

	cheapest, _, _ := factory.FindCheapestCarrier(order.Weight)
	return sm.CreateShipment(order, strings.ToLower(cheapest.GetCarrierName()), factory)
}

// GetShipmentHistory displays the shipment history
func (sm *ShipmentManager) GetShipmentHistory() {
	fmt.Println("\n📊 Shipment History:")
	for i, shipment := range sm.shipments {
		fmt.Printf("   %d. Order %s - %s - Tracking: %s - Cost: $%.2f\n",
			i+1, shipment.OrderID, shipment.Carrier, shipment.Tracking, shipment.Cost)
	}
}

// GetTotalShippingCost calculates total shipping cost
func (sm *ShipmentManager) GetTotalShippingCost() float64 {
	total := 0.0
	for _, shipment := range sm.shipments {
		total += shipment.Cost
	}
	return total
}
