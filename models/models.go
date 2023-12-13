package models

type Product struct {
	ID       int
	Name     string
	Price    float64
	Quantity int
	Category string
}

type Category struct {
	ID   int
	Name string
}
