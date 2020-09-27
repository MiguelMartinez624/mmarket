package auth

// RegisterUser object to be passed when want to register a user it contains
// basic to the register process
type RegisterUser struct {
	Password string      `json:"password"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Resource interface{} `json:"resource"`
}

//LoginAccount regular values for login
type LoginAccount struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
