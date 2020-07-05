package model

import (
	"github.com/google/uuid"
	"parallelSystems/user_gateway/cmd/db"
	"parallelSystems/user_gateway/pkg/utils/hasher"
)

type UserRepo interface {
	Save() (user *User, err error)
	Get(username string) (user *User, err error)
}

func (u *User) Save() (user *User, err error) {
	u.Id = uuid.New()
	tx := db.GetConnection().MustBegin()
	password := hasher.Input{
		Text:         u.Password,
		CipheredText: "",
	}
	password.Encrypt()
	u.Password = password.CipheredText
	_, err = tx.NamedExec(
		"INSERT INTO PUBLIC.user (id, username, password) VALUES (:id, :username, :password)", &u)
	if err != nil {
		return user, err
	}
	err = tx.Commit()
	if err != nil {
		return user, err
	}
	user = &User{
		Id:       u.Id,
		Username: u.Username,
		Password: u.Password,
	}
	return user, nil
}

func (u *User) Get() (user *User, err error) {
	db := db.GetConnection()
	err = db.Get(&user, "select * from public.user where username=$1", u.Username)
	if err != nil {
		return user, err
	}
	return user, nil
}
