package models

import (
	"github.com/putragabrielll/go-backend/src/lib"
	"github.com/putragabrielll/go-backend/src/services"
)


// SELECT * products
func ListAllProducts() ([]services.Products, error) {
	sql := `SELECT * FROM "products"`
	data := []services.Products{}
	err := lib.DbConnection().Select(&data, sql)
	return data, err
}