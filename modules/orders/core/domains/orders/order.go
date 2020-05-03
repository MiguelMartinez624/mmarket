package orders

type OrderStatus = string

var (

	//Pending status
	Pending OrderStatus = "pending"

	// Canceled and order was canceled
	Canceled OrderStatus = "canceled"

	// Completed and order was completed by all parts of the deal
	Completed OrderStatus = "completed"
)

type Order struct {
	ID           string
	StoreID      string
	CustomerID   string
	Number       string
	Status       OrderStatus
	Details      *OrderDetails
	ShippingInfo *ShippingInfo
}

type ShippingInfo struct {
	ItsDelivery bool
	AddressID   string
}
