/**
 * Abstract PaymentStrategy Interface
 * Defines the contract that all payment strategies must follow
 * In JavaScript, we use a base class to simulate an interface
 */
class PaymentStrategy {
  /**
   * Execute the payment strategy
   */
  pay(amount) {
    throw new Error('pay() method must be implemented by concrete strategy');
  }

  /**
   * Validate payment credentials
   */
  validate() {
    throw new Error('validate() method must be implemented by concrete strategy');
  }

  /**
   * Get strategy name
   */
  getStrategyName() {
    throw new Error('getStrategyName() method must be implemented by concrete strategy');
  }
}

// ============================================================
// CONCRETE STRATEGIES - Payment Methods
// ============================================================

/**
 * Credit Card Payment Strategy
 */
class CreditCardPayment extends PaymentStrategy {
  constructor(cardNumber, cardHolder, expiryDate, cvv) {
    super();
    this.cardNumber = cardNumber;
    this.cardHolder = cardHolder;
    this.expiryDate = expiryDate;
    this.cvv = cvv;
  }

  pay(amount) {
    console.log(`🏦 Processing Credit Card Payment`);
    console.log(`   Card: ${this.maskCardNumber(this.cardNumber)}`);
    console.log(`   Holder: ${this.cardHolder}`);
    console.log(`   Amount: $${amount.toFixed(2)}`);
    console.log(`   Status: ✅ Payment successful!\n`);
    return {
      success: true,
      method: 'Credit Card',
      amount: amount,
      timestamp: new Date().toISOString()
    };
  }

  maskCardNumber(cardNumber) {
    const last4 = cardNumber.slice(-4);
    return `****-****-****-${last4}`;
  }

  validate() {
    // Check if card number is valid (basic check)
    return this.cardNumber && this.cardNumber.length >= 13 && this.cvv && this.cvv.length >= 3;
  }

  getStrategyName() {
    return 'Credit Card';
  }
}

/**
 * PayPal Payment Strategy
 */
class PayPalPayment extends PaymentStrategy {
  constructor(email) {
    super();
    this.email = email;
  }

  pay(amount) {
    console.log(`📱 Processing PayPal Payment`);
    console.log(`   Email: ${this.email}`);
    console.log(`   Amount: $${amount.toFixed(2)}`);
    console.log(`   Authenticating with PayPal...`);
    console.log(`   Status: ✅ Payment successful!\n`);
    return {
      success: true,
      method: 'PayPal',
      amount: amount,
      timestamp: new Date().toISOString()
    };
  }

  validate() {
    // Check if email is valid
    return this.email && /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(this.email);
  }

  getStrategyName() {
    return 'PayPal';
  }
}

/**
 * Bitcoin Payment Strategy
 */
class BitcoinPayment extends PaymentStrategy {
  constructor(walletAddress) {
    super();
    this.walletAddress = walletAddress;
  }

  pay(amount) {
    console.log(`🪙 Processing Bitcoin Payment`);
    console.log(`   Wallet: ${this.maskWallet(this.walletAddress)}`);
    console.log(`   Amount: ${(amount * 0.000025).toFixed(6)} BTC`);
    console.log(`   Confirming transaction on blockchain...`);
    console.log(`   Status: ✅ Payment successful!\n`);
    return {
      success: true,
      method: 'Bitcoin',
      amount: amount,
      btcAmount: (amount * 0.000025).toFixed(6),
      timestamp: new Date().toISOString()
    };
  }

  maskWallet(address) {
    const first6 = address.slice(0, 6);
    const last4 = address.slice(-4);
    return `${first6}...${last4}`;
  }

  validate() {
    // Check if wallet address is valid (basic check)
    return this.walletAddress && this.walletAddress.length > 25;
  }

  getStrategyName() {
    return 'Bitcoin';
  }
}

/**
 * Apple Pay Strategy
 */
class ApplePayPayment extends PaymentStrategy {
  constructor(applerId) {
    super();
    this.applerId = applerId;
  }

