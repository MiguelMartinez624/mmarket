package orders

type Validator struct{}

func (v Validator) Validate(order *Order) error {
	if order.Details.Items == nil || len(order.Details.Items) == 0 {
		return ErrNoItemsOnOrder{}
	}

	if order.ShippingInfo.ItsDelivery == true && order.ShippingInfo.AddressID == "" {
		return ErrMissingShippingInfo{}
	}

	if order.StoreID == "" {
		return ErrMissingStore{}
	}

	if order.CustomerID == "" {
		return ErrMissingCustomerID{}
	}

	return nil
}
