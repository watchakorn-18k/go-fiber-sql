package entities

type NewUserBody struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type UserDataFormat struct {
	UserID   string `json:"user_id,omitempty"`
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
}
