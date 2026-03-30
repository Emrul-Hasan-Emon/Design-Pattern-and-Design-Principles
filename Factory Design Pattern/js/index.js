/**
 * Factory Design Pattern - JavaScript Implementation
 * 
 * This example demonstrates the Factory Pattern with a shipping carrier system.
 * A factory creates different types of shipping carriers based on the type parameter,
 * without exposing the creation logic to the client.
 */

// ============================================================
// PRODUCT INTERFACE (Abstract Base Class)
// ============================================================

/**
 * ShippingCarrier Interface (Abstract Base Class)
 * Defines the contract that all shipping carriers must implement
 */
class ShippingCarrier {
  /**
   * Ship a package
   */
  ship(package) {
    throw new Error('ship() method must be implemented by concrete carrier');
  }

  /**
   * Calculate shipping cost
   */
  calculateCost(weight) {
    throw new Error('calculateCost() method must be implemented by concrete carrier');
  }

  /**
   * Get carrier name
   */
  getCarrierName() {
    throw new Error('getCarrierName() method must be implemented by concrete carrier');
  }

  /**
   * Get estimated delivery time
   */
  getEstimatedDelivery(destination) {
    throw new Error('getEstimatedDelivery() method must be implemented by concrete carrier');
  }
}

// ============================================================
// CONCRETE PRODUCTS - Shipping Carriers
// ============================================================

/**
 * FedEx Shipping Carrier
 */
class FedExCarrier extends ShippingCarrier {
  ship(package) {
    console.log(`📦 FedEx: Shipping package to ${package.destination}`);
    console.log(`   Tracking: FX${Math.random().toString().slice(2, 12).toUpperCase()}`);
    console.log(`   Status: ✅ Package picked up\n`);
    return {
      carrier: this.getCarrierName(),
      tracking: `FX${Math.random().toString().slice(2, 12).toUpperCase()}`,
      destination: package.destination,
      status: 'Picked up',
      timestamp: new Date().toISOString()
    };
  }

  calculateCost(weight) {
    // FedEx: $1.50 per kg + $5 base fee
    const baseFee = 5;
    const perKgCost = 1.5;
    return baseFee + (weight * perKgCost);
  }

  getCarrierName() {
    return 'FedEx';
  }

  getEstimatedDelivery(destination) {
    // FedEx: 2-3 days
    return '2-3 business days';
  }
}

/**
 * UPS Shipping Carrier
 */
class UPSCarrier extends ShippingCarrier {
  ship(package) {
    console.log(`📦 UPS: Shipping package to ${package.destination}`);
    console.log(`   Tracking: 1Z${Math.random().toString().slice(2, 12).toUpperCase()}`);
    console.log(`   Status: ✅ In transit\n`);
    return {
      carrier: this.getCarrierName(),
      tracking: `1Z${Math.random().toString().slice(2, 12).toUpperCase()}`,
      destination: package.destination,
      status: 'In transit',
      timestamp: new Date().toISOString()
    };
  }

  calculateCost(weight) {
    // UPS: $1.25 per kg + $4 base fee
    const baseFee = 4;
    const perKgCost = 1.25;
    return baseFee + (weight * perKgCost);
  }

  getCarrierName() {
    return 'UPS';
  }

  getEstimatedDelivery(destination) {
    // UPS: 1-2 days
    return '1-2 business days';
  }
}

/**
 * DHL Shipping Carrier
 */
class DHLCarrier extends ShippingCarrier {
  ship(package) {
    console.log(`📦 DHL: Shipping package to ${package.destination}`);
    console.log(`   Tracking: DHL${Math.random().toString().slice(2, 12).toUpperCase()}`);
    console.log(`   Status: ✅ Label created\n`);
    return {
      carrier: this.getCarrierName(),
      tracking: `DHL${Math.random().toString().slice(2, 12).toUpperCase()}`,
      destination: package.destination,
      status: 'Label created',
      timestamp: new Date().toISOString()
    };
  }

  calculateCost(weight) {
    // DHL: $1.75 per kg + $6 base fee
    const baseFee = 6;
    const perKgCost = 1.75;
    return baseFee + (weight * perKgCost);
  }

  getCarrierName() {
    return 'DHL';
  }

  getEstimatedDelivery(destination) {
    // DHL: 3-5 days (international)
    return '3-5 business days';
  }
}

/**
 * Express Overnight Carrier (Premium option)
 */
