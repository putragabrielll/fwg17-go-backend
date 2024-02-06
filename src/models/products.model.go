package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * products
func ListAllProducts(filter string, sortby string, order string, limit int, offset int) ([]services.ProductsNet, error) {
	sql := `
	SELECT * 
	FROM "products"
	WHERE "name" ILIKE $1
	ORDER BY "`+sortby+`" `+order+`
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.ProductsNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountAllProducts(filter string) (int, error){
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "products"
	WHERE "name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT users BY id
func FindProductsId(id int) (services.ProductsNet, error){
	sql := `SELECT * FROM "products" WHERE "id"=$1`
	data := services.ProductsNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// CREATE products
func CreateProducts(data services.Products) (services.ProductsNet, error){
	sql := `
	INSERT INTO "products"
    ("name", "price", "image", "description", "discount", "isRecommended", "qty", "isActive")
    VALUES
    (:name, :price, :image, :description, :discount, :isRecommended, :qty, :isActive)
    RETURNING *
    `
	returning := services.ProductsNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next(){
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE products
func UpdateProducts(data services.Products) (services.ProductsNet, error){
	sql := `
	UPDATE "products" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"price"=COALESCE(NULLIF(:price, 0),"price"),
	"image"=COALESCE(NULLIF(:image,''),"image"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"discount"=COALESCE(NULLIF(:discount, 0),"discount"),
	"isRecommended"=COALESCE(NULLIF(:isRecommended, false),"isRecommended"),
	"qty"=COALESCE(NULLIF(:qty, 0),"qty"),
	"isActive"=COALESCE(NULLIF(:isActive, false),"isActive"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.ProductsNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)

	dataaja := data.IsRecommended
	dataaja2 := *dataaja
	fmt.Println(dataaja2)
	fmt.Println(data.IsActive)
	
	for rows.Next(){
		rows.StructScan(&returning)
	}

	return returning, err
}

// DELETE products BY id
func DeleteProducts(id int) (services.ProductsNet, error){
	sql := `DELETE FROM "products" WHERE "id"= $1 RETURNING *`
	data := services.ProductsNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}