
// ============================================================
// PRODUCT - Complex Object
// ============================================================

/**
 * Pizza Product Class
 * Represents a pizza with multiple optional and required fields
 */
class Pizza {
  constructor() {
    this.crust = 'Regular';
    this.size = 'Medium';
    this.sauce = 'Tomato';
    this.toppings = [];
    this.cheese = true;
    this.extraCheese = false;
    this.isSpicy = false;
    this.price = 0;
  }

  display() {
    console.log('\n🍕 PIZZA DETAILS:');
    console.log(`   Crust: ${this.crust}`);
    console.log(`   Size: ${this.size}`);
    console.log(`   Sauce: ${this.sauce}`);
    console.log(`   Toppings: ${this.toppings.length > 0 ? this.toppings.join(', ') : 'None'}`);
    console.log(`   Cheese: ${this.cheese ? 'Yes' : 'No'}${this.extraCheese ? ' (Extra)' : ''}`);
    console.log(`   Spicy: ${this.isSpicy ? 'Yes 🌶️' : 'No'}`);
    console.log(`   Price: $${this.price.toFixed(2)}\n`);
  }

  getDescription() {
    return `${this.size} ${this.crust} Pizza with ${this.toppings.length} toppings`;
  }
}

// ============================================================
// BUILDER - Step-by-step Construction with Fluent Interface
// ============================================================

/**
 * PizzaBuilder Class
 * Uses fluent interface to build Pizza step-by-step
 * Each method returns 'this' to enable method chaining
 */
class PizzaBuilder {
  constructor() {
    this.pizza = new Pizza();
  }

  /**
   * Set the crust type
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  setCrust(crust) {
    const validCrusts = ['Regular', 'Thin', 'Thick', 'Stuffed Crust', 'Cauliflower'];
    if (!validCrusts.includes(crust)) {
      throw new Error(`Invalid crust type: ${crust}. Must be one of: ${validCrusts.join(', ')}`);
    }
    this.pizza.crust = crust;
    return this; // Enable method chaining
  }

  /**
   * Set the size
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  setSize(size) {
    const validSizes = ['Small', 'Medium', 'Large', 'X-Large', '2X-Large'];
    if (!validSizes.includes(size)) {
      throw new Error(`Invalid size: ${size}. Must be one of: ${validSizes.join(', ')}`);
    }
    this.pizza.size = size;
    
    // Auto-calculate price based on size
    const priceMap = { Small: 8, Medium: 10, Large: 12, 'X-Large': 14, '2X-Large': 16 };
    this.pizza.price = priceMap[size];
    
    return this;
  }

  /**
   * Set the sauce type
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  setSauce(sauce) {
    const validSauces = ['Tomato', 'White', 'BBQ', 'Pesto', 'Buffalo'];
    if (!validSauces.includes(sauce)) {
      throw new Error(`Invalid sauce: ${sauce}. Must be one of: ${validSauces.join(', ')}`);
    }
    this.pizza.sauce = sauce;
    return this;
  }

  /**
   * Add a topping
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  addTopping(topping) {
    const validToppings = [
      'Pepperoni', 'Mushrooms', 'Onions', 'Sausage',
      'Bacon', 'Chicken', 'Olives', 'Bell Peppers',
      'Spinach', 'Feta', 'Jalapeños'
    ];
    
    if (!validToppings.includes(topping)) {
      throw new Error(`Invalid topping: ${topping}. Must be one of: ${validToppings.join(', ')}`);
    }
    
    // Avoid duplicate toppings
    if (!this.pizza.toppings.includes(topping)) {
      this.pizza.toppings.push(topping);
      this.pizza.price += 1.5; // $1.50 per topping
    }
    
    return this;
  }

  /**
   * Add multiple toppings at once
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  addToppings(toppings) {
    if (!Array.isArray(toppings)) {
      throw new Error('Toppings must be an array');
    }
    
    toppings.forEach(topping => this.addTopping(topping));
    return this;
  }

  /**
   * Set whether to include cheese
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  setCheese(hasCheese) {
    this.pizza.cheese = hasCheese;
    return this;
  }

  /**
   * Add extra cheese
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  addExtraCheese() {
    this.pizza.extraCheese = true;
    this.pizza.price += 2.0; // $2.00 for extra cheese
    return this;
  }

  /**
   * Make the pizza spicy
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  makeSpicy() {
    this.pizza.isSpicy = true;
    this.addTopping('Jalapeños'); // Add jalapeños if not already there
    return this;
  }

  /**
   * Apply a preset configuration
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  applyPreset(preset) {
    switch (preset.toLowerCase()) {
      case 'margherita':
        this.setCrust('Regular')
          .setSize('Medium')
          .setSauce('Tomato')
          .setCheese(true);
        break;

      case 'pepperoni':
        this.setCrust('Regular')
          .setSize('Medium')
          .setSauce('Tomato')
          .addTopping('Pepperoni')
          .setCheese(true);
        break;

      case 'vegetarian':
        this.setCrust('Regular')
          .setSize('Medium')
          .setSauce('Tomato')
          .addToppings(['Mushrooms', 'Onions', 'Bell Peppers', 'Olives'])
          .setCheese(true);
        break;

      case 'spicy':
        this.setCrust('Thick')
          .setSize('Large')
          .setSauce('Buffalo')
          .addToppings(['Pepperoni', 'Jalapeños', 'Sausage'])
          .addExtraCheese()
          .makeSpicy();
        break;

      case 'deluxe':
        this.setCrust('Stuffed Crust')
          .setSize('Large')
          .setSauce('Tomato')
          .addToppings(['Pepperoni', 'Sausage', 'Mushrooms', 'Onions'])
          .addExtraCheese();
        break;

      default:
        throw new Error(`Unknown preset: ${preset}`);
    }
    return this;
  }

  /**
   * Validate the pizza configuration
   * @private
   */
  validate() {
    if (!this.pizza.crust) throw new Error('Crust must be set');
    if (!this.pizza.size) throw new Error('Size must be set');
    if (!this.pizza.sauce) throw new Error('Sauce must be set');
    if (this.pizza.toppings.length === 0 && !this.pizza.cheese) {
      throw new Error('Pizza must have either toppings or cheese');
    }
  }

