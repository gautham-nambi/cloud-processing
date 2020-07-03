package io

import "github.com/google/uuid"

var JwtSecret = []byte("gautham kipso")

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Credentials struct {
	Id          uuid.UUID `json:"id"`
	AccessToken string    `json:"access_token"`
	Username    string    `json:"username"`
}
