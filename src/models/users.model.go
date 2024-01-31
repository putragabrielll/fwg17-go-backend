package modelsUsers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/putragabrielll/go-backend/src/lib"
)


// type Person struct {
//     FirstName string `db:"first_name"`
//     LastName  string `db:"last_name"`
//     Email     string
// }

// type TodoStorage struct{
// 	Conn *sqlx.DB
// }

// func NewTodoStorage(conn *sqlx.DB) *TodoStorage{
// 	return &TodoStorage{Conn: conn}
// }

var DbConnect *sqlx.DB = lib.DBClient // menggunakan koneksi dari lib.

func ListAllUsers(db *sqlx.DB){
	lib.DbConnection()

	people := []Person{}
    db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
    jason, john := people[0], people[1]

    fmt.Printf("%#v\n%#v", jason, john)
	
	
	// tx := db.MustBegin()
	// tx.Select(`SELECT * FROM "users"`)
}