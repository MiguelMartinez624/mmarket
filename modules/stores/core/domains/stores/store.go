package stores

type Store struct {
	ID        string   `json:"id" bson:"_id,omitempty"`
	ProfileID string   `json:"profile_id" bson:"profile_id,omitempty"`
	Name      string   `json:"name" bson:"name,omitempty"`
	Address   *Address `json:"address" bson:"address,omitempty"`
}
