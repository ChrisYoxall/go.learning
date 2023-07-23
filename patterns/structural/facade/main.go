// A facade provides a simple interface to a complex subsystem which contains lots of moving parts. A facade might
// provide limited functionality in comparison to working with the subsystem directly. However, it includes only
// those features that clients really care about.

package main

// OrderSystem is a complex subsystem with lots of moving parts.
type OrderSystem struct{}

func (os *OrderSystem) ValidateOrder() {}

func (os *OrderSystem) ProcessPayment() {}

func (os *OrderSystem) UpdateInventory() {}

func (os *OrderSystem) ShipOrder() {}

// OrderSystemFacade simplifies placing an order by providing a simple interface to a complex subsystem.
type OrderSystemFacade struct {
	OrderSystem *OrderSystem
}

func (osf *OrderSystemFacade) PlaceOrder() {
	osf.OrderSystem.ValidateOrder()
	osf.OrderSystem.ProcessPayment()
	osf.OrderSystem.UpdateInventory()
	osf.OrderSystem.ShipOrder()
}

func main() {

	// Without facade:
	os := &OrderSystem{}
	os.ValidateOrder()
	os.ProcessPayment()
	os.UpdateInventory()
	os.ShipOrder()

	// With facade:
	osf := &OrderSystemFacade{}
	osf.PlaceOrder()
}
