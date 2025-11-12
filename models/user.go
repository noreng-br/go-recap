package models

// User entity model
type User struct {
	ID    string `json:"id"`
	Username  string `json:"username" validate="required"`
	Email string `json:"email" validate="required,email"`
	Password string `json:"password"`
  IsAdmin bool `json:"is_admin"`
}
