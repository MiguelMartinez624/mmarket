package orders

import "fmt"

type ErrNoItemsOnOrder struct{}

func (e ErrNoItemsOnOrder) Error() string {
	return fmt.Sprint("0 items on order, a least 1 item its required.")
}

type ErrMissingShippingInfo struct{}

func (e ErrMissingShippingInfo) Error() string {
	return fmt.Sprint("Orders with delivery,are required to have a shipping address.")
}

type ErrMissingStore struct{}

func (e ErrMissingStore) Error() string {
	return fmt.Sprint("missing storeID,orders require a store origin.")
}

type ErrMissingCustomerID struct{}

func (e ErrMissingCustomerID) Error() string {
	return fmt.Sprint("missing customerID,orders require a customer.")
}