  pay(amount) {
    console.log(`🍎 Processing Apple Pay Payment`);
    console.log(`   Apple ID: ${this.applerId}`);
    console.log(`   Amount: $${amount.toFixed(2)}`);
    console.log(`   Using biometric authentication...`);
    console.log(`   Status: ✅ Payment successful!\n`);
    return {
      success: true,
      method: 'Apple Pay',
      amount: amount,
      timestamp: new Date().toISOString()
    };
  }

  validate() {
    // Check if Apple ID is valid
    return this.applerId && this.applerId.includes('@');
  }

  getStrategyName() {
    return 'Apple Pay';
  }
}

// ============================================================
// CONTEXT - Payment Processor
// ============================================================

/**
 * Payment Processor (Context Class)
 * Delegates payment processing to different strategies based on selection
 */
class PaymentProcessor {
  constructor(paymentStrategy) {
    this.validateStrategy(paymentStrategy);
    this.strategy = paymentStrategy;
  }

  /**
   * Validate that strategy implements PaymentStrategy interface
   */
  validateStrategy(strategy) {
    if (!(strategy instanceof PaymentStrategy)) {
      throw new Error('Strategy must be an instance of PaymentStrategy');
    }
    if (typeof strategy.pay !== 'function') {
      throw new Error('Strategy must implement pay() method');
    }
    if (typeof strategy.validate !== 'function') {
      throw new Error('Strategy must implement validate() method');
    }
    if (typeof strategy.getStrategyName !== 'function') {
      throw new Error('Strategy must implement getStrategyName() method');
    }
  }

  /**
   * Set/Change payment strategy at runtime
   */
  setPaymentStrategy(paymentStrategy) {
    this.validateStrategy(paymentStrategy);
    console.log(`💱 Switching to ${paymentStrategy.getStrategyName()}...\n`);
    this.strategy = paymentStrategy;
  }

  /**
   * Process payment using the current strategy
   */
  processPayment(amount) {
    if (!this.strategy) {
      console.error('❌ No payment strategy selected!');
      return null;
    }
    return this.strategy.pay(amount);
  }

  /**
   * Get current payment method info
   */
  getCurrentStrategy() {
    return this.strategy.constructor.name;
  }
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

console.log('='.repeat(60));
console.log('STRATEGY DESIGN PATTERN - PAYMENT SYSTEM EXAMPLE');
console.log('='.repeat(60));
console.log('\n');

/**
 * Example 1: Payment with Credit Card
 */
console.log('📋 EXAMPLE 1: Credit Card Payment');
console.log('-'.repeat(60));
const creditCardStrategy = new CreditCardPayment(
  '4532-1234-5678-9010',
  'John Doe',
  '12/25',
  '123'
);
const processor = new PaymentProcessor(creditCardStrategy);
processor.processPayment(150.00);

/**
 * Example 2: Switch to PayPal Strategy
 */
console.log('📋 EXAMPLE 2: Switch to PayPal');
console.log('-'.repeat(60));
processor.setPaymentStrategy(
  new PayPalPayment('john.doe@example.com')
);
processor.processPayment(75.50);

/**
 * Example 3: Switch to Bitcoin Strategy
 */
console.log('📋 EXAMPLE 3: Switch to Bitcoin');
console.log('-'.repeat(60));
processor.setPaymentStrategy(
  new BitcoinPayment('1A1z7agoat5mLrQH5r8RN85dy1eWkTqeUP')
);
processor.processPayment(200.00);

/**
 * Example 4: Switch to Apple Pay Strategy
 */
console.log('📋 EXAMPLE 4: Switch to Apple Pay');
console.log('-'.repeat(60));
processor.setPaymentStrategy(
  new ApplePayPayment('john.doe@icloud.com')
);
processor.processPayment(99.99);

// ============================================================
// ADVANCED EXAMPLE: Using Strategies in a Shopping Cart
// ============================================================

console.log('\n' + '='.repeat(60));
console.log('ADVANCED EXAMPLE: E-Commerce Shopping Cart');
console.log('='.repeat(60));
console.log('\n');

/**
 * Shopping Cart with flexible payment options
 */
class ShoppingCart {
  constructor() {
    this.items = [];
    this.paymentProcessor = null;
  }

