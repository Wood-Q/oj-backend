package models

// SignUp struct to describe register a new user.
type SignUp struct {
	UserAccount string `json:"user_account" validate:"required,lte=255"`
	Password    string `json:"password" validate:"required,lte=255"`
	UserRole    string `json:"user_role" validate:"required,lte=25"`
	UserName    string `json:"user_name" validate:"required,lte=25"`
}

// SignIn struct to describe login user.
type SignIn struct {
	UserAccount string `json:"user_account" validate:"required,lte=255"`
	Password    string `json:"password" validate:"required,lte=255"`
}
