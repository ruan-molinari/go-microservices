package data

import (
	"fmt"
	"time"
)

// swagger:model
type Product struct {
	// The id for this Product
	//
	// required: true
	// min: 1
	ID          int       `json:"id"`
	Name        string    `json:"name" validate:"min=3,max=24"`
	Description string    `json:"description"`
	Price       float32   `json:"price" validate:"gt=0"`
	SKU         string    `json:"sku" validate:"required,sku"`
	CreatedOn   time.Time `json:"-"`
	UpdatedOn   time.Time `json:"-"`
	DeletedOn   time.Time `json:"-"`
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
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffee without milk",
		Price:       1.99,
		SKU:         "frd32",
		CreatedOn:   time.Now(),
		UpdatedOn:   time.Now(),
	},
}
