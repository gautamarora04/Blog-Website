package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  []byte `json:"-"` //should not be visible on the frontend side and []byte because we want to encryt it, So it stays secure even if someone got access to database.
	Phone     string `json:"phone"`
}

// to encrypt password before saving into database usihg golang bcrypt this is basically hashing, not encryting.
// For revison purpose, https://stackoverflow.com/questions/18084595/how-to-decrypt-hash-stored-by-bcrypt
func (user *User) HashPassword(password string) {
	hashedpassword, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	user.Password = hashedpassword
}

// Now for the login we need to check weather the password is the same as the passord in the database

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword(user.Password, []byte(password))
	// user.pass is the hashed password stored in the database and password string is provided by the user during the time of login
	return err
}
