package profiles

//Address phisical direction
type Address struct {
	ID          string `json:"id" bson:"_id,omitempty"`
	City        string `json:"city" bson:"city"`
	Country     string `json:"country" bson:"country"`
	Street      string `json:"street" bson:"street"`
	Description string `json:"description" bson:"description"`
}
