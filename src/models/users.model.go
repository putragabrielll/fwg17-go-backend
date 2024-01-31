package modelsUsers

import (
	"database/sql"
	"time"
	// "github.com/jmoiron/sqlx"
	"github.com/putragabrielll/go-backend/src/lib"
)


type Person struct {
    Id 				int 				`db:"id"`
	FullName 		sql.NullString 		`db:"fullName"`
	Email 			string 				`db:"email"`
	PhoneNumber 	sql.NullString 		`db:"phoneNumber"`
	Address 		sql.NullString 		`db:"address"`
	Picture 		sql.NullString 		`db:"picture"`
	Role 			string 		`db:"role"`
	Password 		string 				`db:"password"`
	CreatedAt 		time.Time 			`db:"createdAt"`
	UpdatedAt 		sql.NullTime 		`db:"updatedAt"`
}

// var DbConnect *sqlx.DB = lib.DBClient // V1. menggunakan koneksi dari lib.
// var DbConnect = lib.DbConnection() // V2. menggunakan koneksi dari lib.

func ListAllUsers() ([]Person, error) {
	// people := []Person{} // V1
    // db.Select(&people, "SELECT * FROM users") // V1

	sql := `SELECT * FROM users`
	data := []Person{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}