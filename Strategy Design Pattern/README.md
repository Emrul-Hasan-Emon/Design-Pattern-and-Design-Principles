# Strategy Design Pattern

## Overview

The **Strategy Design Pattern** is a behavioral pattern that defines a family of algorithms, encapsulates each one, and makes them interchangeable. It lets the algorithm vary independently from clients that use it.

In simple terms, instead of implementing a single algorithm, you define multiple algorithms and encapsulate each one in its own class. The client can then choose which algorithm to use at runtime.

## Problem Statement

Imagine you're building a payment system that supports multiple payment methods (Credit Card, PayPal, Apple Pay, etc.). Without the Strategy Pattern, you might end up with:

```
if (paymentMethod === 'creditCard') {
  // Process credit card payment
} else if (paymentMethod === 'paypal') {
  // Process PayPal payment
} else if (paymentMethod === 'applePay') {
  // Process Apple Pay payment
}
```

This approach has several issues:
- **Code becomes complex and hard to maintain** as more payment methods are added
- **Violates Open/Closed Principle** - you have to modify existing code for new strategies
- **Hard to test** - testing requires mocking multiple conditions
- **Tightly coupled** - the algorithm logic is mixed with client code

## Solution: Strategy Pattern

The Strategy Pattern solves this by:
1. Creating a **Strategy Interface** that all algorithms must implement
2. Creating **Concrete Strategy Classes** for each algorithm
3. Using a **Context Class** that uses a strategy without knowing its implementation details

## Key Components

### 1. **Strategy Interface**
Defines the common interface/contract that all concrete strategies must implement.

**TypeScript Approach (Recommended):**
```typescript
interface PaymentStrategy {
  pay(amount: number): PaymentResult;
  validate(): boolean;
  getStrategyName(): string;
}
```

**JavaScript Approach (Using Abstract Base Class):**
```javascript
class PaymentStrategy {
  pay(amount) {
    throw new Error('pay() method must be implemented');
  }

  validate() {
    throw new Error('validate() method must be implemented');
  }

  getStrategyName() {
    throw new Error('getStrategyName() method must be implemented');
  }
}
```

### 2. **Concrete Strategies**
All strategies must implement the PaymentStrategy interface and provide implementations for all required methods.

**TypeScript Approach:**
```typescript
class CreditCardPayment implements PaymentStrategy {
  constructor(
    private cardNumber: string,
    private cardHolder: string,
    private expiryDate: string,
    private cvv: string
  ) {}

  pay(amount: number): PaymentResult {
    // Process credit card payment
    return { success: true, method: 'Credit Card', amount, timestamp: new Date().toISOString() };
  }

  validate(): boolean {
    return this.cardNumber && this.cardNumber.length >= 13;
  }

  getStrategyName(): string {
    return 'Credit Card';
  }
}

// PayPalPayment, BitcoinPayment, ApplePayPayment similarly implement PaymentStrategy
```

**JavaScript Approach:**
```javascript
class CreditCardPayment extends PaymentStrategy {
  constructor(cardNumber, cardHolder, expiryDate, cvv) {
    super();
    this.cardNumber = cardNumber;
    this.cardHolder = cardHolder;
    this.expiryDate = expiryDate;
    this.cvv = cvv;
  }

  pay(amount) {
    // Process credit card payment
    return { success: true, method: 'Credit Card', amount, timestamp: new Date().toISOString() };
  }

  validate() {
    return this.cardNumber && this.cardNumber.length >= 13;
  }

  getStrategyName() {
    return 'Credit Card';
  }
}
```

### 3. **Context (Payment Processor)**
Uses a strategy object to perform the algorithm. The context works with the strategy through the interface, ensuring type safety.

**TypeScript Approach:**
```typescript
class PaymentProcessor {
  private strategy: PaymentStrategy;

  constructor(paymentStrategy: PaymentStrategy) {
    this.strategy = paymentStrategy;
  }

  setPaymentStrategy(paymentStrategy: PaymentStrategy): void {
    this.strategy = paymentStrategy;
  }
  
  processPayment(amount: number): PaymentResult {
    if (!this.strategy.validate()) {
      throw new Error('Invalid payment credentials');
    }
    return this.strategy.pay(amount);
  }
}
```

