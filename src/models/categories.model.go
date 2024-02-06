package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * categories
func ListAllCtgr(filter string, sortby string, order string, limit int, offset int) ([]services.Categories, error) {
	sql := `
	SELECT * 
	FROM "categories"
	WHERE "name" ILIKE $1
	ORDER BY "` + sortby + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.Categories{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountCategories(filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "categories"
	WHERE "name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT categories BY id
func FindCategories(id int) (services.Categories, error) {
	sql := `SELECT * FROM "categories" WHERE "id"=$1`
	data := services.Categories{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE categories
func CreateCat(data services.Categories) (services.Categories, error) {
	sql := `
	INSERT INTO "categories"
    ("name")
    VALUES
    (:name)
    RETURNING *
    `
	returning := services.Categories{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE categories
func UpdateCat(data services.Categories) (services.Categories, error) {
	sql := `
	UPDATE "categories" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Categories{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE categories
func DeleteCat(id int) (services.Categories, error){
	sql := `DELETE FROM "categories" WHERE "id"= $1 RETURNING *`
	data := services.Categories{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}
