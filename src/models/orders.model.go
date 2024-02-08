package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * orders
func ListOrders(table string, column string, filter string, order string, limit int, offset int) ([]services.OrdersNet, error) {
	sql := `
	SELECT 
	"o"."id", 
    "user"."fullName" AS "userName",
    "o"."orderNumber",
    "prom"."name" AS "promoName",
    "o"."total",
    "o"."taxAmount",
    "o"."status",
    "o"."deliveryAddress",
    "o"."fullName",
    "o"."email",
    "o"."createdAt",
    "o"."updatedAt"
	FROM "orders" "o"
	INNER JOIN "users" "user" ON "user"."id" = "o"."usersId"
    INNER JOIN "promo" "prom" ON "prom"."id" = "o"."promoId"
	WHERE "` + table + `"."` + column + `" ILIKE $1
	ORDER BY "` + table + `"."` + column + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.OrdersNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountOrders(table string, column string, filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("o"."id") AS "counts"
    FROM "orders" "o"
	INNER JOIN "users" "user" ON "user"."id" = "o"."usersId"
    INNER JOIN "promo" "prom" ON "prom"."id" = "o"."promoId"
	WHERE "` + table + `"."` + column + `" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT orders BY id
func FindOrders(id int) (services.OrdersNet, error) {
	sql := `
	SELECT 
	"o"."id", 
    "user"."fullName" AS "userName",
    "o"."orderNumber",
    "prom"."name" AS "promoName",
    "o"."total",
    "o"."taxAmount",
    "o"."status",
    "o"."deliveryAddress",
    "o"."fullName",
    "o"."email",
    "o"."createdAt",
    "o"."updatedAt"
	FROM "orders" "o"
	INNER JOIN "users" "user" ON "user"."id" = "o"."usersId"
    INNER JOIN "promo" "prom" ON "prom"."id" = "o"."promoId"
	WHERE "o"."id"=$1
	`
	data := services.OrdersNet{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE orders
func CreateOrders(data services.Orders) (services.Orders, error) {
	sql := `
	INSERT INTO "orders"
    ("usersId", 
	"orderNumber",
	"promoId", 
	"total", 
	"taxAmount", 
	"status", 
	"deliveryAddress", 
	"fullName", 
	"email")
    VALUES
    (:usersId, :orderNumber, :promoId, :total, :taxAmount, :status, :deliveryAddress, :fullName, :email)
    RETURNING *
    `
	returning := services.Orders{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE orders
func UpdateOrders(data services.Orders) (services.Orders, error) {
	sql := `
	UPDATE "orders" SET 
	"usersId"=COALESCE(NULLIF(:usersId, 0),"usersId"),
	"orderNumber"=COALESCE(NULLIF(:orderNumber, ''),"orderNumber"),
	"promoId"=COALESCE(NULLIF(:promoId, 0),"promoId"),
	"total"=COALESCE(NULLIF(:total, 0),"total"),
	"taxAmount"=COALESCE(NULLIF(:taxAmount, 0),"taxAmount"),
	"status"=:status,
	"deliveryAddress"=COALESCE(NULLIF(:deliveryAddress, ''),"deliveryAddress"),
	"fullName"=COALESCE(NULLIF(:fullName, ''),"fullName"),
	"email"=COALESCE(NULLIF(:email, ''),"email"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Orders{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE orders
func DeleteOrders(id int) (services.Orders, error){
	sql := `DELETE FROM "orders" WHERE "id"= $1 RETURNING *`
	data := services.Orders{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}