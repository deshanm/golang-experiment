package models

import (
	"fmt"
	"time"

	"github.com/revel/revel"
)

// User struct
type User struct {
	userID        int
	email         string
	name          string
	password      string
	emailVerified string
	createdAt     time.Time
	updatedAt     time.Time
}

/**
Interface to validate User
*/
func (user User) Validate(v *revel.Validation) {
	v.Required(user.email)
	v.Required(user.password)
}

/*
Insert new User
*/
func (b *User) InsertUser() error {
	fmt.Printf("Get post %#v\n", b)
	return nil
}
