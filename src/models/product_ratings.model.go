package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * product ratings
func ListAllpr(filter string, sortby string, order string, limit int, offset int) ([]services.Pro_RateNet, error) {
	sql := `
	SELECT 
	"pr"."id", 
    "product"."name" AS "namaProduct", 
	"pr"."rate",
	"pr"."reviewMessege",
    "user"."fullName" AS "namaUser",
	"pr"."createdAt",
	"pr"."updatedAt"
	FROM "productRatings" "pr"
	INNER JOIN "products" "product" ON "product"."id" = "pr"."productId"
    INNER JOIN "users" "user" ON "user"."id" = "pr"."usersId"
	WHERE "user"."fullName" ILIKE $1
	ORDER BY "` + sortby + `"."name" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.Pro_RateNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func Countpr(filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("pr"."id") AS "counts"
    FROM "productRatings" "pr"
	INNER JOIN "products" "product" ON "product"."id" = "pr"."productId"
    INNER JOIN "users" "user" ON "user"."id" = "pr"."usersId"
	WHERE "user"."fullName" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT product ratings BY id
func Findpr(id int) (services.Pro_RateNet, error) {
	sql := `
	SELECT 
	"pr"."id", 
    "product"."name" AS "namaProduct", 
	"pr"."rate",
	"pr"."reviewMessege",
    "user"."fullName" AS "namaUser",
	"pr"."createdAt",
	"pr"."updatedAt"
	FROM "productRatings" "pr"
	INNER JOIN "products" "product" ON "product"."id" = "pr"."productId"
    INNER JOIN "users" "user" ON "user"."id" = "pr"."usersId"
	WHERE "pr"."id"=$1`
	data := services.Pro_RateNet{}
	err := lib.DbConnection().Get(&data, sql, id)
	return data, err
}

// CREATE product ratings
func Createpr(data services.Pro_Rate) (services.Pro_Rate, error) {
	sql := `
	INSERT INTO "productRatings"
    ("productId", "rate", "reviewMessege", "usersId")
    VALUES
    (:productId, :rate, :reviewMessege, :usersId)
    RETURNING *
    `
	returning := services.Pro_Rate{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE product ratings
func Updatepr(data services.Pro_Rate) (services.Pro_Rate, error){
	sql := `
	UPDATE "productRatings" SET 
	"productId"=COALESCE(NULLIF(:productId, 0),"productId"),
	"rate"=COALESCE(NULLIF(:rate, 0),"rate"),
	"reviewMessege"=COALESCE(NULLIF(:reviewMessege, ''),"reviewMessege"),
	"usersId"=COALESCE(NULLIF(:usersId, 0),"usersId"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.Pro_Rate{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	if err != nil {
		return returning, err
	}
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// DELETE product ratings
func Deletepr(id int) (services.Pro_Rate, error){
	sql := `DELETE FROM "productRatings" WHERE "id"= $1 RETURNING *`
	data := services.Pro_Rate{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}