package modelsUsers

import (
	// "github.com/jmoiron/sqlx"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * users
func ListAllUsers() ([]services.Person, error) {
	// people := []Person{} // V1
    // db.Select(&people, "SELECT * FROM users") // V1

	sql := `SELECT * FROM "users"`
	data := []services.Person{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}

// SELECT users BY id
func FindUsersId(id int) (services.Person, error){
	sql := `SELECT * FROM "users" WHERE "id"=$1`
	data := services.Person{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// CREATE users
func CreateUsers(data services.Person) (services.Person, error){
	sql := `
	INSERT INTO "users"
    ("email", "role", "password")
    VALUES
    (:email, :role, :password)
    RETURNING *
    `
	returning := services.Person{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE users
func UpdateUsers(data services.Person) (services.Person, error){ // bisa teruddate jika type untuk fullName di ganti jadi string.
	sql := `
	UPDATE "users" SET 
	"fullName"=:fullName, 
	"phoneNumber"=:phoneNumber, 
	"address"=:address, 
	"password"=COALESCE(NULLIF(:password,''),password), 
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	// "fullName"=COALESCE(NULLIF(:fullName,''),"fullName"),
	// "phoneNumber"=COALESCE(NULLIF(:phoneNumber,''),"phoneNumber"),
	// "address"=COALESCE(NULLIF(:address,''),"address"),
	// "password"=COALESCE(NULLIF(:password,''),"password"),
	returning := services.Person{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE users BY id
func DeleteUsersId(id int) (services.Person, error){
	sql := `DELETE FROM "users" WHERE "id"= $1 RETURNING *`
	data := services.Person{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}