  /**
   * Build and return the final Pizza object
   * @returns {Pizza} - The constructed Pizza
   */
  build() {
    this.validate();
    return this.pizza;
  }

  /**
   * Reset the builder to create a new pizza
   * @returns {PizzaBuilder} - Returns this for method chaining
   */
  reset() {
    this.pizza = new Pizza();
    return this;
  }
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

console.log('='.repeat(70));
console.log('BUILDER DESIGN PATTERN - Pizza Ordering System');
console.log('='.repeat(70));
console.log();

/**
 * Example 1: Basic Pizza with Method Chaining
 */
console.log('📋 EXAMPLE 1: Build a Basic Pepperoni Pizza');
console.log('-'.repeat(70));

const pizza1 = new PizzaBuilder()
  .setCrust('Regular')
  .setSize('Large')
  .setSauce('Tomato')
  .addTopping('Pepperoni')
  .setCheese(true)
  .build();

pizza1.display();

/**
 * Example 2: Complex Pizza with Multiple Toppings
 */
console.log('📋 EXAMPLE 2: Build a Complex Pizza with Multiple Toppings');
console.log('-'.repeat(70));

const pizza2 = new PizzaBuilder()
  .setCrust('Stuffed Crust')
  .setSize('X-Large')
  .setSauce('Tomato')
  .addTopping('Pepperoni')
  .addTopping('Sausage')
  .addTopping('Mushrooms')
  .addTopping('Onions')
  .addExtraCheese()
  .build();

pizza2.display();

/**
 * Example 3: Vegetarian Pizza using addToppings
 */
console.log('📋 EXAMPLE 3: Build a Vegetarian Pizza');
console.log('-'.repeat(70));

const pizza3 = new PizzaBuilder()
  .setCrust('Thin')
  .setSize('Medium')
  .setSauce('Pesto')
  .addToppings(['Mushrooms', 'Onions', 'Bell Peppers', 'Spinach'])
  .setCheese(true)
  .build();

pizza3.display();

/**
 * Example 4: Spicy Pizza with Helper Method
 */
console.log('📋 EXAMPLE 4: Build a Spicy Pizza');
console.log('-'.repeat(70));

const pizza4 = new PizzaBuilder()
  .setCrust('Thick')
  .setSize('Large')
  .setSauce('Buffalo')
  .addToppings(['Pepperoni', 'Jalapeños', 'Sausage'])
  .makeSpicy()
  .addExtraCheese()
  .build();

pizza4.display();

/**
 * Example 5: Using Presets
 */
console.log('📋 EXAMPLE 5: Build Pizza Using Presets');
console.log('-'.repeat(70));

const presets = ['Margherita', 'Pepperoni', 'Vegetarian', 'Spicy', 'Deluxe'];

console.log('Available Presets:');
const builder = new PizzaBuilder();

for (const preset of presets) {
  const pizza = builder.reset().applyPreset(preset).build();
  console.log(`✓ ${pizza.crust} Crust ${preset}: $${pizza.price.toFixed(2)}`);
}
console.log();

/**
 * Example 6: Custom Pizza Combination
 */
console.log('📋 EXAMPLE 6: Custom Pizza with White Sauce');
console.log('-'.repeat(70));

const pizza6 = new PizzaBuilder()
  .setCrust('Cauliflower')
  .setSize('Medium')
  .setSauce('White')
  .addToppings(['Chicken', 'Spinach', 'Feta'])
  .addExtraCheese()
  .build();

pizza6.display();

// ============================================================
// ADVANCED EXAMPLE: Restaurant Order Management
// ============================================================

console.log('\n' + '='.repeat(70));
console.log('ADVANCED EXAMPLE: Restaurant Order Management System');
console.log('='.repeat(70));
console.log();

/**
 * Order Class for managing multiple pizzas
 */
class Order {
  constructor(orderId, customerName) {
    this.orderId = orderId;
    this.customerName = customerName;
    this.pizzas = [];
    this.orderTime = new Date();
  }

