package main

import (
	handler "github.com/inventory-management-system/handlers"
	"gofr.dev/pkg/gofr"
)

func main() {
	app := gofr.New()

	app.GET("/products", handler.GetProducts)
	app.GET("/products/{id}", handler.GetProductByID)
	app.POST("/products", handler.AddProduct)
	app.POST("/products/{id}", handler.UpdateProduct)

	// Starts the server, it will listen on the default port 8000.
	// it can be over-ridden through configs
	app.Start()
}
