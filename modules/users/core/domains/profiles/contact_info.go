package profiles

type CommunicationChannel int8

const (
	Email      CommunicationChannel = 0
	MobilPhone                      = 1
	LocalPhone                      = 2
)

type ContactInfo struct {
	ID          string               `json:"id" bson:"_id,omitempty"`
	Channel     CommunicationChannel `json:"channel" bson:"channel,omitempty"`
	ItsVerified bool                 `json:"its_verified" bson:"its_verified,omitempty"`
	Value       string               `json:"value" bson:"value,omitempty"`
	ItsMain     bool                 `json:"its_main" bson:"its_main,omitempty"`
}
