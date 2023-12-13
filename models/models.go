package models

type Product struct {
	ID         int
	Name       string
	Price      float64
	Quantity   int
	CategoryID int
}

type Category struct {
	ID   int
	Name string
}
