package common

type UserResponse struct {
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
}

type RegisterRequest struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
