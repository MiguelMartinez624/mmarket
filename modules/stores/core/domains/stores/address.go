package stores

// Address phisical direction of a store
type Address struct {
	Verified  bool   `json:"verified" bson:"verified,omitempty"`
	Direction string `json:"direction" bson:"direction,omitempty"`
	ZipCode   string `json:"zip_code" bson:"zip_code,omitempty"`
	Country   string `json:"country" bson:"country,omitempty"`
}
