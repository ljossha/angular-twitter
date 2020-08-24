package models

// User present of user model on DAL
type User struct {
	ID    int64  `json:"user_id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
