package models

// Entity basic property of a model
type EntityData struct {
	ID        string `json:"id" bson:"_id,omitempty"`
	CreatedAt int64  `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt int64  `json:"updated_at" bson:"updated_at,omitempty"`
}

type Entity interface {
	ItsEntity() bool
}
