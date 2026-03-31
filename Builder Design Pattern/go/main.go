package main

import (
	"fmt"
	"strings"
)

// ============================================================
// PRODUCT - Complex Object
// ============================================================

// Pizza represents a pizza product with multiple configurable fields
type Pizza struct {
	Crust      string
	Size       string
	Sauce      string
	Toppings   []string
	Cheese     bool
	ExtraCheese bool
	IsSpicy    bool
	Price      float64
}

// GetDescription returns a string description of the pizza
func (p *Pizza) GetDescription() string {
	toppingCount := len(p.Toppings)
	return fmt.Sprintf("%s %s Pizza with %d toppings", p.Size, p.Crust, toppingCount)
}

// Display prints the pizza details in a formatted way
func (p *Pizza) Display() {
	fmt.Println()
	fmt.Println("🍕 PIZZA DETAILS:")
	fmt.Printf("   Crust: %s\n", p.Crust)
	fmt.Printf("   Size: %s\n", p.Size)
	fmt.Printf("   Sauce: %s\n", p.Sauce)

	if len(p.Toppings) > 0 {
		fmt.Printf("   Toppings: %s\n", strings.Join(p.Toppings, ", "))
	} else {
		fmt.Println("   Toppings: None")
	}

	cheeseStr := "No"
	if p.Cheese {
		cheeseStr = "Yes"
		if p.ExtraCheese {
			cheeseStr += " (Extra)"
		}
	}
	fmt.Printf("   Cheese: %s\n", cheeseStr)

	spicyStr := "No"
	if p.IsSpicy {
		spicyStr = "Yes 🌶️"
	}
	fmt.Printf("   Spicy: %s\n", spicyStr)
	fmt.Printf("   Price: $%.2f\n", p.Price)
	fmt.Println()
}

// ============================================================
// BUILDER - Step-by-step Construction with Fluent Interface
// ============================================================

// PizzaBuilder builds Pizza objects step-by-step using fluent interface
// The receiver methods return *PizzaBuilder to enable method chaining
type PizzaBuilder struct {
	pizza *Pizza
}

// NewPizzaBuilder creates and returns a new PizzaBuilder
func NewPizzaBuilder() *PizzaBuilder {
	return &PizzaBuilder{
		pizza: &Pizza{
			Crust:   "Regular",
			Size:    "Medium",
			Sauce:   "Tomato",
			Toppings: []string{},
			Cheese:  true,
			Price:   0,
		},
	}
}

// SetCrust sets the crust type and returns the builder for chaining
func (pb *PizzaBuilder) SetCrust(crust string) *PizzaBuilder {
	validCrusts := map[string]bool{
		"Regular": true,
		"Thin": true,
		"Thick": true,
		"Stuffed Crust": true,
		"Cauliflower": true,
	}

	if !validCrusts[crust] {
		panic(fmt.Sprintf("Invalid crust type: %s", crust))
	}

	pb.pizza.Crust = crust
	return pb
}

// SetSize sets the size and returns the builder for chaining
// Also auto-calculates base price based on size
func (pb *PizzaBuilder) SetSize(size string) *PizzaBuilder {
	priceMap := map[string]float64{
		"Small":     8.0,
		"Medium":    10.0,
		"Large":     12.0,
		"X-Large":   14.0,
		"2X-Large":  16.0,
	}

	price, ok := priceMap[size]
	if !ok {
		panic(fmt.Sprintf("Invalid size: %s", size))
	}

	pb.pizza.Size = size
	pb.pizza.Price = price
	return pb
}

// SetSauce sets the sauce type and returns the builder for chaining
func (pb *PizzaBuilder) SetSauce(sauce string) *PizzaBuilder {
	validSauces := map[string]bool{
		"Tomato":  true,
		"White":   true,
		"BBQ":     true,
		"Pesto":   true,
		"Buffalo": true,
	}

	if !validSauces[sauce] {
		panic(fmt.Sprintf("Invalid sauce: %s", sauce))
	}

	pb.pizza.Sauce = sauce
	return pb
}

