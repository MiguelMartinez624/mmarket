package products

type Product struct {
	ID          string  `json:"id" bson:"_id,omitempty"`
	StoreID     string  `json:"store_id" bson:"store_id,omitempty"`
	Name        string  `json:"name" bson:"Name,omitempty"`
	Description string  `json:"description" bson:"description,omitempty"`
	RawPrice    float32 `json:"raw_price" bson:"raw_price,omitempty"`
}
