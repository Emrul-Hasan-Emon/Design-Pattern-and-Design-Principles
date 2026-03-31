# Object-Oriented Programming (OOP) Principles

## Overview

Object-Oriented Programming (OOP) is a programming paradigm that organizes code into **objects** (data and methods). It's based on four fundamental principles that help create maintainable, scalable, and reusable code.

This section explores the **4 core OOP principles** and how they are supported in JavaScript, TypeScript, and Go.

---

## 🎯 The 4 Core OOP Principles

### 1. **Encapsulation** 🔐

**Definition**: Bundling data (attributes) and methods (functions) into a single unit (class/object) and hiding internal details from the outside world.

**Key Concepts**:
- **Data Hiding**: Keep fields private to prevent unauthorized access
- **Controlled Access**: Use getter/setter methods to manage data modifications
- **Information Hiding**: Hide implementation complexity from users
- **Class as Boundary**: Define what's public and what's private

**Benefits**:
- ✅ Prevents accidental data corruption
- ✅ Allows behavior changes without affecting external code
- ✅ Improves maintainability and debugging
- ✅ Reduces coupling between objects

**Example Concept**:
```javascript
class BankAccount {
  #balance = 0;  // Private field
  
  deposit(amount) {
    if (amount > 0) this.#balance += amount;
  }
  
  getBalance() {
    return this.#balance;
  }
}
```

---

### 2. **Inheritance** 🧬

**Definition**: Creating a new class (child/derived) based on an existing class (parent/base), allowing code reuse and hierarchical relationships.

**Key Concepts**:
- **Parent/Base Class**: Contains common properties and methods
- **Child/Derived Class**: Extends parent and adds/overrides functionality
- **Code Reuse**: Avoid writing the same code multiple times
- **IS-A Relationship**: Child IS-A type of parent
- **Method Overriding**: Child can provide its own implementation

**Benefits**:
- ✅ Promotes code reuse
- ✅ Establishes logical hierarchy
- ✅ Simplifies maintenance through centralization
- ✅ Enables polymorphism (see below)

**Example Concept**:
```javascript
class Animal {
  speak() { console.log("Some sound"); }
}

class Dog extends Animal {
  speak() { console.log("Woof!"); }  // Override
}
```

---

### 3. **Polymorphism** 🔄

**Definition**: The ability of objects to take on multiple forms or the ability of a single function/method to work with different types.

**Key Concepts**:
- **Method Overriding**: Child classes provide their own version of parent methods
- **Method Overloading**: Same method name with different parameters (limited support)
- **Interface Implementation**: Multiple classes can implement the same interface
- **Runtime Behavior**: Actual method called determined at runtime
- **Compile-time vs Runtime**: Static (compile-time) vs Dynamic (runtime) polymorphism

**Benefits**:
- ✅ Write flexible and generic code
- ✅ Write to interfaces, not implementations
- ✅ Reduces code duplication
- ✅ Makes code extensible

**Example Concept**:
```javascript
class Shape {
  area() { }
}

class Circle extends Shape {
  constructor(r) { this.r = r; }
  area() { return Math.PI * this.r * this.r; }
}

class Rectangle extends Shape {
  constructor(w, h) { this.w = w; this.h = h; }
  area() { return this.w * this.h; }
}

// Polymorphism: same method, different behavior
const shapes = [new Circle(5), new Rectangle(4, 6)];
shapes.forEach(s => console.log(s.area()));
```

---

### 4. **Abstraction** 🎭

**Definition**: Hiding complex implementation details and showing only the essential features of an object.

**Key Concepts**:
- **Abstract Classes**: Cannot be instantiated; define interface for subclasses
- **Abstract Methods**: Methods without implementation; must be provided by subclasses
- **Simplification**: Hide complexity, expose only what's necessary
- **Contract Definition**: Define what methods must exist
- **Implementation Agnostic**: User doesn't need to know how it works

**Benefits**:
- ✅ Reduces complexity for users
- ✅ Provides clear interface contracts
- ✅ Enables parallel development
- ✅ Makes code more maintainable

**Example Concept**:
```javascript
class PaymentProcessor {
  // Abstract method - must be implemented by subclasses
  process(amount) {
    throw new Error("Must implement process()");
  }
}

class CreditCardProcessor extends PaymentProcessor {
  process(amount) {
    return `Processing $${amount} with Credit Card`;
  }
}
```

---

## � Types of Encapsulation

