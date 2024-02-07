package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * tags
func ListAllTags(filter string, sortby string, order string, limit int, offset int) ([]services.Tags, error){
	sql := `
	SELECT * 
	FROM "tags"
	WHERE "name" ILIKE $1
	ORDER BY "` + sortby + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.Tags{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountTags(filter string) (int, error){
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "tags"
	WHERE "name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT tags BY id
func FindTags(id int) (services.Tags, error){
	sql := `SELECT * FROM "tags" WHERE "id"=$1`
	data := services.Tags{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE tags
func CreateTags(data services.Tags) (services.Tags, error){
	sql := `
	INSERT INTO "tags"
    ("name")
    VALUES
    (:name)
    RETURNING *
    `
	returning := services.Tags{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE tags
func UpdateTags(data services.Tags) (services.Tags, error){
	sql := `
	UPDATE "tags" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Tags{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE tags
func DeleteTags(id int) (services.Tags, error){
	sql := `DELETE FROM "tags" WHERE "id"= $1 RETURNING *`
	data := services.Tags{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}