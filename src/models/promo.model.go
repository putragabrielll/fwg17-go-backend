package models

import (
	"fmt"
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)

// SELECT * promo
func ListAllPromo(filter string, sortby string, order string, limit int, offset int) ([]services.PromoNet, error) {
	sql := `
	SELECT * 
	FROM "promo"
	WHERE "name" ILIKE $1
	ORDER BY "` + sortby + `" ` + order + `
	LIMIT $2
	OFFSET $3
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	data := []services.PromoNet{}
	err := lib.DbConnection().Select(&data, sql, fmtsearch, limit, offset)
	return data, err
}

// Count total data
func CountAllPromo(filter string) (int, error) {
	var count int
	sql := `
	SELECT COUNT("id") AS "counts"
    FROM "promo"
	WHERE "name" ILIKE $1
	`
	fmtsearch := fmt.Sprintf("%%%v%%", filter)
	err := lib.DbConnection().Get(&count, sql, fmtsearch)
	return count, err
}

// SELECT promo BY id
func FindPromoId(id int) (services.PromoNet, error) {
	sql := `SELECT * FROM "promo" WHERE "id"=$1`
	data := services.PromoNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}

// CREATE promo
func CreatePromo(data services.Promo) (services.PromoNet, error){
	temp1 := data.Percentage
	temp2 := *temp1
	fmtpersen := fmt.Sprintf("%.2f", temp2)
	sql := `
	INSERT INTO "promo"
    ("name", "code", "description", "percentage", "isExpired", "maximumPromo", "minimumAmount")
    VALUES
    (:name, :code, :description, `+fmtpersen+`, :isExpired, :maximumPromo, :minimumAmount)
    RETURNING *
    `
	returning := services.PromoNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next() {
		rows.StructScan(&returning)
	}
	return returning, err
}

// UPDATE PROMO
func UpdatePromo(data services.Promo) (services.PromoNet, error){
	temp1 := data.Percentage
	temp2 := *temp1
	fmtpersen := fmt.Sprintf("%.2f", temp2)
	sql := `
	UPDATE "promo" SET 
	"name"=COALESCE(NULLIF(:name,''),"name"),
	"code"=COALESCE(NULLIF(:code, ''),"code"),
	"description"=COALESCE(NULLIF(:description,''),"description"),
	"percentage"=COALESCE(NULLIF(`+fmtpersen+`,0),"percentage"),
	"isExpired"=COALESCE(NULLIF(:isExpired, false),"isExpired"),
	"maximumPromo"=COALESCE(NULLIF(:maximumPromo, 0),"maximumPromo"),
	"minimumAmount"=COALESCE(NULLIF(:minimumAmount, 0),"minimumAmount"),
	"updatedAt"=NOW()
    WHERE id=:id
    RETURNING *
    `
	returning := services.PromoNet{}
	rows, err := lib.DbConnection().NamedQuery(sql, data)
	for rows.Next(){
		rows.StructScan(&returning)
	}

	return returning, err
}

// DELETE promo BY id
func DeletePromo(id int) (services.PromoNet, error){
	sql := `DELETE FROM "promo" WHERE "id"= $1 RETURNING *`
	data := services.PromoNet{}
	err := lib.DbConnection().Get(&data, sql, id) // id diambil dari parameter id.
	return data, err
}