class ExpressOvernight extends ShippingCarrier {
  ship(package) {
    console.log(`🚀 Express Overnight: Shipping package to ${package.destination}`);
    console.log(`   Tracking: EXP${Math.random().toString().slice(2, 12).toUpperCase()}`);
    console.log(`   Status: ✅ Priority handling\n`);
    return {
      carrier: this.getCarrierName(),
      tracking: `EXP${Math.random().toString().slice(2, 12).toUpperCase()}`,
      destination: package.destination,
      status: 'Priority handling',
      timestamp: new Date().toISOString()
    };
  }

  calculateCost(weight) {
    // Express: $5 per kg + $15 base fee + 50% surcharge
    const baseFee = 15;
    const perKgCost = 5;
    const totalCost = baseFee + (weight * perKgCost);
    return totalCost * 1.5; // 50% premium for overnight
  }

  getCarrierName() {
    return 'Express Overnight';
  }

  getEstimatedDelivery(destination) {
    // Express: Next day
    return 'Next business day (guaranteed)';
  }
}

// ============================================================
// FACTORY - Object Creator
// ============================================================

/**
 * ShippingCarrierFactory
 * Encapsulates the creation logic for all shipping carriers
 */
class ShippingCarrierFactory {
  /**
   * Create a shipping carrier based on type
   * This is the factory method
   */
  static createCarrier(carrierType) {
    const type = carrierType.toLowerCase().trim();

    switch (type) {
      case 'fedex':
      case 'fedx':
        return new FedExCarrier();

      case 'ups':
        return new UPSCarrier();

      case 'dhl':
        return new DHLCarrier();

      case 'express':
      case 'overnight':
      case 'express-overnight':
        return new ExpressOvernight();

      default:
        throw new Error(
          `Unknown carrier: "${carrierType}". Available: fedex, ups, dhl, express`
        );
    }
  }

  /**
   * Get list of available carriers
   */
  static getAvailableCarriers() {
    return ['fedex', 'ups', 'dhl', 'express'];
  }

  /**
   * Create carrier with automatic price comparison
   * Returns the cheapest option
   */
  static findCheapestCarrier(weight) {
    const carriers = this.getAvailableCarriers();
    let cheapest = {
      carrier: null,
      cost: Infinity
    };

    for (const carrierType of carriers) {
      const carrier = this.createCarrier(carrierType);
      const cost = carrier.calculateCost(weight);

      if (cost < cheapest.cost) {
        cheapest = {
          carrier: carrier,
          cost: cost
        };
      }
    }

    return cheapest;
  }
}

// ============================================================
// CLIENT CODE - Usage Examples
// ============================================================

console.log('FACTORY DESIGN PATTERN - Shipping Carrier System');
console.log('\n');

/**
 * Example 1: Basic Factory Usage
 */
console.log('📋 EXAMPLE 1: Create Different Carriers Using Factory');

const package1 = {
  id: 'PKG001',
  destination: 'New York, NY',
  weight: 2.5
};

// Create FedEx carrier using factory
const fedexCarrier = ShippingCarrierFactory.createCarrier('fedex');
console.log(`Carrier: ${fedexCarrier.getCarrierName()}`);
console.log(`Cost: $${fedexCarrier.calculateCost(package1.weight).toFixed(2)}`);
console.log(`Delivery: ${fedexCarrier.getEstimatedDelivery(package1.destination)}`);
fedexCarrier.ship(package1);

/**
 * Example 2: Switch Carriers Without Changing Client Code
 */
console.log('📋 EXAMPLE 2: Switch to Different Carrier');

const upsCarrier = ShippingCarrierFactory.createCarrier('ups');
console.log(`Carrier: ${upsCarrier.getCarrierName()}`);
console.log(`Cost: $${upsCarrier.calculateCost(package1.weight).toFixed(2)}`);
console.log(`Delivery: ${upsCarrier.getEstimatedDelivery(package1.destination)}`);
upsCarrier.ship(package1);

/**
 * Example 3: Process Multiple Packages with Different Carriers
 */
console.log('📋 EXAMPLE 3: Multiple Packages with Different Carriers');

const packages = [
  { id: 'PKG002', destination: 'Los Angeles, CA', weight: 1.2, carrier: 'dhl' },
  { id: 'PKG003', destination: 'Chicago, IL', weight: 5.0, carrier: 'fedex' },
  { id: 'PKG004', destination: 'Miami, FL', weight: 0.8, carrier: 'ups' }
];

for (const pkg of packages) {
  const carrier = ShippingCarrierFactory.createCarrier(pkg.carrier);
  const cost = carrier.calculateCost(pkg.weight);
  console.log(
    `📦 ${pkg.id}: ${carrier.getCarrierName()} - $${cost.toFixed(2)} - ${pkg.destination}`
  );
  carrier.ship(pkg);
}

