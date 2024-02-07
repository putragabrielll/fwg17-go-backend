package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product categories
func ListAllpc(filterby string, filter string, sortby string, order string, limit int, offset int) ([]services.Pro_CateNet, error) {
	sql := `
	SELECT 
	"pc"."id", 
    "product"."name" AS "namaProduct", 
    "categori"."name" AS "namaCategories",
	"pc"."createdAt",
	"pc"."updatedAt"
	FROM "productCategories" "pc"
	INNER JOIN "products" "product" ON "product"."id" = "pc"."productId"
    INNER JOIN "categories" "categori" ON "categori"."id" = "pc"."categoriesId"
	WHERE "` + filterby + `"."name" ILIKE $1
	ORDER BY "` + sortby + `"."name" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.Pro_CateNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func Countpc(filterby string, filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("pc"."id") AS "counts"
    FROM "productCategories" "pc"
	INNER JOIN "products" "product" ON "product"."id" = "pc"."productId"
    INNER JOIN "categories" "categori" ON "categori"."id" = "pc"."categoriesId"
	WHERE "` + filterby + `"."name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT product categories BY id
func FindPC(id int) (services.Pro_CateNet, error) {
	sql := `
	SELECT 
	"pc"."id", 
    "product"."name" AS "namaProduct", 
    "categori"."name" AS "namaCategories",
	"pc"."createdAt",
	"pc"."updatedAt"
	FROM "productCategories" "pc"
	INNER JOIN "products" "product" ON "product"."id" = "pc"."productId"
    INNER JOIN "categories" "categori" ON "categori"."id" = "pc"."categoriesId"
	WHERE "pt"."id"=$1`
	data := services.Pro_CateNet{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE product categories
func CreatePC(data services.Pro_Cate) (services.Pro_Cate, error){
	sql := `
	INSERT INTO "productCategories"
    ("productId", "categoriesId")
    VALUES
    (:productId, :categoriesId)
    RETURNING *
    `
	returning := services.Pro_Cate{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil{
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE product categories
func UpdatePC(data services.Pro_Cate) (services.Pro_Cate, error){
	sql := `
	UPDATE "productCategories" SET 
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"categoriesId"=COALESCE(NULLIF(:categoriesId, 0),"categoriesId"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Pro_Cate{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE product categories
func DeletePC(id int) (services.Pro_Cate, error){
	sql := `DELETE FROM "productCategories" WHERE "id"= $1 RETURNING *`
	data := services.Pro_Cate{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}