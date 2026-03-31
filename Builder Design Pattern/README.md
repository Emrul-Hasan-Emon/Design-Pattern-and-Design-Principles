# Builder Design Pattern

## Overview

The **Builder Design Pattern** is a creational pattern that separates the construction of a complex object from its representation. It allows you to build objects step-by-step, without requiring all parameters upfront and avoiding long constructor chains.

In simple terms, instead of creating objects with complex constructors or multiple parameters, you use a builder that constructs the object piece-by-piece.

## Problem Statement

Imagine you're building a pizza ordering system. Without the Builder Pattern, you might end up with:

```javascript
// Constructor hell
const pizza = new Pizza(
  'Thin Crust',
  true,           // has cheese
  true,           // has pepperoni
  false,          // has mushrooms
  true,           // has onions
  false,          // has olives
  'Large',
  '$18.99'
);

// Or multiple constructors
const pizza1 = new Pizza('Thin Crust', true);
const pizza2 = new Pizza('Thin Crust', true, true, false);
const pizza3 = new Pizza('Thin Crust', true, true, false, true, false);
```

This approach has several issues:
- **Parameter bloat** - Too many constructor parameters
- **Readability** - Hard to understand what each parameter means
- **Flexibility** - Must specify all parameters even if you only need a few
- **Optional parameters** - Difficult to handle optional fields
- **Immutability** - Hard to create truly immutable objects
- **Maintenance** - Adding new options requires new constructors or parameter changes

## Solution: Builder Pattern

The Builder Pattern solves this by:
1. Creating a **Builder Class** that constructs the object step-by-step
2. Providing **fluent interface** (method chaining) for easy configuration
3. **Separating construction** from the actual object representation
4. Allowing **optional fields** without multiple constructors

## Key Components

### 1. **Product**
The complex object being built.

```javascript
class Pizza {
  constructor(
    crust,
    size,
    toppings,
    cheese,
    sauce,
    price
  ) {
    this.crust = crust;
    this.size = size;
    this.toppings = toppings;
    this.cheese = cheese;
    this.sauce = sauce;
    this.price = price;
  }
}
```

### 2. **Builder**
The interface defining construction steps.

```javascript
class PizzaBuilder {
  constructor() {
    this.pizza = new Pizza();
  }

  setCrust(crust) {
    this.pizza.crust = crust;
    return this;  // Enable method chaining
  }

  addTopping(topping) {
    if (!this.pizza.toppings) this.pizza.toppings = [];
    this.pizza.toppings.push(topping);
    return this;
  }

  build() {
    return this.pizza;
  }
}
```

### 3. **Client**
Uses the builder to construct objects.

```javascript
const pizza = new PizzaBuilder()
  .setCrust('Thin Crust')
  .setSize('Large')
  .addTopping('pepperoni')
  .addTopping('mushrooms')
  .setCheese(true)
  .build();
```

## Structure Diagram

```
╔═══════════════════════════════════════════════════════════════════════════════╗
║                           BUILDER DESIGN PATTERN                             ║
╚═══════════════════════════════════════════════════════════════════════════════╝

                            ┌──────────────────┐
                            │  Client Code     │
                            │ (Your App)       │
                            └────────┬─────────┘
                                     │
                         Creates and configures
                                     │
                                     ▼
                        ┌────────────────────────┐
                        │  <<Builder>>           │
                        │  PizzaBuilder          │
                        ├────────────────────────┤
                        │ - pizza: Pizza         │
                        │ + setCrust()           │
                        │ + setSize()            │
                        │ + addTopping()         │
                        │ + setCheese()          │
                        │ + build()              │
                        │                        │
                        │ Returns: PizzaBuilder  │
                        │ (for method chaining)  │
                        └────────────┬───────────┘
                                     │
                    Step-by-step configuration
                                     │
                                     ▼
                    ┌────────────────────────────┐
                    │  Product: Pizza            │
                    ├────────────────────────────┤
                    │ - crust: String            │
                    │ - size: String             │
                    │ - toppings: List           │
                    │ - cheese: Boolean          │
                    │ - sauce: String            │
                    │ - price: Double            │
                    │                            │
                    │ + display()                │
                    │ + getPrice()               │
                    └────────────────────────────┘

                      Client calls build()
                            │
                            ▼
          Returns fully constructed Pizza object
```

