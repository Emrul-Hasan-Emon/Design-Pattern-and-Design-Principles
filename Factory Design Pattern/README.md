# Factory Design Pattern

## Overview

The **Factory Design Pattern** is a creational pattern that provides an interface for creating objects without specifying the exact classes to create. It encapsulates the object creation logic, making it easy to create different types of objects based on certain conditions.

In simple terms, instead of directly instantiating classes using the `new` keyword, you use a factory function or class that handles the creation logic for you.

## Problem Statement

Imagine you're building a shipping system that supports multiple shipping carriers (FedEx, UPS, DHL, etc.). Without the Factory Pattern, you might end up with:

```javascript
if (carrierType === 'fedex') {
  const carrier = new FedExCarrier();
} else if (carrierType === 'ups') {
  const carrier = new UPSCarrier();
} else if (carrierType === 'dhl') {
  const carrier = new DHLCarrier();
}
```

This approach has several issues:
- **Scattered object creation** - Creation logic is spread throughout the application
- **Tight coupling** - Client code depends directly on concrete classes
- **Hard to maintain** - Adding new shipping carriers requires modifying client code
- **Difficult to test** - Hard to mock or test different carrier implementations
- **Violates Single Responsibility** - Client code has multiple reasons to change

## Solution: Factory Pattern

The Factory Pattern solves this by:
1. Creating a **Factory Interface or Factory Class** that handles object creation
2. Creating **Concrete Products** that implement a common interface
3. Moving **creation logic into the factory**, away from client code

## Key Components

### 1. **Product Interface**
Defines the common interface for all products created by the factory.

```javascript
interface ShippingCarrier {
  ship(package): ShippingResult;
  calculateCost(weight): number;
  getCarrierName(): string;
}
```

### 2. **Concrete Products**
Implement the Product interface. Each represents a different type of product.

```javascript
class FedExCarrier implements ShippingCarrier {
  ship(package) { /* implementation */ }
  calculateCost(weight) { /* implementation */ }
  getCarrierName() { return 'FedEx'; }
}

class UPSCarrier implements ShippingCarrier {
  ship(package) { /* implementation */ }
  calculateCost(weight) { /* implementation */ }
  getCarrierName() { return 'UPS'; }
}
```

### 3. **Factory (Creator)**
Encapsulates the object creation logic and returns products based on input parameters.

```javascript
class ShippingCarrierFactory {
  static createCarrier(carrierType) {
    switch(carrierType) {
      case 'fedex':
        return new FedExCarrier();
      case 'ups':
        return new UPSCarrier();
      case 'dhl':
        return new DHLCarrier();
      default:
        throw new Error(`Unknown carrier: ${carrierType}`);
    }
  }
}
```

### 4. **Client**
Uses the factory to create objects instead of creating them directly.

```javascript
const carrier = ShippingCarrierFactory.createCarrier('fedex');
const cost = carrier.calculateCost(5);
```

## Structure Diagram

```
╔═══════════════════════════════════════════════════════════════════════════════╗
║                        FACTORY DESIGN PATTERN                                ║
╚═══════════════════════════════════════════════════════════════════════════════╝

                            ┌──────────────────┐
                            │  Client Code     │
                            │ (Your App)       │
                            └────────┬─────────┘
                                     │
                              Creates/Uses
                                     │
                                     ▼
                        ┌────────────────────────┐
                        │  <<Factory>>           │
                        │  ShippingCarrier       │
                        │    Factory             │
                        ├────────────────────────┤
                        │ + createCarrier()      │
                        └────────────┬───────────┘
                                     │
                         Decides which to create
                       ┌─────────────┼──────────────┐
                       │             │              │
                       ▼             ▼              ▼
            ┌──────────────┐  ┌─────────────┐  ┌──────────┐
            │ FedExCarrier │  │ UPSCarrier  │  │DHLCarrier│
            ├──────────────┤  ├─────────────┤  ├──────────┤
            │ + ship()     │  │ + ship()    │  │+ship()   │
            │ + calculate()│  │ + calculate │  │+calculate│
            │ + getName()  │  │ + getName() │  │+getName()│
            └──────────────┘  └─────────────┘  └──────────┘
                       │             │              │
                       └─────────────┼──────────────┘
                                     │
                            Implements Interface
                                     │
                                     ▼
                        ┌────────────────────────┐
                        │  <<Interface>>         │
                        │  ShippingCarrier       │
                        ├────────────────────────┤
                        │ + ship()               │
                        │ + calculateCost()      │
                        │ + getCarrierName()     │
                        └────────────────────────┘
```

