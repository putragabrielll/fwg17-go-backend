package models

import (
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product size
func ListAllps() ([]services.PsNet, error) {
	sql := `
	SELECT * 
	FROM "productSize"
	`
	data := []services.PsNet{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}

// SELECT products size BY id
func FindPSId(id int) (services.PsNet, error) {
	sql := `SELECT * FROM "productSize" WHERE "id"=$1`
	data := services.PsNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// UPDATE product size
func UpdatePS(data services.Ps) (services.PsNet, error) {
	sql := `
	UPDATE "productSize" SET
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.PsNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}
