package handlers

import (
	"net/http"

	"github.com/ruan-molinari/go-microservices/data"
)

// swagger:route POST /product products createProduct
// Create a new product
//
// responses:
//	200: productResponse
//	422: errorValidation
//  501: errorResponse

// Create hangles POST requests to add new products
func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(&prod)
}
