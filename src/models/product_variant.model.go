package models

import (
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product variant
func ListAllpv() ([]services.PvNet, error){
	sql := `
	SELECT * 
	FROM "productVariant"
	`
	data := []services.PvNet{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}

// SELECT product variant BY id
func FindPVId(id int) (services.PvNet, error){
	sql := `SELECT * FROM "productVariant" WHERE "id"=$1`
	data := services.PvNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// CREATE product variant
func CreatePv(data services.Pv) (services.PvNet, error){
	sql := `
	INSERT INTO "productVariant"
    ("name", "additionalPrice")
    VALUES
    (:name, :additionalPrice)
    RETURNING *
    `
	returning := services.PvNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next(){
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE product variant
func UpdatePv(data services.Pv) (services.PvNet, error){
	sql := `
	UPDATE "productVariant" SET
	"name"=COALESCE(NULLIF(:name, ''),"name"),
	"additionalPrice"=COALESCE(NULLIF(:additionalPrice, 0),"additionalPrice"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.PvNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE product variant
func DeletePv(id int) (services.PvNet, error){
	sql := `DELETE FROM "productVariant" WHERE "id"= $1 RETURNING *`
	data := services.PvNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}