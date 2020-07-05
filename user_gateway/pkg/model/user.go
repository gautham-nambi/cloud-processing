package model

import (
	"parallelSystems/user_gateway/pkg/utils/hasher"

	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `db:"id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
}

func (u *User) CheckPassword() (isCorrectPassword bool, err error) {
	user, err := u.Get()
	if err != nil {
		return false, err
	}
	cipherUtil := hasher.Input{
		CipheredText: user.Password,
	}
	return cipherUtil.VerifyCipheredText(), nil
}
