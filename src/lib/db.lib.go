package lib

import (
	"os"
	"github.com/jmoiron/sqlx"
)



func DbConnection() *sqlx.DB{

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_CONNECT"))
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}

var DBClient *sqlx.DB = DbConnection() // untuk memanggil function koneksi yg sudah di buat