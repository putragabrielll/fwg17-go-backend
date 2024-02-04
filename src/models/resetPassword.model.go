package models


import (
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)



// CREATE otp
func CreateRP(data services.FormReset) (services.FormReset, error){
	sql := `
	INSERT INTO "forgotPassword"
    ("email", "otp")
    VALUES
    (:email, :otp)
    RETURNING *
    `
	returning := services.FormReset{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	// if err != nil {
	// 	return returning, err
	// }
	for rows.Next(){ // rows.Next() => akan mengembalikan boolean.
		rows.StructScan(&returning)
	}
	return returning, err
}


// SELECT otp BY email
func FindRPByOTP(otp string) (services.FormReset, error){
	sql := `SELECT * FROM "forgotPassword" WHERE "otp"=$1`
	data := services.FormReset{}
	err := lib.DbConnection().Get(&data, sql, otp) // id diambil dari parameter id.
	return data, err
}


// DELETE otp BY email
func DeleteOTP(id int) (services.FormReset, error){
	sql := `DELETE FROM "forgotPassword" WHERE "id"= $1`
	data := services.FormReset{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}