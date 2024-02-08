package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * order details
func ListAllod(filter string, sortby string, order string, limit int, offset int) ([]services.OrdersDetailsNet, error) {
	sql := `
	SELECT 
	"od"."id", 
    "order"."orderNumber",
    "produk"."name" AS "namaProduct",
    "size"."size",
    "variant"."name",
    "od"."qty",
    "od"."subTotal",
    "od"."createdAt",
    "od"."updatedAt"
	FROM "orderDetails" "od"
    INNER JOIN "orders" "order" ON "order"."id" = "od"."ordersId"
    INNER JOIN "products" "produk" ON "produk"."id" = "od"."productId"
    INNER JOIN "productSize" "size" ON "size"."id" = "od"."productSizeId"
    INNER JOIN "productVariant" "variant" ON "variant"."id" = "od"."productVariantId"
	WHERE "order"."orderNumber" ILIKE $1
	ORDER BY "od"."` + sortby + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.OrdersDetailsNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	fmt.Println(err)
	return data, err
}

// Count total data
func Countod(filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("od"."id") AS "counts"
    FROM "orderDetails" "od"
    INNER JOIN "orders" "order" ON "order"."id" = "od"."ordersId"
    INNER JOIN "products" "produk" ON "produk"."id" = "od"."productId"
    INNER JOIN "productSize" "size" ON "size"."id" = "od"."productSizeId"
    INNER JOIN "productVariant" "variant" ON "variant"."id" = "od"."productVariantId"
	WHERE "order"."orderNumber" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT order details BY id
func Findod(id int) (services.OrdersDetailsNet, error){
	sql := `
	SELECT 
	"od"."id", 
    "order"."orderNumber",
    "produk"."name" AS "namaProduct",
    "size"."size",
    "variant"."name",
    "od"."qty",
    "od"."subTotal",
    "od"."createdAt",
    "od"."updatedAt"
	FROM "orderDetails" "od"
    INNER JOIN "orders" "order" ON "order"."id" = "od"."ordersId"
    INNER JOIN "products" "produk" ON "produk"."id" = "od"."productId"
    INNER JOIN "productSize" "size" ON "size"."id" = "od"."productSizeId"
    INNER JOIN "productVariant" "variant" ON "variant"."id" = "od"."productVariantId"
	WHERE "od"."id"=$1
	`
	data := services.OrdersDetailsNet{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE order details
func Createod(data services.OrdersDetails) (services.OrdersDetails, error){
	sql := `
	INSERT INTO "orderDetails"
    ("ordersId", "productId", "productSizeId", "productVariantId", "qty", "subTotal")
    VALUES
    (:ordersId, :productId, :productSizeId, :productVariantId, :qty, :subTotal)
    RETURNING *
    `
	returning := services.OrdersDetails{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil{
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE order details
func Updateod(data services.OrdersDetails) (services.OrdersDetails, error){
	sql := `
	UPDATE "orderDetails" SET 
	"ordersId"=COALESCE(NULLIF(:ordersId, 0),"ordersId"),
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"productSizeId"=COALESCE(NULLIF(:productSizeId, 0),"productSizeId"),
	"productVariantId"=COALESCE(NULLIF(:productVariantId, 0),"productVariantId"),
	"qty"=COALESCE(NULLIF(:qty, 0),"qty"),
	"subTotal"=COALESCE(NULLIF(:subTotal, 0),"subTotal"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.OrdersDetails{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE order details
func Deleteod(id int) (services.OrdersDetails, error){
	sql := `DELETE FROM "orderDetails" WHERE "id"= $1 RETURNING *`
	data := services.OrdersDetails{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}