// AddTopping adds a single topping and returns the builder for chaining
func (pb *PizzaBuilder) AddTopping(topping string) *PizzaBuilder {
	validToppings := map[string]bool{
		"Pepperoni":    true,
		"Mushrooms":    true,
		"Onions":       true,
		"Sausage":      true,
		"Bacon":        true,
		"Chicken":      true,
		"Olives":       true,
		"Bell Peppers": true,
		"Spinach":      true,
		"Feta":         true,
		"Jalapeños":    true,
	}

	if !validToppings[topping] {
		panic(fmt.Sprintf("Invalid topping: %s", topping))
	}

	// Avoid duplicate toppings
	for _, t := range pb.pizza.Toppings {
		if t == topping {
			return pb
		}
	}

	pb.pizza.Toppings = append(pb.pizza.Toppings, topping)
	pb.pizza.Price += 1.5 // $1.50 per topping

	return pb
}

// AddToppings adds multiple toppings and returns the builder for chaining
func (pb *PizzaBuilder) AddToppings(toppings []string) *PizzaBuilder {
	for _, topping := range toppings {
		pb.AddTopping(topping)
	}
	return pb
}

// SetCheese sets whether to include cheese and returns the builder for chaining
func (pb *PizzaBuilder) SetCheese(hasCheese bool) *PizzaBuilder {
	pb.pizza.Cheese = hasCheese
	return pb
}

// AddExtraCheese adds extra cheese and returns the builder for chaining
func (pb *PizzaBuilder) AddExtraCheese() *PizzaBuilder {
	pb.pizza.ExtraCheese = true
	pb.pizza.Price += 2.0 // $2.00 for extra cheese
	return pb
}

// MakeSpicy makes the pizza spicy and returns the builder for chaining
func (pb *PizzaBuilder) MakeSpicy() *PizzaBuilder {
	pb.pizza.IsSpicy = true
	pb.AddTopping("Jalapeños")
	return pb
}

// ApplyPreset applies a preset configuration and returns the builder for chaining
func (pb *PizzaBuilder) ApplyPreset(preset string) *PizzaBuilder {
	switch strings.ToLower(preset) {
	case "margherita":
		pb.SetCrust("Regular").
			SetSize("Medium").
			SetSauce("Tomato").
			SetCheese(true)

	case "pepperoni":
		pb.SetCrust("Regular").
			SetSize("Medium").
			SetSauce("Tomato").
			AddTopping("Pepperoni").
			SetCheese(true)

	case "vegetarian":
		pb.SetCrust("Regular").
			SetSize("Medium").
			SetSauce("Tomato").
			AddToppings([]string{"Mushrooms", "Onions", "Bell Peppers", "Olives"}).
			SetCheese(true)

	case "spicy":
		pb.SetCrust("Thick").
			SetSize("Large").
			SetSauce("Buffalo").
			AddToppings([]string{"Pepperoni", "Jalapeños", "Sausage"}).
			AddExtraCheese().
			MakeSpicy()

	case "deluxe":
		pb.SetCrust("Stuffed Crust").
			SetSize("Large").
			SetSauce("Tomato").
			AddToppings([]string{"Pepperoni", "Sausage", "Mushrooms", "Onions"}).
			AddExtraCheese()

	default:
		panic(fmt.Sprintf("Unknown preset: %s", preset))
	}

	return pb
}

// validate checks if the pizza configuration is valid
func (pb *PizzaBuilder) validate() error {
	if pb.pizza.Crust == "" {
		return fmt.Errorf("crust must be set")
	}
	if pb.pizza.Size == "" {
		return fmt.Errorf("size must be set")
	}
	if pb.pizza.Sauce == "" {
		return fmt.Errorf("sauce must be set")
	}
	if len(pb.pizza.Toppings) == 0 && !pb.pizza.Cheese {
		return fmt.Errorf("pizza must have either toppings or cheese")
	}
	return nil
}

// Build validates and returns the final Pizza object
func (pb *PizzaBuilder) Build() *Pizza {
	if err := pb.validate(); err != nil {
		panic(err)
	}
	return pb.pizza
}