**Encapsulation can be categorized based on access levels and implementation methods:**

### **1. Data Encapsulation**
- **Purpose**: Hide data members and control access through methods
- **Method**: Private fields + public getter/setter methods
- **Example**: Bank account hiding balance, providing deposit/withdraw methods

```typescript
class BankAccount {
  private _balance: number = 0;
  
  getBalance(): number { return this._balance; }
  deposit(amount: number): void { this._balance += amount; }
}
```

### **2. Method Encapsulation**
- **Purpose**: Hide implementation details of methods
- **Method**: Private/protected methods used only internally
- **Example**: Complex calculation hidden inside a public method

```typescript
class Calculator {
  public add(a: number, b: number): number {
    return this.internalAdd(a, b);  // Private method
  }
  
  private internalAdd(x: number, y: number): number {
    return x + y;
  }
}
```

### **3. Access Level Based**

**Public Encapsulation**
- Accessible from anywhere
- Use for essential interfaces

**Protected Encapsulation**
- Accessible within class and subclasses
- For inheritance scenarios

**Private Encapsulation**
- Only accessible within the class
- Strictest level of data hiding

### **4. Closure-Based Encapsulation** 
- Uses function scope to hide variables
- No formal access modifiers needed

```javascript
function createAccount() {
  let balance = 0;  // Hidden in closure
  
  return {
    deposit: (amount) => { balance += amount; },
    getBalance: () => balance
  };
}
```

---

## 🔍 Types of Inheritance

**Inheritance can be classified based on structure and relationships:**

### **1. Single Inheritance**
- Child class inherits from one parent class
- Simplest and most common form
- Supported by: JavaScript ✅, TypeScript ✅, Go (as embedding) ✅

```typescript
class Animal { }
class Dog extends Animal { }
```

### **2. Multi-Level Inheritance**
- Inheritance chain: Grandparent → Parent → Child
- Creates hierarchical structure
- Supported by: JavaScript ✅, TypeScript ✅

```typescript
class LivingBeing { }
class Animal extends LivingBeing { }
class Dog extends Animal { }
```

**⚠️ Warning**: Deep chains reduce maintainability.