## How It Works

**The Flow:**

1. **Client requests a product** via the factory: `factory.createCarrier('fedex')`
2. **Factory receives the type parameter** and decides which product to create
3. **Factory encapsulates creation logic**:
   - Instantiates the correct concrete product class
   - Handles initialization if needed
   - Validates parameters
4. **Factory returns the product** to the client
5. **Client uses the product** through the common interface

### Example Workflow:

```
┌──────────────────────────┐
│ Client: "Need FedEx"     │
└────────────┬─────────────┘
             │
             ▼
┌──────────────────────────────────────┐
│ Factory.createCarrier('fedex')       │
│                                      │
│ 1. Recognize type: 'fedex'           │
│ 2. new FedExCarrier()                │
│ 3. Return instance                   │
└────────────┬─────────────────────────┘
             │
             ▼
        FedExCarrier instance
             │
             ▼
┌──────────────────────────┐
│ Client uses carrier      │
│ carrier.ship(package)    │
│ carrier.calculateCost()  │
└──────────────────────────┘
```

## Benefits

✅ **Loose Coupling** - Client code doesn't depend on concrete classes  
✅ **Single Responsibility** - Creation logic is isolated in the factory  
✅ **Open/Closed Principle** - Easy to add new products without modifying existing code  
✅ **Centralized Control** - All object creation happens in one place  
✅ **Easy to Extend** - Add new types by creating new concrete classes  
✅ **Easier Testing** - Can mock the factory or products easily  
✅ **Maintainability** - Changes to creation logic only need to happen in one place  
✅ **Flexibility** - Can add complex initialization logic in the factory  

## Drawbacks

❌ **Extra Classes** - Creates more classes for each product type  
❌ **Complexity** - Might be overkill for simple object creation  
❌ **Indirection** - Extra layer between client and actual objects  
❌ **Factory Bloat** - Factory method can become large with many products  

## When to Use

- When you have multiple types of objects that share a common interface
- When object creation logic is complex
- When the type of object to create is determined at runtime
- When you want to decouple object creation from usage
- When you want to centralize object creation
- When you need to add new types frequently without modifying client code

**Real-World Examples:**
- Database connections (MySQL, PostgreSQL, MongoDB)
- UI components (Button, TextField, CheckBox for different themes)
- Shipping carriers (FedEx, UPS, DHL)
- Payment processors (Stripe, PayPal, Square)
- Document exporters (PDF, Word, Excel)
- Image handlers (JPEG, PNG, GIF)
- Notification senders (Email, SMS, Push)

## Factory Pattern Variations

### 1. **Simple Factory (Static Method)**
```javascript
class CarrierFactory {
  static createCarrier(type) {
    // Logic here
  }
}
```
- Simplest form
- Uses static method
- Good for straightforward creation

### 2. **Factory Method (Virtual Constructor)**
```javascript
class AbstractFactory {
  createProduct() { /* override in subclass */ }
}
```
- Uses inheritance
- Each subclass decides which product to create
- More flexible for complex hierarchies

### 3. **Abstract Factory**
Creates families of related objects
- Used when you have multiple product families
- More complex but powerful

## Project Structure

```
Factory Design Pattern/
├── js/
│   └── index.js          # JavaScript implementation with class-based factory
├── go/
│   └── main.go           # Go implementation with interface-based factory
└── README.md             # This file
```

## Running the Examples

### JavaScript
```bash
cd js
node index.js
```

### Go
```bash
cd go
go run main.go
```