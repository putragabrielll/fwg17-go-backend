package models

import (
	// "github.com/jmoiron/sqlx"
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * users
func ListAllUsers(filter string, sortby string, order string, limit int, offset int) ([]services.PersonNet, error) {
	// people := []Person{} // V1
    // db.Select(&people, "SELECT * FROM users") // V1

	sql := `
	SELECT * 
	FROM "users"
	WHERE "fullName" ILIKE $1
	ORDER BY "`+sortby+`" `+order+`
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.PersonNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountAllUsers(filter string) (int, error){
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "users"
	WHERE "fullName" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT users BY id
func FindUsersId(id int) (services.PersonNet, error){
	sql := `SELECT * FROM "users" WHERE "id"=$1`
	data := services.PersonNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// CREATE users ADMIN
func CreateUsers(data services.Person) (services.Person, error){
	sql := `
	INSERT INTO "users"
    ("fullName", "email", "phoneNumber", "address", "picture", "role", "password")
    VALUES
    (:fullName, :email, :phoneNumber, :address, :picture, :role, :password)
    RETURNING *
    `
	returning := services.Person{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE users
func UpdateUsers(data services.Person) (services.Person, error){ // bisa teruddate jika type untuk fullName di ganti jadi string.
	sql := `
	UPDATE "users" SET 
	"fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	"phoneNumber"=COALESCE(NULLIF(:phoneNumber,''),"phoneNumber"),
	"address"=COALESCE(NULLIF(:address,''),"address"),
	"picture"=COALESCE(NULLIF(:picture,''),"picture"),
	"password"=COALESCE(NULLIF(:password,''),"password"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Person{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE users BY id
func DeleteUsersId(id int) (services.PersonNet, error){
	sql := `DELETE FROM "users" WHERE "id"= $1 RETURNING *`
	data := services.PersonNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}




//------------ AUTH ------------
// LOGIN users BY email
func FindUsersByEmail(email string) (services.PersonNet, error){
	sql := `SELECT * FROM "users" WHERE "email"=$1`
	data := services.PersonNet{}
	err := lib.DbConnection().Get(&data, sql, email) // id diambil dari parameter id.
	return data, err
}

// REGISTER users
func RegisterUsers(data services.RLUsers) (services.PersonNet, error){
	sql := `
	INSERT INTO "users"
    ("email", "role", "password")
    VALUES
    (:email, :role, :password)
    RETURNING *
    `
	returning := services.PersonNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}