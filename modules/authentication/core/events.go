package auth

const (
	ACCOUNT_CREATED        string = "ACCOUNT_CREATED"
	ACCOUNT_EMAIL_VERIFIED string = "ACCOUNT_EMAIL_VERIFIED"
)

type AccountCreatedResult struct {
	Email string `json:"email"`
}
