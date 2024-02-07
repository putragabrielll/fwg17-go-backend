package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product tags
func ListAllpt(filterby string, filter string, sortby string, order string, limit int, offset int) ([]services.Pro_TagsNet, error) {
	sql := `
	SELECT 
	"pt"."id", 
    "product"."name" AS "namaProduct", 
    "tag"."name" AS "namaTags",
	"pt"."createdAt",
	"pt"."updatedAt"
	FROM "productTags" "pt"
	INNER JOIN "products" "product" ON "product"."id" = "pt"."productId"
    INNER JOIN "tags" "tag" ON "tag"."id" = "pt"."tagsId"
	WHERE "` + filterby + `"."name" ILIKE $1
	ORDER BY "` + sortby + `"."name" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.Pro_TagsNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func Countpt(filterby string, filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("pt"."id") AS "counts"
    FROM "productTags" "pt"
	INNER JOIN "products" "product" ON "product"."id" = "pt"."productId"
    INNER JOIN "tags" "tag" ON "tag"."id" = "pt"."tagsId"
	WHERE "` + filterby + `"."name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT product tags BY id
func FindPT(id int) (services.Pro_TagsNet, error){
	sql := `
	SELECT 
	"pt"."id", 
    "product"."name" AS "namaProduct", 
    "tag"."name" AS "namaTags",
	"pt"."createdAt",
	"pt"."updatedAt"
	FROM "productTags" "pt" 
	INNER JOIN "products" "product" ON "product"."id" = "pt"."productId" 
    INNER JOIN "tags" "tag" ON "tag"."id" = "pt"."tagsId" 
	WHERE "pt"."id"=$1`
	data := services.Pro_TagsNet{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE product tags
func CreatePT(data services.Pro_Tags) (services.Pro_Tags, error){
	sql := `
	INSERT INTO "productTags"
    ("productId", "tagsId")
    VALUES
    (:productId, :tagsId)
    RETURNING *
    `
	returning := services.Pro_Tags{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE product tags
func UpdatePT(data services.Pro_Tags) (services.Pro_Tags, error){
	sql := `
	UPDATE "productTags" SET 
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"tagsId"=COALESCE(NULLIF(:tagsId, 0),"tagsId"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Pro_Tags{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE product tags
func DeletePT(id int) (services.Pro_Tags, error){
	sql := `DELETE FROM "productTags" WHERE "id"= $1 RETURNING *`
	data := services.Pro_Tags{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}