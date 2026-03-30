package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// ============================================================
// STRATEGY INTERFACE
// ============================================================

// PaymentStrategy interface defines the contract for all payment strategies
type PaymentStrategy interface {
	// Pay executes the payment strategy
	Pay(amount float64) PaymentResult

	// Validate checks if the payment credentials are valid
	Validate() bool

	// GetStrategyName returns the name of the payment strategy
	GetStrategyName() string
}

// PaymentResult holds the result of a payment operation
type PaymentResult struct {
	Success   bool
	Method    string
	Amount    float64
	BTCAmount string
	Timestamp string
}

// ============================================================
// CONCRETE STRATEGIES - Payment Methods
// ============================================================

// CreditCardPayment is a strategy for credit card payments
type CreditCardPayment struct {
	CardNumber string
	CardHolder string
	ExpiryDate string
	CVV        string
}

func (c *CreditCardPayment) Pay(amount float64) PaymentResult {
	fmt.Printf("🏦 Processing Credit Card Payment\n")
	fmt.Printf("   Card: %s\n", c.maskCardNumber(c.CardNumber))
	fmt.Printf("   Holder: %s\n", c.CardHolder)
	fmt.Printf("   Amount: $%.2f\n", amount)
	fmt.Printf("   Status: ✅ Payment successful!\n\n")

	return PaymentResult{
		Success:   true,
		Method:    c.GetStrategyName(),
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (c *CreditCardPayment) Validate() bool {
	// Check if card number is valid (basic check)
	return len(c.CardNumber) >= 13 && len(c.CVV) >= 3
}

func (c *CreditCardPayment) GetStrategyName() string {
	return "Credit Card"
}

func (c *CreditCardPayment) maskCardNumber(cardNumber string) string {
	if len(cardNumber) < 4 {
		return "****"
	}
	last4 := cardNumber[len(cardNumber)-4:]
	return "****-****-****-" + last4
}

// PayPalPayment is a strategy for PayPal payments
type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) PaymentResult {
	fmt.Printf("📱 Processing PayPal Payment\n")
	fmt.Printf("   Email: %s\n", p.Email)
	fmt.Printf("   Amount: $%.2f\n", amount)
	fmt.Printf("   Authenticating with PayPal...\n")
	fmt.Printf("   Status: ✅ Payment successful!\n\n")

	return PaymentResult{
		Success:   true,
		Method:    p.GetStrategyName(),
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (p *PayPalPayment) Validate() bool {
	// Check if email is valid using regex
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(p.Email)
}

func (p *PayPalPayment) GetStrategyName() string {
	return "PayPal"
}

// BitcoinPayment is a strategy for Bitcoin payments
type BitcoinPayment struct {
	WalletAddress string
}

func (b *BitcoinPayment) Pay(amount float64) PaymentResult {
	btcAmount := fmt.Sprintf("%.6f", amount*0.000025)

	fmt.Printf("🪙 Processing Bitcoin Payment\n")
	fmt.Printf("   Wallet: %s\n", b.maskWallet(b.WalletAddress))
	fmt.Printf("   Amount: %s BTC\n", btcAmount)
	fmt.Printf("   Confirming transaction on blockchain...\n")
	fmt.Printf("   Status: ✅ Payment successful!\n\n")

	return PaymentResult{
		Success:   true,
		Method:    b.GetStrategyName(),
		Amount:    amount,
		BTCAmount: btcAmount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (b *BitcoinPayment) Validate() bool {
	// Check if wallet address is valid (basic check)
	return len(b.WalletAddress) > 25
}

func (b *BitcoinPayment) GetStrategyName() string {
	return "Bitcoin"
}

func (b *BitcoinPayment) maskWallet(address string) string {
	if len(address) < 10 {
		return address
	}
	first6 := address[:6]
	last4 := address[len(address)-4:]
	return first6 + "..." + last4
}

// ApplePayPayment is a strategy for Apple Pay payments
type ApplePayPayment struct {
	AppleID string
}

func (a *ApplePayPayment) Pay(amount float64) PaymentResult {
	fmt.Printf("🍎 Processing Apple Pay Payment\n")
	fmt.Printf("   Apple ID: %s\n", a.AppleID)
	fmt.Printf("   Amount: $%.2f\n", amount)
	fmt.Printf("   Using biometric authentication...\n")
	fmt.Printf("   Status: ✅ Payment successful!\n\n")

	return PaymentResult{
		Success:   true,
		Method:    a.GetStrategyName(),
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (a *ApplePayPayment) Validate() bool {
	// Check if Apple ID is valid
	return strings.Contains(a.AppleID, "@")
}

func (a *ApplePayPayment) GetStrategyName() string {
	return "Apple Pay"
}

// GooglePayPayment is a strategy for Google Pay payments
type GooglePayPayment struct {
	Email string
}

func (g *GooglePayPayment) Pay(amount float64) PaymentResult {
	fmt.Printf("🔵 Processing Google Pay Payment\n")
	fmt.Printf("   Email: %s\n", g.Email)
	fmt.Printf("   Amount: $%.2f\n", amount)
	fmt.Printf("   Authenticating with Google Account...\n")
	fmt.Printf("   Status: ✅ Payment successful!\n\n")

	return PaymentResult{
		Success:   true,
		Method:    g.GetStrategyName(),
		Amount:    amount,
		Timestamp: time.Now().Format(time.RFC3339),
	}
}

func (g *GooglePayPayment) Validate() bool {
	// Check if email is valid
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	return emailRegex.MatchString(g.Email)
}

func (g *GooglePayPayment) GetStrategyName() string {
	return "Google Pay"
}

// ============================================================
// CONTEXT - Payment Processor
// ============================================================

// PaymentProcessor is the context that uses PaymentStrategy
type PaymentProcessor struct {
	strategy PaymentStrategy
}

// NewPaymentProcessor creates a new payment processor with a strategy
func NewPaymentProcessor(strategy PaymentStrategy) *PaymentProcessor {
	return &PaymentProcessor{
		strategy: strategy,
	}
}

// SetPaymentStrategy changes the payment strategy at runtime
func (pp *PaymentProcessor) SetPaymentStrategy(strategy PaymentStrategy) {
	fmt.Printf("💱 Switching to %s...\n\n", strategy.GetStrategyName())
	pp.strategy = strategy
}

// ProcessPayment processes a payment using the current strategy
func (pp *PaymentProcessor) ProcessPayment(amount float64) PaymentResult {
	if pp.strategy == nil {
		fmt.Println("❌ No payment strategy selected!")
		return PaymentResult{Success: false}
	}

	if !pp.strategy.Validate() {
		fmt.Printf("❌ Invalid %s credentials!\n\n", pp.strategy.GetStrategyName())
		return PaymentResult{Success: false}
	}

	return pp.strategy.Pay(amount)
}

// GetCurrentStrategy returns the name of the current strategy
func (pp *PaymentProcessor) GetCurrentStrategy() string {
	if pp.strategy == nil {
		return "No strategy selected"
	}
	return pp.strategy.GetStrategyName()
}

// ============================================================
// ADVANCED EXAMPLE: Shopping Cart
// ============================================================

// CartItem represents an item in the shopping cart
type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// ShoppingCart represents a shopping cart
type ShoppingCart struct {
	items              []CartItem
	paymentProcessor   *PaymentProcessor
}

// NewShoppingCart creates a new shopping cart
func NewShoppingCart() *ShoppingCart {
	return &ShoppingCart{
		items: []CartItem{},
	}
}

// AddItem adds an item to the cart
func (sc *ShoppingCart) AddItem(name string, price float64, quantity int) {
	sc.items = append(sc.items, CartItem{
		Name:     name,
		Price:    price,
		Quantity: quantity,
	})
	fmt.Printf("✅ Added %dx %s ($%.2f) to cart\n", quantity, name, price)
}

// GetTotal calculates the total price of items in the cart
func (sc *ShoppingCart) GetTotal() float64 {
	total := 0.0
	for _, item := range sc.items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

// ShowCart displays the cart contents
func (sc *ShoppingCart) ShowCart() {
	fmt.Println("\n🛒 CART CONTENTS:")
	for _, item := range sc.items {
		itemTotal := item.Price * float64(item.Quantity)
		fmt.Printf("   - %s: $%.2f x %d = $%.2f\n", item.Name, item.Price, item.Quantity, itemTotal)
	}
	fmt.Printf("   ├─ Total: $%.2f\n\n", sc.GetTotal())
}

// Checkout processes the checkout with a payment strategy
func (sc *ShoppingCart) Checkout(strategy PaymentStrategy) PaymentResult {
	fmt.Printf("📦 Proceeding to checkout with %s...\n", strategy.GetStrategyName())
	sc.paymentProcessor = NewPaymentProcessor(strategy)
	total := sc.GetTotal()
	fmt.Printf("💰 Processing payment of $%.2f:\n\n", total)
	return sc.paymentProcessor.ProcessPayment(total)
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

func main() {
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println("STRATEGY DESIGN PATTERN - Go Implementation")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	// Example 1: Payment with Credit Card
	fmt.Println("📋 EXAMPLE 1: Credit Card Payment")
	fmt.Println(strings.Repeat("-", 60))
	creditCardStrategy := &CreditCardPayment{
		CardNumber: "4532-1234-5678-9010",
		CardHolder: "John Doe",
		ExpiryDate: "12/25",
		CVV:        "123",
	}
	processor := NewPaymentProcessor(creditCardStrategy)
	processor.ProcessPayment(150.0)

	// Example 2: Switch to PayPal Strategy
	fmt.Println("📋 EXAMPLE 2: Switch to PayPal")
	fmt.Println(strings.Repeat("-", 60))
	processor.SetPaymentStrategy(&PayPalPayment{
		Email: "john.doe@example.com",
	})
	processor.ProcessPayment(75.50)

	// Example 3: Switch to Bitcoin Strategy
	fmt.Println("📋 EXAMPLE 3: Switch to Bitcoin")
	fmt.Println(strings.Repeat("-", 60))
	processor.SetPaymentStrategy(&BitcoinPayment{
		WalletAddress: "1A1z7agoat5mLrQH5r8RN85dy1eWkTqeUP",
	})
	processor.ProcessPayment(200.0)

	// Example 4: Switch to Apple Pay Strategy
	fmt.Println("📋 EXAMPLE 4: Switch to Apple Pay")
	fmt.Println(strings.Repeat("-", 60))
	processor.SetPaymentStrategy(&ApplePayPayment{
		AppleID: "john.doe@icloud.com",
	})
	processor.ProcessPayment(99.99)

	// Example 5: Switch to Google Pay Strategy
	fmt.Println("📋 EXAMPLE 5: Switch to Google Pay")
	fmt.Println(strings.Repeat("-", 60))
	processor.SetPaymentStrategy(&GooglePayPayment{
		Email: "user@gmail.com",
	})
	processor.ProcessPayment(125.0)

	// ============================================================
	// ADVANCED EXAMPLE: Shopping Cart
	// ============================================================

	fmt.Println("\n" + strings.Repeat("=", 60))
	fmt.Println("ADVANCED EXAMPLE: E-Commerce Shopping Cart")
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()

	// Simulate shopping scenario
	fmt.Println("👤 Customer Shopping...\n")
	cart := NewShoppingCart()
	cart.AddItem("Laptop", 999.99, 1)
	cart.AddItem("Mouse", 29.99, 2)
	cart.AddItem("Keyboard", 79.99, 1)

	cart.ShowCart()

	// Checkout with Credit Card
	fmt.Println("📋 CHECKOUT METHOD 1: Credit Card")
	fmt.Println(strings.Repeat("-", 60))
	cart.Checkout(&CreditCardPayment{
		CardNumber: "4532-1234-5678-9010",
		CardHolder: "Jane Smith",
		ExpiryDate: "06/26",
		CVV:        "456",
	})

	// New customer with different payment method
	fmt.Println("📋 CHECKOUT METHOD 2: PayPal")
	fmt.Println(strings.Repeat("-", 60))
	cartNew := NewShoppingCart()
	cartNew.AddItem("USB-C Cable", 15.99, 3)
	cartNew.AddItem("Monitor Stand", 45.0, 1)
	cartNew.ShowCart()
	cartNew.Checkout(&PayPalPayment{
		Email: "jane.smith@example.com",
	})
}
