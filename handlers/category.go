package handler

import (
	"database/sql"
	"strconv"

	"github.com/inventory-management-system/models"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
)

func GetCategories(ctx *gofr.Context) (interface{}, error) {

	rows, err := ctx.DB().QueryContext(ctx, " SELECT category_id, name FROM categories")
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	defer rows.Close()

	categories := make([]models.Category, 0)

	for rows.Next() {
		var category models.Category

		err = rows.Scan(&category.ID, &category.Name)
		if err != nil {
			return nil, errors.DB{Err: err}
		}

		categories = append(categories, category)
	}

	err = rows.Err()
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return categories, nil
}

func GetCategoryByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var category models.Category

	err := ctx.DB().QueryRowContext(ctx, " SELECT category_id, name FROM categories WHERE category_id=$1", id).
		Scan(&category.ID, &category.Name)

	if err == sql.ErrNoRows {
		return nil, errors.EntityNotFound{Entity: "product", ID: id}
	}

	return category, nil
}

func AddCategory(ctx *gofr.Context) (interface{}, error) {
	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO categories(name) VALUES($1)", category.Name)
	if err != nil {
		return nil, err
	}

	return "Successfully Added", nil
}

func UpdateCategory(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	var category models.Category

	if err := ctx.Bind(&category); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}

	_, err := ctx.DB().ExecContext(ctx, " UPDATE categories SET name = $1 WHERE category_id=$2 ", category.Name, id)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return "Update successfully", nil
}

func DeleteCategory(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := strconv.Atoi(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM categories WHERE category_id=$1", id)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	return "Deleted successfully", nil

}
