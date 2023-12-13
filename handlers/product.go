package handler

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/inventory-management-system/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetProducts(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, " SELECT pd.product_id as ID, pd.name as Name, pd.quantity as Quantity, pd.Price as Price, ctg.name as Category FROM product pd inner join categories ctg on pd.category_id = ctg.category_id")
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	products := make([]models.Product, 0)

	for rows.Next() {
		var product models.Product

		err = rows.Scan(&product.ID, &product.Name, &product.Quantity, &product.Price, &product.Category)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		products = append(products, product)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return products, nil
}

func GetProductByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var resp models.Product

	err := ctx.DB().QueryRowContext(ctx, " SELECT pd.product_id as ID, pd.name as Name, pd.quantity as Quantity, pd.Price as Price, ctg.name as Category FROM product pd inner join categories ctg on pd.category_id = ctg.category_id where product_id=$1", id).
		Scan(&resp.ID, &resp.Name, &resp.Quantity, &resp.Price, &resp.Category)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: id}
	}

	return resp, nil
}