  /**
   * Add a pizza to the order
   */
  addPizza(pizza) {
    this.pizzas.push(pizza);
    return this;
  }

  /**
   * Get total order price
   */
  getTotalPrice() {
    return this.pizzas.reduce((total, pizza) => total + pizza.price, 0);
  }

  /**
   * Display order details
   */
  display() {
    console.log(`\n📦 ORDER #${this.orderId}`);
    console.log(`   Customer: ${this.customerName}`);
    console.log(`   Time: ${this.orderTime.toLocaleString()}`);
    console.log(`   Items: ${this.pizzas.length} pizza(s)`);
    console.log('   ─────────────────────────────');
    
    this.pizzas.forEach((pizza, index) => {
      console.log(`   ${index + 1}. ${pizza.getDescription()} - $${pizza.price.toFixed(2)}`);
    });
    
    console.log('   ─────────────────────────────');
    console.log(`   Total: $${this.getTotalPrice().toFixed(2)}`);
    console.log();
  }
}

// Simulate restaurant orders
console.log('Processing Customer Orders:');
console.log('-'.repeat(70));

const order1 = new Order('ORD001', 'John Doe');
order1.addPizza(new PizzaBuilder().applyPreset('Margherita').build());
order1.addPizza(new PizzaBuilder().applyPreset('Pepperoni').build());
order1.display();

const order2 = new Order('ORD002', 'Jane Smith');
order2.addPizza(new PizzaBuilder().applyPreset('Vegetarian').build());
order2.addPizza(new PizzaBuilder().applyPreset('Deluxe').build());
order2.addPizza(new PizzaBuilder().applyPreset('Spicy').build());
order2.display();

const order3 = new Order('ORD003', 'Bob Johnson');
order3.addPizza(new PizzaBuilder()
  .setCrust('Thick')
  .setSize('X-Large')
  .setSauce('BBQ')
  .addToppings(['Chicken', 'Bacon', 'Onions'])
  .addExtraCheese()
  .build());
order3.display();

/**
 * Summary Statistics
 */
console.log('📊 Restaurant Daily Summary:');
console.log('-'.repeat(70));

const allOrders = [order1, order2, order3];
const totalOrders = allOrders.length;
const totalPizzas = allOrders.reduce((sum, order) => sum + order.pizzas.length, 0);
const totalRevenue = allOrders.reduce((sum, order) => sum + order.getTotalPrice(), 0);
const averageOrderValue = totalRevenue / totalOrders;

console.log(`Total Orders: ${totalOrders}`);
console.log(`Total Pizzas Sold: ${totalPizzas}`);
console.log(`Total Revenue: $${totalRevenue.toFixed(2)}`);
console.log(`Average Order Value: $${averageOrderValue.toFixed(2)}`);
console.log(`Average Pizzas per Order: ${(totalPizzas / totalOrders).toFixed(1)}`);