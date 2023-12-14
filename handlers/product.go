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

	var product models.Product

	err := ctx.DB().QueryRowContext(ctx, " SELECT pd.product_id as ID, pd.name as Name, pd.quantity as Quantity, pd.Price as Price, ctg.name as Category FROM product pd inner join categories ctg on pd.category_id = ctg.category_id where product_id=$1", id).
		Scan(&product.ID, &product.Name, &product.Quantity, &product.Price, &product.Category)
	fmt.Println(err)
	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: id}
	}

	return product, nil
}

func AddProduct(ctx *gofr.Context) (interface{}, error) {
	var product models.Product
	var category models.Category

	if err := ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	if err := ctx.DB().QueryRowContext(ctx, "SELECT category_id from categories where name=$1", product.Category).Scan(&category.ID); err == sql.ErrNoRows {
		if err := ctx.DB().QueryRowContext(ctx, "INSERT INTO categories(name) VALUES($1)", product.Category); err != nil {
			return nil, sql.ErrConnDone
		}

	}
	fmt.Print(product.Category)
	ctx.DB().QueryRowContext(ctx, "INSERT INTO product(name,quantity,price,category_id) VALUES($1,$2,$3,(SELECT category_id from categories where name = $4))", product.Name, product.Quantity, product.Price, product.Category).Scan(&product.Name, &product.Quantity, &product.Price)
	// if err != nil {
	// 	return nil, errors.DB{Err: err}
	// }

	return "Successfull added", nil
}

func UpdateProduct(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var product models.Product

	if err := ctx.Bind(&product); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	_, err := ctx.DB().ExecContext(ctx, " UPDATE product SET name=$1, quantity=$2, price=$3, category_id=(SELECT category_id from categories where name=$4) where product_id=$5 ", product.Name, product.Quantity, product.Price, product.Category, id)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return "update successfully", nil
}

func DeleteProduct(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM product WHERE product_id=$1", id)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return "Deleted successfully", nil

}
