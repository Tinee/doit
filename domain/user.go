package domain

import (
	"time"

	"github.com/Tinee/doit/pkg/validation"

	"github.com/globalsign/mgo/bson"
	"golang.org/x/crypto/bcrypt"
)

// User represents the user domain object.
type User struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	Email     string        `json:"email"`
	FirstName string        `json:"firstName,omitempty"`
	LastName  string        `json:"lastName,omitempty"`
	Password  string        `json:"password,omitempty"`
	CreatedAt time.Time     `json:"createdAt"`
	UpdatedAt time.Time     `json:"updatedAt"`
}

// UserRepository interface for the user repo
type UserRepository interface {
	Create(User) (*User, error)
}

// EncryptPassword using bcrypt to generate a hashed password and asign to pointer.
func (c *User) EncryptPassword() {
	bs, _ := bcrypt.GenerateFromPassword([]byte(c.Password), bcrypt.DefaultCost)

	c.Password = string(bs)
}

// CompareHashedPasswordWith takes a string that it's comparing with the pointer User.
// if they don't match it throws back an error.
func (c *User) CompareHashedPasswordWith(compare string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(c.Password), []byte(compare)); err != nil {
		return ErrPasswordMissmatch
	}
	return nil
}

//ClearPassword removes the password, useful when we wan't to pass this object back to the caller.
func (c *User) ClearPassword() {
	c.Password = ""
}

// Validate tries to validate the struct, if it does not validate correctly it give back an error.
func (c *User) Validate() error {
	errs := validation.New()
	if c.Password == "" {
		errs.Add("Password", "Is required")
	}

	if c.Email == "" {
		errs.Add("Email", "Is required")
	}

	return errs
}
