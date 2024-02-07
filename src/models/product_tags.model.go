package models

import (
	"fmt"

	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product tags
func ListAllpt(filter string, sortby string, order string, limit int, offset int) (){
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
func Countpt() (int, error){
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "productTags"
	WHERE "name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}