package orders

type ItemsList struct {
	ItemID       string
	PricePerItem float64
	Quantity     int
}

type OrderDetails struct {
	itemsID []ItemsList
}
