package handlers

import (
	"encoding/json"

	"github.com/inventory-management-system/models"
	"gofr.dev/pkg/gofr"
)

func AddProduct(ctx *gofr.Context) (interface{}, error) {
	var product models.Product
	err := json.NewDecoder(ctx.Request().Body).Decode(&product)

	if err != nil {
		ctx.Request().Response.Status = "404 Something went wrong"
	}

	return "addprod", nil
}