### **3. Multiple Inheritance**
- Child inherits from multiple parents
- **NOT supported by any of these languages directly**
- ❌ JavaScript: No multiple class inheritance
- ❌ TypeScript: No multiple class inheritance (classes extend only ONE class)
- ✅ TypeScript ONLY: Can implement multiple INTERFACES (not inheritance, it's type contracts)
- ⚠️ Go: No traditional inheritance at all (uses embedding)

**TypeScript allows multiple INTERFACES (not class inheritance):**
```typescript
interface Flyable { fly(): void; }
interface Swimmable { swim(): void; }

// ✅ This is implementing multiple interfaces, NOT multiple inheritance
class Duck implements Flyable, Swimmable {
  fly() { }
  swim() { }
}

// ❌ This does NOT work (multiple class inheritance)
// class Duck extends Animal, Bird { }  // ERROR!

// ✅ But you CAN extend one class AND implement multiple interfaces
class Duck extends Animal implements Flyable, Swimmable {
  fly() { }
  swim() { }
}
```

**JavaScript alternative - use mixins:**
```javascript
// Mixins provide similar functionality without true inheritance
const canFly = {
  fly() { console.log("Flying"); }
};

const canSwim = {
  swim() { console.log("Swimming"); }
};

class Duck extends Animal { }
Object.assign(Duck.prototype, canFly, canSwim);
```

### **4. Hierarchical Inheritance**
- Multiple children inherit from same parent
- Creates tree structure
- All languages support this ✅

```typescript
class Vehicle { }
class Car extends Vehicle { }
class Bike extends Vehicle { }
class Truck extends Vehicle { }
```

### **5. Composition (Preferred Alternative)**
- Object contains instances of other classes
- Avoids inheritance complexity
- **"Favor composition over inheritance"**
- Supported by: All languages ✅

```typescript
class Engine { }
class Car {
  private engine: Engine;  // Composition
}
```

### **6. Aggregation**
- Weak relationship between objects
- Objects can exist independently
- No ownership implied (shared references)
- **"Has-a" relationship but objects are loosely coupled**
- Supported by: All languages ✅

```typescript
class Department {
  private employees: Employee[] = [];  // Aggregation
  
  addEmployee(employee: Employee) {
    this.employees.push(employee);
  }
}

class Employee {
  name: string;
  // Employee can exist without Department
}

// Employee can belong to multiple departments
const employee = new Employee();
const dept1 = new Department();
const dept2 = new Department();
dept1.addEmployee(employee);
dept2.addEmployee(employee);  // Same employee in multiple departments
```

**Key Characteristics**:
- ✅ Objects can exist independently
- ✅ No strict ownership
- ✅ Objects can be shared across multiple containers
- ✅ Weaker relationship than composition
- ✅ Can have cycles (A contains B, B contains A)

### **7. Has-A Relationship (General Composition)**
- General concept describing object containment
- **"A Car HAS-A Engine"** (composition)
- **"A Department HAS-A list of Employees"** (aggregation)
- Different from **"IS-A"** relationship (inheritance)
- Supported by: All languages ✅

**Has-A vs IS-A**:

| Aspect | IS-A (Inheritance) | HAS-A (Composition) |
|--------|-------------------|-------------------|
| **Relationship** | Parent-Child | Container-Component |
| **Example** | Dog IS-A Animal | Car HAS-A Engine |
| **Code** | `class Dog extends Animal` | `Car { Engine engine; }` |
| **Flexibility** | Tight coupling | Loose coupling |
| **Multiple** | Limited | More flexible |
| **Runtime Change** | Fixed at compile time | Can change at runtime |

**Composition vs Aggregation Comparison**:

| Aspect | Composition | Aggregation |
|--------|-------------|-------------|
| **Ownership** | Strong (owns parts) | Weak (references parts) |
| **Lifecycle** | Parts die with whole | Parts can exist independently |
| **Relationship Type** | Part-of | Member-of |
| **Example** | Car HAS-A Engine | Department HAS-A Employee |
| **Multiplicity** | Usually 1:1 or 1:N | Usually 1:N or N:M |

**Composition Example** (Strong Ownership):
```typescript
class Car {
  private engine: Engine;  // Owned by Car
  
  constructor() {
    this.engine = new Engine();  // Created here
  }
  
  // When Car is destroyed, Engine is destroyed
}
```

**Aggregation Example** (Weak Ownership):
```typescript
class Team {
  private members: Developer[] = [];  // Referenced, not owned
  
  addMember(developer: Developer) {
    this.members.push(developer);
  }
  
  // When Team is destroyed, Developers still exist
}

const dev = new Developer("John");
const team1 = new Team();
const team2 = new Team();
team1.addMember(dev);
team2.addMember(dev);  // Same developer in multiple teams
```

### **8. Embedding (Go's Approach)**
- Promotes fields and methods of embedded type
- Alternative to inheritance
- Go's preferred method

```go
type Animal struct { Name string }
type Dog struct {
  Animal  // Embedded
  Breed   string
}
```

### **9. Mixins (JavaScript/TypeScript Alternative to Multiple Inheritance)**
- **Purpose**: Share functionality across multiple classes/objects without inheritance
- **Method**: Object composition + method sharing
- **Philosophy**: "Behavior sharing" instead of "IS-A relationship"
- **Key Benefit**: Avoids deep inheritance chains, allows code reuse without inheritance
- Supported by: JavaScript ✅, TypeScript ✅, Go (via embedding) ✅

**What is a Mixin?**
- A plain object containing methods/properties
- Mixed into other objects/classes to share functionality
- NOT inheritance (no parent-child relationship)
- Multiple mixins can be combined

**JavaScript Mixin Example:**
```javascript
// Define reusable behavior as mixins
const canFly = {
  fly() { console.log(`${this.name} is flying`); }
};

const canSwim = {
  swim() { console.log(`${this.name} is swimming`); }
};

const canRun = {
  run() { console.log(`${this.name} is running`); }
};

// Base class
class Animal {
  constructor(name) {
    this.name = name;
  }
}

// Apply multiple mixins to a class
class Duck extends Animal {}
Object.assign(Duck.prototype, canFly, canSwim, canRun);

// Use it
const duck = new Duck("Donald");
duck.fly();   // Donald is flying
duck.swim();  // Donald is swimming
duck.run();   // Donald is running
```

**TypeScript Mixin Example (with type safety):**
```typescript
// Mixin functions that add functionality
function canFly<T extends { name: string }>(constructor: T) {
  return class extends constructor {
    fly() { console.log(`${this.name} is flying`); }
  };
}

function canSwim<T extends { name: string }>(constructor: T) {
  return class extends constructor {
    swim() { console.log(`${this.name} is swimming`); }
  };
}

// Apply mixins
class Animal {
  constructor(public name: string) {}
}

class Duck extends canSwim(canFly(Animal)) {}

const duck = new Duck("Donald");
duck.fly();   // Donald is flying
duck.swim();  // Donald is swimming
```

**Alternative TypeScript Mixin (using interfaces):**
```typescript
// Define interfaces for behaviors
interface Swimmer {
  swim(): void;
}

interface Flyer {
  fly(): void;
}

// Mixin implementations
const swimMixin: Swimmer = {
  swim() { console.log("Swimming..."); }
};

const flyMixin: Flyer = {
  fly() { console.log("Flying..."); }
};

class Duck implements Swimmer, Flyer {
  name: string;
  
  constructor(name: string) {
    this.name = name;
  }
  
  // Manually include mixin methods (or use Object.assign)
  swim = swimMixin.swim;
  fly = flyMixin.fly;
}
```

**Go Composition (similar to mixins):**
```go
type Flyer interface {
  Fly() string
}

type Swimmer interface {
  Swim() string
}

// Mixin-like embedded types
type Bird struct {
  Name string
}

func (b Bird) Fly() string {
  return fmt.Sprintf("%s is flying", b.Name)
}

type Fish struct {
  Name string
}

func (f Fish) Swim() string {
  return fmt.Sprintf("%s is swimming", f.Name)
}

// Combine via embedding
type Duck struct {
  Bird
  Fish
}

func main() {
  duck := Duck{Bird{"Donald"}, Fish{"Donald"}}
  fmt.Println(duck.Bird.Fly())
  fmt.Println(duck.Fish.Swim())
}
```

**Mixins vs Inheritance:**

| Aspect | Inheritance | Mixins |
|--------|------------|--------|
| **Relationship** | IS-A (Parent-Child) | Behavior sharing |
| **Structure** | Class hierarchy | Flat, combining objects |
| **Multiple** | Limited | Multiple allowed |
| **Flexibility** | Fixed at class definition | Can be applied dynamically |
| **Code Reuse** | Through hierarchy | Direct method sharing |
| **Complexity** | Can get deep chains | Simpler, more modular |
| **Example** | Dog extends Animal | Duck gets canFly, canSwim |

**When to Use Mixins:**
- ✅ Multiple unrelated behaviors to share
- ✅ Avoid deep inheritance hierarchies
- ✅ Share behavior across unrelated classes
- ✅ Dynamic behavior addition

**When NOT to Use Mixins:**
- ❌ When IS-A relationship truly exists
- ❌ When you need clear hierarchy
- ❌ Performance critical code (extra indirection)

---

## 🔍 Types of Polymorphism

**Polymorphism is divided into static (compile-time) and dynamic (runtime):**

### **1. Compile-Time (Static) Polymorphism**

#### **Method Overloading**
- Same method name with different parameters
- Resolved at compile time
- Supported by: TypeScript ✅ (through unions/overloads), Go ❌, JavaScript ❌

```typescript
class Calculator {
  add(a: number, b: number): number;
  add(a: string, b: string): string;
  add(a: any, b: any) {
    return a + b;
  }
}
```

#### **Operator Overloading**
- Redefining operators for custom types
- Limited support in most languages
- Supported by: TypeScript (limited) ⚠️, Go ❌, JavaScript ❌

### **2. Runtime (Dynamic) Polymorphism**

#### **Method Overriding**
- Child class provides different implementation of parent method
- Decided at runtime based on actual object type
- Supported by: JavaScript ✅, TypeScript ✅, Go ✅

```typescript
class Animal {
  speak(): void { console.log("Sound"); }
}

class Dog extends Animal {
  speak(): void { console.log("Woof!"); }  // Override
}

const animal: Animal = new Dog();
animal.speak();  // Calls Dog's version
```

#### **Interface-Based Polymorphism**
- Different classes implement same interface
- Methods called based on actual type
- Supported by: TypeScript ✅, Go ✅, JavaScript ❌

```typescript
interface Shape {
  area(): number;
}

class Circle implements Shape { area() { } }
class Rectangle implements Shape { area() { } }
```

#### **Duck Typing (JavaScript)**
- "If it walks like a duck and quacks like a duck, it's a duck"
- No formal interface needed
- Type determined by behavior, not declaration

```javascript
const makeSound = (creature) => creature.speak();

makeSound(dog);   // Works if dog has speak()
makeSound(person); // Works if person has speak()
```

### **3. Parametric Polymorphism (Generics)**
- Using type parameters for generic types
- Supported by: TypeScript ✅, Go ✅, JavaScript ❌

```typescript
class Container<T> {
  item: T;
  
  getItem(): T { return this.item; }
}

const stringContainer = new Container<string>();
```

---

## 🔍 Types of Abstraction

**Abstraction methods and levels:**

### **1. Abstract Classes**
- Cannot be instantiated directly
- Define interface for subclasses
- Contain abstract methods (must be implemented)
- **⚠️ IMPORTANT: Abstract classes are TypeScript compile-time only**
- ❌ JavaScript: No true abstract classes (after compilation, becomes normal class)
- ✅ TypeScript: Supports abstract keyword
- ✅ Go: Uses interfaces instead

**TypeScript Abstract Class (compile-time feature):**
```typescript
abstract class PaymentProcessor {
  abstract process(amount: number): boolean;
  
  log() { console.log("Processing payment..."); }
}

class CreditCardProcessor extends PaymentProcessor {
  process(amount: number): boolean { return true; }
}
```

**After JavaScript Compilation (abstract disappears):**
```javascript
// Abstract keyword is GONE - just a regular class now
class PaymentProcessor {
  process(amount) {
    throw new Error("Must implement process()");
  }
  
  log() { console.log("Processing payment..."); }
}

class CreditCardProcessor extends PaymentProcessor {
  process(amount) { return true; }
}
```

**JavaScript workaround (convention-based):**
```javascript
// JavaScript doesn't support abstract, so throw error manually
class PaymentProcessor {
  process() {
    throw new Error("Must implement process() method");
  }
}

class CreditCardProcessor extends PaymentProcessor {
  process(amount) { return true; }
}
```

### **2. Interfaces**
- Define contract/protocol
- Specify what methods must exist
- Classes implement interfaces
- **⚠️ IMPORTANT: Interfaces are TypeScript compile-time only**
- ❌ JavaScript: No interfaces (after compilation, they vanish)
- ✅ TypeScript: Full interface support
- ✅ Go: Full interface support (implicit implementation)

**TypeScript Interfaces (compile-time, erased after compilation):**
```typescript
interface PaymentMethod {
  pay(amount: number): boolean;
  validate(): boolean;
}

class CreditCard implements PaymentMethod {
  pay(amount: number): boolean { return true; }
  validate(): boolean { return true; }
}
```

**After JavaScript Compilation (interface disappears):**
```javascript
// The interface VANISHES - JavaScript has no interfaces
class CreditCard {
  pay(amount) { return true; }
  validate() { return true; }
}
```

**JavaScript workaround (duck typing):**
```javascript
// JavaScript doesn't enforce interfaces, use duck typing
function processPayment(paymentMethod) {
  if (typeof paymentMethod.pay === 'function' && 
      typeof paymentMethod.validate === 'function') {
    paymentMethod.validate();
    paymentMethod.pay(100);
  }
}

// Any object with these methods works
const creditCard = {
  pay(amount) { console.log(`Paid $${amount}`); },
  validate() { console.log("Validated"); }
};

processPayment(creditCard);  // Works!
```

### **3. Abstract Methods**
- Methods declared but not implemented in parent
- Child classes must provide implementation
- Enforces contract
- Supported by: TypeScript ✅

```typescript
abstract class Animal {
  abstract makeSound(): void;  // No implementation
}
```

### **4. Method Hiding**
- Keeping implementation details private
- Exposing only necessary public interface
- Reduces complexity for users
- Supported by: All languages ✅

```typescript
class Database {
  public connect() { this.openConnection(); }
  
  private openConnection() { } // Hidden
  private authenticate() { }    // Hidden
}
```

### **5. Inheritance-Based Abstraction**
- Parent class defines abstract operations
- Child classes provide specific implementation
- Creates abstraction hierarchy

```typescript
abstract class Shape {
  abstract calculateArea(): number;
  abstract draw(): void;
}

class Circle extends Shape {
  calculateArea() { }
  draw() { }
}
```

### **6. Protocol/Contract Abstraction**
- Define expected behavior without implementation
- Used heavily in Go (implicit interfaces)
- Reduces coupling between types

```go
type Reader interface {
  Read(p []byte) (n int, err error)
}
// Any type with this method signature
// implicitly becomes a Reader
```

---

## 📊 Language Support Summary Table

| Feature | JavaScript | TypeScript | Go |
|---------|-----------|-----------|-----|
| **Encapsulation** | ⚠️ Closures | ✅ All types | ✅ Capitalization |
| **Single Inheritance** | ✅ | ✅ | ✅ (Embedding) |
| **Multiple Inheritance** | ❌ | ❌ | ❌ |
| **Composition** | ✅ | ✅ | ✅ |
| **Aggregation** | ✅ | ✅ | ✅ |
| **Has-A Relationship** | ✅ | ✅ | ✅ |
| **Mixins** | ✅ YES | ✅ YES | ✅ (via embedding) |
| **Method Overloading** | ❌ | ✅ | ❌ |
| **Method Overriding** | ✅ | ✅ | ✅ |
| **Abstract Classes** | ⚠️ | ✅ | ❌ |
| **Interfaces** | ❌ | ✅ | ✅ (Implicit) |
| **Access Modifiers** | ⚠️ Limited | ✅ Full | ✅ Convention |
| **Generics** | ❌ | ✅ | ✅ |
| **Duck Typing** | ✅ | ⚠️ | ❌ |

---

## 🎯 CRITICAL: JavaScript Runtime vs TypeScript Compile-Time

### ⚠️ The Big Rule You Must Remember

**JavaScript runtime has NO concept of:**
- ❌ `interface`
- ❌ `abstract class`
- ❌ `multiple inheritance`

**JavaScript runtime ONLY knows:**
- ✅ Functions
- ✅ Objects
- ✅ Prototypes (single prototype chain)
- ✅ Single inheritance

**Everything else is TypeScript compile-time only!**

---

## 📋 Correct TypeScript Summary

### ✅ What TypeScript Supports

| Feature | Support | Example |
|---------|---------|---------|
| **Interfaces** | ✅ Yes | `interface Shape { area(): number; }` |
| **AbstractClasses** | ✅ Yes | `abstract class Animal {}` |
| **Extend ONE class** | ✅ Yes | `class Dog extends Animal {}` |
| **Extend multiple classes** | ❌ NO | `class Dog extends Animal, Mammal {}` ❌ |
| **Implement multiple interfaces** | ✅ Yes | `class Bird implements Flyable, Swimmable {}` |
| **Extend 1 class + multiple interfaces** | ✅ Yes | `class Bird extends Animal implements Flyable, Swimmable {}` |

### 🔑 TypeScript Final Rule

```
A TypeScript class can:
  ✅ extend ONE class (normal or abstract)
  ✅ implement MULTIPLE interfaces
  ✅ extend one class AND implement multiple interfaces simultaneously

Examples:
✅ class Dog extends Animal {}
✅ class Bird implements Flyable, Walkable, Swimmable {}
✅ class Bird extends Animal implements Flyable, Walkable, Swimmable {}

❌ class Dog extends Animal, Mammal {}
```

---

## 📋 Correct JavaScript Summary

### ✅ What JavaScript Supports

| Feature | Support | Example |
|---------|---------|---------|
| **Interfaces** | ❌ NO | Not in JavaScript |
| **Abstract classes** | ❌ NO | Compile-time only in TS |
| **Extend ONE class** | ✅ Yes | `class Dog extends Animal {}` |
| **Extend multiple classes** | ❌ NO | `class Dog extends Animal, Mammal {}` ❌ |
| **Multiple inheritance** | ❌ NO | Prototype chain is linear |

### 🔑 JavaScript Final Rule

```
JavaScript supports:
  ✅ Single inheritance only
  ✅ One class can extend ONE class
  ❌ No interfaces
  ❌ No abstract classes
  ❌ No multiple inheritance

Why? Prototype chain is LINEAR:
  Dog → Animal → Object
  (Cannot branch to multiple parents)
```

### ⚠️ What Happens to TypeScript Features in JavaScript?

**At Runtime (after TypeScript compilation):**

| TypeScript | Compiles to JavaScript | JavaScript sees |
|-----------|----------------------|-----------------|
| `interface Shape { }` | (removed) | Nothing (erased) |
| `abstract class Animal { }` | `class Animal { }` | Normal class |
| `class Dog implements Shape { }` | `class Dog { }` | Just a class |
| `class Dog extends Animal implements Shape { }` | `class Dog extends Animal { }` | Single inheritance |

**Example:**
```typescript
// TypeScript Source
abstract class PaymentProcessor {
  abstract process(): void;
}

class CreditCard implements PaymentProcessor {
  process() { }
}
```

```javascript
// Compiled JavaScript (interfaces/abstract removed)
class PaymentProcessor {
  // No "abstract" - just a regular class
}

class CreditCard extends PaymentProcessor {
  process() { }
}
```

---

## 🎭 Comparison: JavaScript vs TypeScript OOP

### Full Feature Comparison

| Feature | JavaScript | TypeScript | Notes |
|---------|-----------|-----------|-------|
| **Interfaces** | ❌ No | ✅ Yes (compile-time) | Erased after compilation |
| **Abstract Classes** | ❌ No | ✅ Yes (compile-time) | Become regular classes in JS |
| **Encapsulation** | ⚠️ Partial | ✅ Full | JS: closures / TS: access modifiers |
| **Single Inheritance** | ✅ Yes | ✅ Yes | Only one parent allowed |
| **Multiple Inheritance** | ❌ No | ❌ No | Never supported |
| **Multiple Interfaces** | ❌ No | ✅ Yes | TS only, removed in JS |
| **Access Modifiers** | ❌ Limited | ✅ Full | public/private/protected |
| **Type System** | ❌ No | ✅ Yes | Runtime erased in JS |
| **Method Overloading** | ❌ No | ✅ Yes | Compile-time checking |
| **Generics** | ❌ No | ✅ Yes | Type parameters |

### 🎓 Practical Examples

#### Example 1: Trying Multiple Inheritance

**JavaScript (Only single inheritance allowed):**
```javascript
// ❌ This DOES NOT work in JavaScript
class Dog extends Animal, Mammal { }

// ✅ This is what you do instead
class Dog extends Animal {
  // Copy Mammal functionality manually or use mixins
}
```

**TypeScript (Same - classes can't extend multiple classes):**
```typescript
// ❌ This does NOT work in TypeScript either
class Dog extends Animal, Mammal { }

// ✅ But you can implement multiple interfaces
class Dog extends Animal implements Swimmable, Runnable { }
```

#### Example 2: Abstract Classes in TypeScript → JavaScript

**TypeScript Source:**
```typescript
abstract class Animal {
  abstract makeSound(): void;
  
  move() { console.log("Moving..."); }
}

class Dog extends Animal {
  makeSound() { console.log("Woof!"); }
}
```

**After JavaScript Compilation:**
```javascript
// Interfaces/abstract are GONE - JS doesn't know about them
class Animal {
  move() { console.log("Moving..."); }
}

class Dog extends Animal {
  makeSound() { console.log("Woof!"); }
}
```

**At Runtime in JavaScript:**
```javascript
// JavaScript enforces nothing automatically
// If Dog doesn't have makeSound(), that's OK at runtime
// But TypeScript would catch it at compile time
```

#### Example 3: Interfaces in TypeScript → JavaScript

**TypeScript Source:**
```typescript
interface Shape {
  area(): number;
  perimeter(): number;
}

class Circle implements Shape {
  area() { return Math.PI * this.r * this.r; }
  perimeter() { return 2 * Math.PI * this.r; }
}
```

**After JavaScript Compilation:**
```javascript
// The interface DISAPPEARS - it's not part of JavaScript
class Circle {
  area() { return Math.PI * this.r * this.r; }
  perimeter() { return 2 * Math.PI * this.r; }
}
```

**At Runtime:**
```javascript
// JavaScript has no idea that Circle "implements" Shape
// It just sees a class with those methods
```

---

## ✨ The Golden Rule Summary

| Aspect | TypeScript | JavaScript |
|--------|-----------|-----------|
| **Compile-Time** | ✅ Interfaces ✅ Abstract classes ✅ Type checking | ❌ None (no compilation) |
| **Runtime** | Erases interfaces/abstract → runs as JS | ✅ Classes ✅ Prototypes ✅ Single inheritance |
| **One class extension** | ✅ Yes | ✅ Yes |
| **Multiple interfaces** | ✅ Yes | ❌ They don't exist |
| **Multiple class inheritance** | ❌ Never | ❌ Never |
| **What matters** | Type safety before runtime | What actually executes |

### 🎯 Simple Rule to Remember

```
TypeScript = JavaScript + Type Checking + OOP Features (compile-time)
JavaScript = Just functions, objects, and prototypes (at runtime)

When TypeScript compiles to JavaScript:
  - All interfaces vanish
  - All abstract keywords vanish
  - All type hints vanish
  - Only plain JavaScript classes remain
```

---

## 💡 Practical Advice

### ✅ DO (In TypeScript):
```typescript
// Use interfaces for contracts
interface PaymentMethod {
  pay(): void;
}

// Use abstract classes for base implementations
abstract class BaseProcessor {
  abstract process(): void;
}

// Implement multiple interfaces
class CreditCard implements PaymentMethod, Serializable { }

// Extend one class + implement interfaces
class CreditCard extends BaseProcessor implements PaymentMethod { }
```

### ❌ DON'T:
```typescript
// DON'T try to extend multiple classes
class Dog extends Animal, Mammal { }  // ❌ Error

// DON'T assume JavaScript has interfaces
// JavaScript sees none of this:
interface Shape { }  // Just type checking, removed at runtime!
```

### 📝 JavaScript Workarounds:

If you need interface-like behavior in vanilla JavaScript:

```javascript
// Use duck typing
function makeSound(creature) {
  if (typeof creature.speak === 'function') {
    creature.speak();
  }
}

// Use mixins for shared functionality
const canFly = {
  fly() { console.log("Flying..."); }
};

class Bird extends Animal { }
Object.assign(Bird.prototype, canFly);
```

---

## 🎓 Summary Table: What's Actually True



| Requirement | JavaScript | TypeScript | Go |
|---|---|---|---|
| **Strict Encapsulation** | ⚠️ | ✅ | ✅ |
| **Type Safety** | ❌ | ✅ | ✅ |
| **Interface Contracts** | ❌ | ✅ | ✅ |
| **Easy Learning Curve** | ✅ | ⚠️ | ✅ |
| **Rapid Development** | ✅ | ⚠️ | ✅ |
| **Large Projects** | ⚠️ | ✅ | ✅ |
| **Simplicity** | ✅ | ❌ | ✅ |
| **Web Development** | ✅ | ✅ | ⚠️ |
| **Systems Programming** | ❌ | ❌ | ✅ |
| **Flexible Typing** | ✅ | ⚠️ | ❌ |

---

## 🎓 OOP Principles Application Examples

### **Example 1: Payment System**

**Encapsulation**: Hide payment details, expose only necessary methods
**Inheritance**: Different payment methods inherit from base
**Polymorphism**: Same interface, different implementations
**Abstraction**: User doesn't care how payment is processed

See: [Strategy Design Pattern](../Strategy%20Design%20Pattern/README.md)

### **Example 2: Object Creation**

**Encapsulation**: Constructor hides initialization logic
**Abstraction**: Hide object creation complexity
**Polymorphism**: Different ways to create objects

See: [Factory Design Pattern](../Factory%20Design%20Pattern/README.md)

### **Example 3: Complex Configuration**

**Encapsulation**: Protect configuration state
**Abstraction**: Hide builder complexity
**Polymorphism**: Build different object types

See: [Builder Design Pattern](../Builder%20Design%20Pattern/README.md)

---

## 💡 Best Practices

### ✅ DO:
1. **Use encapsulation** to protect data integrity
2. **Favor composition over inheritance** (reduces complexity)
3. **Write to interfaces, not implementations**
4. **Keep classes single-responsibility** (one reason to change)
5. **Use meaningful names** for classes and methods
6. **Document public interfaces** clearly

### ❌ DON'T:
1. **Deep inheritance chains** (hard to follow)
2. **Expose internal state** unnecessarily
3. **Violate encapsulation** through backdoors
4. **Create "god objects"** (classes doing too much)
5. **Use inheritance for code reuse only** (use composition instead)
6. **Forget about polymorphism benefits** (write flexible code)

---

## 🔗 Related Resources

- [Encapsulation Examples](./Encapsulation/)
- [Inheritance Examples](./Inheritance/)
- [Polymorphism Examples](./Polymorphism/)
- [Abstraction Examples](./Abstraction/)

---


### 🎯 What This Means:

**JavaScript:**
- ✅ Can use: Single inheritance, encapsulation (closures/private fields), polymorphism via duck typing
- ❌ NO: Interfaces, abstract classes, multiple inheritance, type checking at runtime

**TypeScript:**
- ✅ Can use: Everything TypeScript (interfaces, abstract classes, multiple interfaces, type checking)
- ⚠️ But: All compile-time features are erased when compiled to JavaScript

**Go:**
- ✅ Can use: Single inheritance (via embedding), interfaces, implicit implementation
- ❌ NO: Abstract classes, explicit inheritance (uses composition principle instead)

---

