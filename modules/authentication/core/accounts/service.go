package accounts

import "context"

// Service contains the logic of this domain, accounts it use its the gate
// to validate and store, search delete accounts
type Service interface {
	//CreateAccount take care of all steps and validations to create a account, a username and email can only belong to 1
	//account and once that its taked it wont allow to create another account with it,
	//
	//* Step 1 : validate that the account data is valid this would be the email and the password as those are
	//the only required fields for a account creation, if a username its not provided it will set the email
	//as a default username that can be changed later.
	//
	//* Step 2 : check that the username or the email its not been used already taked by another account,
	//
	//* Step 3 : hash the password before stored as we never persist plain password, and set the first State
	//of the account to be UNVERIFIED
	//
	//* Step 4 : generate a validation hash that that can be used for validate the account after its creation this
	//can be send via email to the account email.
	//
	//* Step 5: the final step its to create a Unique ID to a resource that this account its guarding and return the
	//account keys (accountID, resourceID, validationHash).
	CreateAccount(ctx context.Context, acc Account) (keys *NewAccountKeys, err error)

	//Authenticate method to validate and account it follows the next step to performe this action.
	//1- check that the account actually exist.
	//2- validate the password against the hashed password stored.
	//3- checkout the state for the state of the account this may be on 3 different states.
	// STATUS
	//	BLOCKED : may be for multiple reasons, it doesnt matter to the authentication process with
	//	the account is blocked.
	//
	//	UNVERIFIED : the account email its not verified yet so the ownership of the email provided
	//	is still on a unknown state and the account cant be used.
	//
	//	ACTIVE : this is the ideal state of the account and the only one were the account can be used.
	Authenticate(ctx context.Context, username string, password string) (account *Account, err error)


	//ValidateAccountWithHash : validate an account that has been just created, each account have a unique
	//hash that its used to validate email ownership so we can be sure that the email belongs to the
	//account creator, will change the AccountStatus to be 	Active if everything goes well
	//
	//Returns the account without the password for security reasons or a error if something goes wrong
	//in the process
	ValidateAccountWithHash(ctx context.Context, hash string) (acc *Account, err error)
}

type DefaultService struct {
	accountRepository Repository
	encrypter         Encrypter
}

func NewService(accountRepository Repository, encrypter Encrypter) Service {
	return &DefaultService{
		accountRepository: accountRepository,
		encrypter:         encrypter,
	}
}
