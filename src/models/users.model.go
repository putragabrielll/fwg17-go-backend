package modelsUsers

import "github.com/jmoiron/sqlx"

func UsersModel(db *sqlx.DB){
	tx := db.MustBegin()
	tx.Select(`SELECT * FROM "users"`)
}