/**
 * Example 4: Available Carriers List
 */
console.log('📋 EXAMPLE 4: Available Carriers');
const availableCarriers = ShippingCarrierFactory.getAvailableCarriers();
console.log('Available carriers:');
availableCarriers.forEach(carrierType => {
  const carrier = ShippingCarrierFactory.createCarrier(carrierType);
  console.log(
    `  ✓ ${carrier.getCarrierName()}: $${carrier.calculateCost(1).toFixed(2)}/kg (base)`
  );
});
console.log('\n');

/**
 * Example 5: Price Comparison - Find Cheapest Carrier
 */
console.log('📋 EXAMPLE 5: Price Comparison - Find Cheapest Carrier');

const comparisonWeights = [1, 5, 10, 25];

for (const weight of comparisonWeights) {
  const cheapest = ShippingCarrierFactory.findCheapestCarrier(weight);
  console.log(
    `For ${weight}kg: ${cheapest.carrier.getCarrierName()} is cheapest at $${cheapest.cost.toFixed(2)}`
  );
}
console.log('\n');

// ============================================================
// ADVANCED EXAMPLE: Order Shipment Manager
// ============================================================

console.log('\n' + '='.repeat(70));
console.log('ADVANCED EXAMPLE: Order Shipment Manager');
console.log('='.repeat(70));
console.log('\n');

/**
 * ShipmentManager using the factory
 */
class ShipmentManager {
  constructor() {
    this.shipments = [];
  }

  /**
   * Create shipment with specified carrier
   */
  createShipment(order, carrierType) {
    console.log(`\n📦 Processing Order: ${order.id}`);
    console.log(`   Items: ${order.items}`);
    console.log(`   Weight: ${order.weight}kg`);
    console.log(`   Destination: ${order.destination}`);

    // Use factory to create carrier
    const carrier = ShippingCarrierFactory.createCarrier(carrierType);

    console.log(`   Carrier: ${carrier.getCarrierName()}`);
    const shippingCost = carrier.calculateCost(order.weight);
    console.log(`   Shipping Cost: $${shippingCost.toFixed(2)}`);
    console.log(`   Delivery: ${carrier.getEstimatedDelivery(order.destination)}`);

    // Ship the package
    const result = carrier.ship(order);

    // Record shipment
    this.shipments.push({
      ...result,
      orderId: order.id,
      cost: shippingCost
    });

    return result;
  }

  /**
   * Auto-select cheapest carrier
   */
  createShipmentCheapest(order) {
    console.log(`\n📦 Processing Order: ${order.id} (Auto-select cheapest)`);
    console.log(`   Items: ${order.items}`);
    console.log(`   Weight: ${order.weight}kg`);

    const cheapest = ShippingCarrierFactory.findCheapestCarrier(order.weight);
    return this.createShipment(order, cheapest.carrier.getCarrierName().toLowerCase());
  }

  /**
   * Get shipment history
   */
  getShipmentHistory() {
    console.log('\n📊 Shipment History:');
    this.shipments.forEach((shipment, index) => {
      console.log(
        `   ${index + 1}. Order ${shipment.orderId} - ${shipment.carrier} - Tracking: ${shipment.tracking} - Cost: $${shipment.cost.toFixed(2)}`
      );
    });
  }

  /**
   * Calculate total shipping cost
   */
  getTotalShippingCost() {
    return this.shipments.reduce((total, shipment) => total + shipment.cost, 0);
  }
}

// Simulate order processing
const manager = new ShipmentManager();

const orders = [
  { id: 'ORD001', items: '3 books', weight: 2.5, destination: 'New York, NY' },
  { id: 'ORD002', items: '1 laptop', weight: 3.2, destination: 'San Francisco, CA' },
  { id: 'ORD003', items: '5 shoes', weight: 5.5, destination: 'Miami, FL' },
  { id: 'ORD004', items: '2 watches', weight: 0.5, destination: 'Boston, MA' }
];

console.log('Processing Orders:');
console.log('-'.repeat(70));

// Process some orders with specific carriers
manager.createShipment(orders[0], 'fedex');
manager.createShipment(orders[1], 'ups');

// Process remaining orders with cheapest option
manager.createShipmentCheapest(orders[2]);
manager.createShipmentCheapest(orders[3]);

// Show history and costs
manager.getShipmentHistory();
console.log(`\nTotal Shipping Cost: $${manager.getTotalShippingCost().toFixed(2)}`);