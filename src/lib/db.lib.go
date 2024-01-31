package lib

import (
	"log"
	"os"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)



func DbConnection() *sqlx.DB{
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_CONNECT"))
	if err != nil {
		// panic(err.Error())
		log.Fatalln(err)
	}

	err = db.Ping()
	if err != nil {
		// panic(err.Error())
		log.Fatalln(err)
	}

	return db
}

var DBClient *sqlx.DB = DbConnection() // untuk memanggil function koneksi yg sudah di buat