// Reset resets the builder to create a new pizza
func (pb *PizzaBuilder) Reset() *PizzaBuilder {
	pb.pizza = &Pizza{
		Crust:   "Regular",
		Size:    "Medium",
		Sauce:   "Tomato",
		Toppings: []string{},
		Cheese:  true,
		Price:   0,
	}
	return pb
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

func main() {
	printHeader("BUILDER DESIGN PATTERN - Pizza Ordering System")

	// Example 1: Basic Pizza with Method Chaining
	exampleBasicPizza()

	// Example 2: Complex Pizza with Multiple Toppings
	exampleComplexPizza()

	// Example 3: Vegetarian Pizza
	exampleVegetarianPizza()

	// Example 4: Spicy Pizza
	exampleSpicyPizza()

	// Example 5: Using Presets
	examplePresets()

	// Example 6: Custom Pizza Combination
	exampleCustomPizza()

	// Advanced Example: Restaurant Order Management
	advancedRestaurantExample()
}

// Example 1: Basic Pizza with Method Chaining
func exampleBasicPizza() {
	fmt.Println("\n📋 EXAMPLE 1: Build a Basic Pepperoni Pizza")
	fmt.Println(strings.Repeat("-", 70))

	pizza := NewPizzaBuilder().
		SetCrust("Regular").
		SetSize("Large").
		SetSauce("Tomato").
		AddTopping("Pepperoni").
		SetCheese(true).
		Build()

	pizza.Display()
}

// Example 2: Complex Pizza with Multiple Toppings
func exampleComplexPizza() {
	fmt.Println("📋 EXAMPLE 2: Build a Complex Pizza with Multiple Toppings")
	fmt.Println(strings.Repeat("-", 70))

	pizza := NewPizzaBuilder().
		SetCrust("Stuffed Crust").
		SetSize("X-Large").
		SetSauce("Tomato").
		AddTopping("Pepperoni").
		AddTopping("Sausage").
		AddTopping("Mushrooms").
		AddTopping("Onions").
		AddExtraCheese().
		Build()

	pizza.Display()
}

// Example 3: Vegetarian Pizza
func exampleVegetarianPizza() {
	fmt.Println("📋 EXAMPLE 3: Build a Vegetarian Pizza")
	fmt.Println(strings.Repeat("-", 70))

	pizza := NewPizzaBuilder().
		SetCrust("Thin").
		SetSize("Medium").
		SetSauce("Pesto").
		AddToppings([]string{"Mushrooms", "Onions", "Bell Peppers", "Spinach"}).
		SetCheese(true).
		Build()

	pizza.Display()
}

// Example 4: Spicy Pizza
func exampleSpicyPizza() {
	fmt.Println("📋 EXAMPLE 4: Build a Spicy Pizza")
	fmt.Println(strings.Repeat("-", 70))

	pizza := NewPizzaBuilder().
		SetCrust("Thick").
		SetSize("Large").
		SetSauce("Buffalo").
		AddToppings([]string{"Pepperoni", "Jalapeños", "Sausage"}).
		MakeSpicy().
		AddExtraCheese().
		Build()

	pizza.Display()
}

// Example 5: Using Presets
func examplePresets() {
	fmt.Println("📋 EXAMPLE 5: Build Pizza Using Presets")
	fmt.Println(strings.Repeat("-", 70))

	presets := []string{"Margherita", "Pepperoni", "Vegetarian", "Spicy", "Deluxe"}

	fmt.Println("Available Presets:")
	builder := NewPizzaBuilder()

	for _, preset := range presets {
		pizza := builder.Reset().ApplyPreset(preset).Build()
		fmt.Printf("✓ %s Crust %s: $%.2f\n", pizza.Crust, preset, pizza.Price)
	}
	fmt.Println()
}

// Example 6: Custom Pizza with White Sauce
func exampleCustomPizza() {
	fmt.Println("📋 EXAMPLE 6: Custom Pizza with White Sauce")
	fmt.Println(strings.Repeat("-", 70))

	pizza := NewPizzaBuilder().
		SetCrust("Cauliflower").
		SetSize("Medium").
		SetSauce("White").
		AddToppings([]string{"Chicken", "Spinach", "Feta"}).
		AddExtraCheese().
		Build()

	pizza.Display()
}

// ============================================================
// ADVANCED EXAMPLE: Restaurant Order Management
// ============================================================

// Order represents a customer order with multiple pizzas
type Order struct {
	OrderID     string
	CustomerName string
	Pizzas      []*Pizza
	OrderTime   string
}

// NewOrder creates a new order
func NewOrder(orderID, customerName string) *Order {
	return &Order{
		OrderID:      orderID,
		CustomerName: customerName,
		Pizzas:       make([]*Pizza, 0),
		OrderTime:    getTimeString(),
	}
}

// AddPizza adds a pizza to the order
func (o *Order) AddPizza(pizza *Pizza) *Order {
	o.Pizzas = append(o.Pizzas, pizza)
	return o
}

// GetTotalPrice calculates the total order price
func (o *Order) GetTotalPrice() float64 {
	total := 0.0
	for _, pizza := range o.Pizzas {
		total += pizza.Price
	}
	return total
}

// Display prints the order details
func (o *Order) Display() {
	fmt.Println()
	fmt.Printf("📦 ORDER #%s\n", o.OrderID)
	fmt.Printf("   Customer: %s\n", o.CustomerName)
	fmt.Printf("   Time: %s\n", o.OrderTime)
	fmt.Printf("   Items: %d pizza(s)\n", len(o.Pizzas))
	fmt.Println("   " + strings.Repeat("─", 36))

	for i, pizza := range o.Pizzas {
		fmt.Printf("   %d. %s - $%.2f\n", i+1, pizza.GetDescription(), pizza.Price)
	}

	fmt.Println("   " + strings.Repeat("─", 36))
	fmt.Printf("   Total: $%.2f\n", o.GetTotalPrice())
	fmt.Println()
}

// advancedRestaurantExample demonstrates restaurant order management
func advancedRestaurantExample() {
	fmt.Println()
	printHeader("ADVANCED EXAMPLE: Restaurant Order Management System")

	fmt.Println("Processing Customer Orders:")
	fmt.Println(strings.Repeat("-", 70))

	// Order 1
	order1 := NewOrder("ORD001", "John Doe").
		AddPizza(NewPizzaBuilder().ApplyPreset("Margherita").Build()).
		AddPizza(NewPizzaBuilder().ApplyPreset("Pepperoni").Build())
	order1.Display()

	// Order 2
	order2 := NewOrder("ORD002", "Jane Smith").
		AddPizza(NewPizzaBuilder().ApplyPreset("Vegetarian").Build()).
		AddPizza(NewPizzaBuilder().ApplyPreset("Deluxe").Build()).
		AddPizza(NewPizzaBuilder().ApplyPreset("Spicy").Build())
	order2.Display()

	// Order 3
	order3 := NewOrder("ORD003", "Bob Johnson").
		AddPizza(NewPizzaBuilder().
			SetCrust("Thick").
			SetSize("X-Large").
			SetSauce("BBQ").
			AddToppings([]string{"Chicken", "Bacon", "Onions"}).
			AddExtraCheese().
			Build())
	order3.Display()

	// Summary Statistics
	fmt.Println("📊 Restaurant Daily Summary:")
	fmt.Println(strings.Repeat("-", 70))

	orders := []*Order{order1, order2, order3}
	totalOrders := len(orders)
	totalPizzas := 0
	totalRevenue := 0.0

	for _, order := range orders {
		totalPizzas += len(order.Pizzas)
		totalRevenue += order.GetTotalPrice()
	}

	averageOrderValue := totalRevenue / float64(totalOrders)
	avgPizzasPerOrder := float64(totalPizzas) / float64(totalOrders)

	fmt.Printf("Total Orders: %d\n", totalOrders)
	fmt.Printf("Total Pizzas Sold: %d\n", totalPizzas)
	fmt.Printf("Total Revenue: $%.2f\n", totalRevenue)
	fmt.Printf("Average Order Value: $%.2f\n", averageOrderValue)
	fmt.Printf("Average Pizzas per Order: %.1f\n", avgPizzasPerOrder)
}

// ============================================================
// UTILITY FUNCTIONS
// ============================================================

// printHeader prints a formatted header
func printHeader(title string) {
	fmt.Println()
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println(title)
	fmt.Println(strings.Repeat("=", 70))
	fmt.Println()
}

// getTimeString returns the current time as a formatted string
func getTimeString() string {
	return "2024-01-15 14:30:00"
}