  addItem(name, price, quantity = 1) {
    this.items.push({ name, price, quantity });
    console.log(`✅ Added ${quantity}x ${name} ($${price}) to cart`);
  }

  getTotal() {
    return this.items.reduce((total, item) => 
      total + (item.price * item.quantity), 0
    );
  }

  showCart() {
    console.log('\n🛒 CART CONTENTS:');
    this.items.forEach(item => {
      console.log(`   - ${item.name}: $${item.price} x ${item.quantity} = $${(item.price * item.quantity).toFixed(2)}`);
    });
    console.log(`   ├─ Total: $${this.getTotal().toFixed(2)}\n`);
  }

  checkout(paymentStrategy) {
    console.log(`📦 Proceeding to checkout with ${paymentStrategy.constructor.name}...`);
    this.paymentProcessor = new PaymentProcessor(paymentStrategy);
    const total = this.getTotal();
    console.log(`💰 Processing payment of $${total.toFixed(2)}:\n`);
    return this.paymentProcessor.processPayment(total);
  }
}

// Simulate shopping scenario
console.log('👤 Customer Shopping...\n');
const cart = new ShoppingCart();
cart.addItem('Laptop', 999.99, 1);
cart.addItem('Mouse', 29.99, 2);
cart.addItem('Keyboard', 79.99, 1);

cart.showCart();

// Checkout with Credit Card
console.log('📋 CHECKOUT METHOD 1: Credit Card');
console.log('-'.repeat(60));
cart.checkout(new CreditCardPayment('4532-1234-5678-9010', 'Jane Smith', '06/26', '456'));

// Same customer wants to pay with PayPal instead
console.log('📋 CHECKOUT METHOD 2: Changing to PayPal');
console.log('-'.repeat(60));
const cartNew = new ShoppingCart();
cartNew.addItem('USB-C Cable', 15.99, 3);
cartNew.addItem('Monitor Stand', 45.00, 1);
cartNew.showCart();
cartNew.checkout(new PayPalPayment('jane.smith@example.com'));

// ============================================================
// KEY BENEFITS DEMONSTRATED
// ============================================================

console.log('='.repeat(60));
console.log('KEY BENEFITS OF STRATEGY PATTERN');
console.log('='.repeat(60));
console.log(`
1. ✅ FLEXIBILITY: Change payment method at runtime
2. ✅ MAINTAINABILITY: Each strategy is independent
3. ✅ EXTENSIBILITY: Easy to add new payment methods
4. ✅ TESTABILITY: Test each strategy separately
5. ✅ NO CONDITIONALS: No massive if-else chains
6. ✅ SINGLE RESPONSIBILITY: Each class does one thing
7. ✅ OPEN/CLOSED PRINCIPLE: Open for extension, closed for modification
`);

// ============================================================
// ADDING A NEW STRATEGY (Demonstrates OCP - Open/Closed Principle)
// ============================================================

/**
 * New Strategy: Google Pay (Can be added without modifying existing code!)
 */
class GooglePayPayment extends PaymentStrategy {
  constructor(email) {
    super();
    this.email = email;
  }

  pay(amount) {
    console.log(`🔵 Processing Google Pay Payment`);
    console.log(`   Email: ${this.email}`);
    console.log(`   Amount: $${amount.toFixed(2)}`);
    console.log(`   Authenticating with Google Account...`);
    console.log(`   Status: ✅ Payment successful!\n`);
    return {
      success: true,
      method: 'Google Pay',
      amount: amount,
      timestamp: new Date().toISOString()
    };
  }

  validate() {
    // Check if Google email is valid
    return this.email && this.email.includes('@');
  }

  getStrategyName() {
    return 'Google Pay';
  }
}

console.log('📋 EXAMPLE 5: Adding Google Pay (New Strategy)');
console.log('-'.repeat(60));
console.log('Notice: We added a new payment method WITHOUT modifying existing code!\n');
const googlePayProcessor = new PaymentProcessor(
  new GooglePayPayment('user@gmail.com')
);
googlePayProcessor.processPayment(125.00);

console.log('='.repeat(60));
console.log('✨ Strategy Pattern Implementation Complete!');
console.log('='.repeat(60));
