package models

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"` //should not be visible on the frontend side and []byte because we want to encryt it, So it stays secure even if someone got access to database.
	Phone     string `json:"phone"`
}