**JavaScript Approach:**
```javascript
class PaymentProcessor {
  constructor(paymentStrategy) {
    this.validateStrategy(paymentStrategy);
    this.strategy = paymentStrategy;
  }

  validateStrategy(strategy) {
    if (!(strategy instanceof PaymentStrategy)) {
      throw new Error('Strategy must be an instance of PaymentStrategy');
    }
  }
  
  processPayment(amount) {
    return this.strategy.pay(amount);
  }

  setPaymentStrategy(newStrategy) {
    this.validateStrategy(newStrategy);
    this.strategy = newStrategy;
  }
}
```

## Structure Diagram

```
╔══════════════════════════════════════════════════════════════════════════════╗
║                        STRATEGY DESIGN PATTERN                              ║
╚══════════════════════════════════════════════════════════════════════════════╝

                    ┌─────────────────────────────────────┐
                    │     <<Interface>>                   │
                    │     PaymentStrategy                 │
                    ├─────────────────────────────────────┤
                    │ + pay(amount): PaymentResult        │
                    │ + validate(): boolean               │
                    │ + getStrategyName(): string         │
                    └──────────────┬──────────────────────┘
                                   │
                                   │ implements
                                   │
              ┌────────────────────┼────────────────────┬──────────────┐
              │                    │                    │              │
        ┌─────┴────┐        ┌──────┴──────┐       ┌─────┴────┐   ┌────┴──────┐
        │  Credit  │        │   PayPal    │       │ Bitcoin  │   │  Apple    │
        │  Card    │        │  Payment    │       │ Payment  │   │   Pay     │
        │ Payment  │        │             │       │          │   │           │
        └──────────┘        └─────────────┘       └──────────┘   └───────────┘


╔═══════════════════════════════════════════════════════════════════════════════╗
║                                  CONTEXT                                      ║
╠═══════════════════════════════════════════════════════════════════════════════╣
║                     PaymentProcessor (Main Processor)                         ║
╠═══════════════════════════════════════════════════════════════════════════════╣
║                                                                               ║
║  Properties:                                                                  ║
║  ├─ strategy: PaymentStrategy                                                ║
║                                                                               ║
║  Methods:                                                                     ║
║  ├─ constructor(strategy: PaymentStrategy)                                   ║
║  ├─ setStrategy(strategy: PaymentStrategy): void                             ║
║  ├─ processPayment(amount: float): PaymentResult                             ║
║  └─ getCurrentStrategy(): string                                              ║
║                                                                               ║
╚═══════════════════════════════════════════════════════════════════════════════╝


╔═══════════════════════════════════════════════════════════════════════════════╗
║                              CLIENT INTERACTION                               ║
╚═══════════════════════════════════════════════════════════════════════════════╝

STEP 1: Create a Strategy
┌──────────────────────────────────────────────────────────┐
│ strategy = new CreditCardPayment(...)                    │
└──────────────────────────────────────────────────────────┘
                         │
                         ▼
STEP 2: Pass Strategy to Processor
┌──────────────────────────────────────────────────────────┐
│ processor = new PaymentProcessor(strategy)               │
│                                                          │
│ ==> Processor receives strategy                          │
│ ==> Processor sets it internally                         │
│ ==> Processor returns itself to client                   │
└──────────────────────────────────────────────────────────┘
                         │
                         ▼
STEP 3: Client Works with Processor
┌──────────────────────────────────────────────────────────┐
│ result = processor.processPayment(150.00)                │
│                                                          │
│ ==> Processor uses stored strategy                       │
│ ==> Calls strategy.validate()                            │
│ ==> Calls strategy.pay(150.00)                           │
│ ==> Returns result                                       │
└──────────────────────────────────────────────────────────┘
                         │
                         ▼
STEP 4: Switch Strategy at Runtime (Optional)
┌──────────────────────────────────────────────────────────┐
│ processor.setStrategy(new PayPalPayment(...))            │
│                                                          │
│ result = processor.processPayment(75.50)                 │
│                                                          │
│ ==> Processor now uses NEW strategy                      │
│ ==> Calls PayPal's validate()                            │
│ ==> Calls PayPal's pay(75.50)                            │
│ ==> Returns result                                       │
└──────────────────────────────────────────────────────────┘
                         │
                         ▼
                    Client Gets Results
```

