package db

import "github.com/jmoiron/sqlx"

var DB * sqlx.DB

func GetConnection() *sqlx.DB {
	return DB
}

func SetConnection(db *sqlx.DB) {
	DB = db
}
