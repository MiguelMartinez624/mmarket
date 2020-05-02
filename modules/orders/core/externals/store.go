package externals

type RequestItem struct {
	ProductID string
	Quantity  int
}

type AvailavilityStatus struct {
	ProductID    string
	Availability bool
}

type StoresModule interface {
	AskForAvailability(storeID string, car []RequestItem) (availability []AvailavilityStatus, err error)
}