## How It Works

The Strategy Pattern operates in a clear four-step process:

**Step 1: Interface & Strategies**
- Define the `PaymentStrategy` interface with methods like `pay()`, `validate()`, `getStrategyName()`
- Create multiple concrete strategies (CreditCard, PayPal, Bitcoin, Apple Pay, etc.)
- Each strategy implements all interface methods

**Step 2: Create Strategy**
- Client instantiates a concrete strategy: `new CreditCardPayment(...)`

**Step 3: Pass to Processor**
- Client passes the strategy to the context/processor: `new PaymentProcessor(strategy)`
- Processor receives the strategy and stores it internally
- Processor returns itself ready to use

**Step 4: Use Processor**
- Client calls methods on the processor: `processor.processPayment(amount)`
- Processor delegates to the stored strategy: `strategy.pay(amount)`
- Client can switch strategies anytime: `processor.setStrategy(newStrategy)`

### Code Example:

```javascript
// Step 1 & 2: Create a strategy
const cardStrategy = new CreditCardPayment("4532-1234-5678-9010", "John Doe", "12/25", "123");

// Step 3: Pass to processor (processor returns itself)
const processor = new PaymentProcessor(cardStrategy);

// Step 4: Client works with processor
processor.processPayment(150.00);  // Uses CreditCard strategy

// Switch strategy at runtime
processor.setStrategy(new PayPalPayment("john@example.com"));
processor.processPayment(75.50);   // Uses PayPal strategy

// No need to create a new processor - reuse the same one!
```

**Key Insight:** The client never directly calls the strategy. Instead:
- ✅ Client interacts only with the Processor
- ✅ Client passes strategies to the Processor
- ✅ Processor handles all delegation to strategies
- ✅ This keeps client code clean and simple

## Why an Interface is Essential

The **Strategy Interface** is crucial because it:

1. **Defines a Contract** - Establishes what methods every strategy must implement
2. **Ensures Consistency** - All strategies follow the same structure
3. **Enables Type Safety** - The Context can validate that a strategy implements the required methods
4. **Simplifies Adding New Strategies** - New strategies must adhere to the interface
5. **Improves Code Readability** - Developers know exactly what methods are available
6. **Supports Polymorphism** - The Context treats all strategies uniformly through the interface

**Without an interface:**
- ❌ New strategies might miss required methods
- ❌ Context code becomes fragile and error-prone
- ❌ No clear contract between context and strategies
- ❌ Difficult to catch errors at design time

**With an interface:**
- ✅ Clear contract that all strategies must follow
- ✅ Context can validate strategy compliance
- ✅ Easy to extend with new strategies
- ✅ Type checking and error detection

## Benefits

✅ **Open/Closed Principle** - Easy to add new strategies without modifying existing code  
✅ **Single Responsibility** - Each strategy handles one algorithm  
✅ **Maintainability** - Strategy logic is isolated and organized  
✅ **Testability** - Easy to test each strategy independently  
✅ **Runtime Flexibility** - Switch strategies at runtime  
✅ **Eliminates Conditionals** - No long if-else chains  

## Drawbacks

❌ **Increased Complexity** - More classes to manage  
❌ **Overhead for Simple Cases** - Overkill for simple algorithms  
❌ **Memory Cost** - Each strategy occupies memory space  

## When to Use

- When you have multiple ways to perform a task
- When you need to switch algorithms at runtime
- When you want to avoid long if-else or switch statements
- When different variants of an algorithm are needed in different scenarios
- When you want to isolate algorithm implementation details from clients

**Real-World Examples:**
- Payment processing systems (Credit Card, PayPal, Crypto)
- Sorting algorithms (QuickSort, MergeSort, BubbleSort)
- Compression algorithms (ZIP, RAR, 7-Zip)
- Routing algorithms (Shortest path, Cheapest path, Fastest route)
- Data export formats (JSON, XML, CSV)
- Authentication methods (OAuth, Basic Auth, JWT)


## Project Structure

```
Strategy Design Pattern/
├── js/
│   └── index.js              # JavaScript implementation 
├── go/
│   └── main.go               # Go implementation with implicit interfaces
└── README.md                 # This file
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