## How It Works

**The Flow:**

1. **Client instantiates a Builder** with optional initial value
2. **Client calls builder methods** step-by-step to configure object
3. **Each method returns the builder** enabling method chaining (fluent interface)
4. **Client calls build()** to get the final constructed object
5. **Product is returned** ready to use

### Example Workflow:

```
┌──────────────────┐
│ new PizzaBuilder │  <- Create builder
└────────┬─────────┘
         │
         ▼
┌──────────────────────────────────┐
│ .setCrust('Thin Crust')          │  <- Configure step 1
│   Returns: this (PizzaBuilder)   │     (Enable chaining)
└────────┬─────────────────────────┘
         │
         ▼
┌──────────────────────────────────┐
│ .setSize('Large')                │  <- Configure step 2
│   Returns: this (PizzaBuilder)   │     (Enable chaining)
└────────┬─────────────────────────┘
         │
         ▼
┌──────────────────────────────────┐
│ .addTopping('pepperoni')         │  <- Configure step 3
│   Returns: this (PizzaBuilder)   │     (Enable chaining)
└────────┬─────────────────────────┘
         │
         ▼
┌──────────────────────────────────┐
│ .build()                         │  <- Final step: Build
│   Returns: Pizza (Product)       │     Return the object
└────────┬─────────────────────────┘
         │
         ▼
    Pizza Instance (Ready to use)
```

## Benefits

✅ **Readable Code** - Clear, self-documenting construction steps  
✅ **Flexibility** - Build objects with only needed parameters  
✅ **Optional Fields** - No need for multiple constructors  
✅ **Method Chaining** - Fluent interface for clean code  
✅ **Immutability** - Can create immutable objects easily  
✅ **Separation of Concerns** - Construction logic separate from representation  
✅ **Telescoping Constructor Problem** - Solves the "parameter bloat" issue  
✅ **Validation** - Can validate at each step  
✅ **Reusability** - Same builder can create multiple objects  
✅ **Maintainability** - Adding new fields doesn't break existing code  

## Drawbacks

❌ **Extra Classes** - Creates a builder class for each product  
❌ **Verbosity** - More code for simple objects  
❌ **Memory Overhead** - Extra builder object in memory  
❌ **Complexity** - Overkill for simple objects with few fields  
❌ **Thread Safety** - Builder itself may not be thread-safe  

## When to Use

- When objects have many optional parameters
- When object construction is complex with many steps
- When you want a fluent, readable API
- When you need different representations of the same data
- When you want to construct objects step-by-step
- When you want to avoid "telescoping constructor" pattern
- When building immutable objects

**Real-World Examples:**
- SQL Query builders
- HTML/XML builders
- URL query parameter builders
- Request/Response builders in APIs
- Configuration object builders
- Document builders (PDF, Word)
- Custom object construction with validation
- Game level/scene builders

## Builder Pattern Variations

### 1. **Classic Builder**
```javascript
class PizzaBuilder {
  build() { return new Pizza(this.crust, this.size, ...); }
}
```
- Returns the product from build()
- Most common variation

### 2. **Fluent Builder (Method Chaining)**
```javascript
class PizzaBuilder {
  setCrust(c) { this.crust = c; return this; }
  setSize(s) { this.size = s; return this; }
}
```
- Returns `this` from each setter
- Enables chaining

### 3. **Director Pattern with Builder**
```javascript
class PizzaDirector {
  constructor(builder) { this.builder = builder; }
  buildMargherita() {
    return this.builder.setCrust('...').addTopping('...').build();
  }
}
```
- Director orchestrates building
- Encapsulates common construction sequences

### 4. **Static Builder Pattern**
```javascript
Pizza.builder().setCrust('...').build();
```
- Static factory method returns builder

## Project Structure

```
Builder Design Pattern/
├── js/
│   └── index.js          # JavaScript implementation with fluent interface
├── go/
│   └── main.go           # Go implementation with method chaining
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