package stores

// Address phisical direction of a store
type Address struct {
	ID         string `json:"id" bson:"_id ,omitempty"`
	Verified   bool   `json:"verified" bson:"verified ,omitempty"`
	Direcction string `json:"direction" bson:"direction ,omitempty"`
	ZipCode    string `json:"zip_code" bson:"zip_code ,omitempty"`
	Country    string `json:"country" bson:"country ,omitempty"`
}
