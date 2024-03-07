package data

import (
	"fmt"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// The id for this product
	//
	// required: true
	// min: 1
	ID int `json:"id"`

	// The name ot the product
	//
	// required: true
	// min: 3
	// max: 255
	Name string `json:"name" validate:"required,min=3,max=255"`

	// The description of the product
	//
	// required: false
	// max length: 10000
	Description string `json:"description" validate:"max=10000"`

	// The price of the product
	//
	// required: true
	// min: 0.1
	Price float32 `json:"price" validate:"required,gt=0"`

	// The SKU for the product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU string `json:"sku" validate:"required,sku"`
}

// `Products` is a collection of `Product`
type Products []*Product

func GetProducts() Products {
	return productList
}

func GetProductById(id int) (*Product, error) {
	i := findIndexByProductId(id)

	if i == -1 {
		return nil, ErrProductNotFound
	}

	return productList[i], nil
}

func AddProduct(p *Product) {
	p.ID = getNextId()
	productList = append(productList, p)
}

func UpdateProduct(p Product) error {
	i := findIndexByProductId(p.ID)

	if i == -1 {
		return ErrProductNotFound
	}
	productList[i] = &p

	return nil
}

func DeleteProduct(id int) error {
	i := findIndexByProductId(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i:]...)

	return nil
}

func getNextId() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrProductNotFound = fmt.Errorf("Product not found")

func findIndexByProductId(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = Products{
	{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc323",
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "frd32",
	},
}
