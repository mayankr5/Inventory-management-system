package main

import (
	handler "github.com/inventory-management-system/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	//Product API's
	app.GET("/products", handler.GetProducts)
	app.GET("/products/{id}", handler.GetProductByID)
	app.POST("/products", handler.AddProduct)
	app.PUT("/products/{id}", handler.UpdateProduct)
	app.DELETE("/products/{id}", handler.DeleteProduct)

	// Categories API's
	app.GET("/categories", handler.GetCategories)
	app.GET("/categories/{id}", handler.GetCategoryByID)
	app.POST("/categories", handler.AddCategory)
	app.PUT("/categories/{id}", handler.UpdateCategory)
	app.DELETE("/categories/{id}", handler.DeleteCategory)

	app.Start()
}
