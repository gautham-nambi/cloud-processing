package model

import (
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}

func (u *User) CreateUser() (user *User, err error) {
	
}

func (u *User) encryptPassword() {
	//TODO Encrypt Password
}

func (u *User) VerifyPassword() bool {
	//TODO Verify encrypted password
	return false
}
