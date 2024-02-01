package modelsUsers

import (
	"database/sql"
	"time"
	// "github.com/jmoiron/sqlx"
	"github.com/putragabrielll/go-backend/src/lib"
)


type Person struct {
    Id 				int 				`db:"id" json:"id"`
	FullName 		sql.NullString 		`db:"fullName" json:"fullname"`
	Email 			string 				`db:"email" json:"email" form:"email"`
	PhoneNumber 	sql.NullString 		`db:"phoneNumber" json:"phoneNumber"`
	Address 		sql.NullString 		`db:"address" json:"address"`
	Picture 		sql.NullString 		`db:"picture" json:"picture"`
	Role 			string 				`db:"role" json:"role" form:"role"`
	Password 		string 				`db:"password" json:"password" form:"password"`
	CreatedAt 		time.Time 			`db:"createdAt" json:"createdAt"`
	UpdatedAt 		sql.NullTime 		`db:"updatedAt" json:"updatedAt"`
}

// var DbConnect *sqlx.DB = lib.DBClient // V1. menggunakan koneksi dari lib.
// var DbConnect = lib.DbConnection() // V2. menggunakan koneksi dari lib.

func ListAllUsers() ([]Person, error) {
	// people := []Person{} // V1
    // db.Select(&people, "SELECT * FROM users") // V1

	sql := `SELECT * FROM "users"`
	data := []Person{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}

func FindUsersId(id int) (Person, error){
	sql := `SELECT * FROM "users" WHERE "id"=$1`
	data := Person{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

func CreateUsers(data Person) (Person, error){
	sql := `
	INSERT INTO "users"
    ("email", "role", "password")
    VALUES
    (:email, :role, :password)
    RETURNING *
    `
	returning := Person{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}