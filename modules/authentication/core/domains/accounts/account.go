package accounts

//  AccountStatus are the actuall status that a account can have
type AccountStatus = int8

var (
	// An account can be blocked for many reason but while this its the
	//current state the access its temporally forbidden
	Blocked AccountStatus = 0

	// Active have full access to the account benefits
	Active AccountStatus = 1

	// Unverified the account have been created but it havent be verified by
	// its owner, it can't be used until the owner of the account verify its
	// owershift by any of the protocol allowed for account verification.
	Unverified AccountStatus = 2
)

//Credetial protect a resource
type Account struct {
	ID             string        `json:"id" bson:"_id,omitempty"`
	Username       string        `json:"username" bson:"username,omitempty"`
	Password       string        `json:"password" bson:"password,omitempty"`
	Status         AccountStatus `json:"status" bson:"status"`
	ValidationHash string        `json:"validation_hash" bson:"validation_hash"`
}

type NewAccountKeys struct {
	AccountID        string
	VerificationHash string
}
