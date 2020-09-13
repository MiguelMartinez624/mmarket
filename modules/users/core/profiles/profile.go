package profiles

type Roles = string

var (

	// Client its a reference to th common user of the platform the one that will be
	// using the service for buying
	Client Roles = "client"

	// Delivery its the user that will handle P2P deliveries
	Delivery Roles = "delivery"

	// Provider its the one who put products to sell
	Provider Roles = "provider"
)

type Profile struct {
	ID        string        `json:"id" bson:"_id,omitempty"`
	AccountID string        `json:"account_id" bson:"account_id,omitempty"`
	FirstName string        `json:"firstname" bson:"firstname,omitempty"`
	LastName  string        `json:"lastname" bson:"lastname,omitempty"`
	Contacts  []ContactInfo `json:"contacts" bson:"contacts,omitempty"`
	Address   []Address     `json:"address" bson:"address,omitempty"`
}

func (p *Profile) IsValid() error {

	if p.AccountID == "" {
		return MissingAccountIDError{}
	}

	if p.FirstName == "" {
		return MissingFirstNameError{}
	}

	if len(p.Contacts) == 0 {
		return NoContactsOnProfileError{}
	}

	return nil
}
