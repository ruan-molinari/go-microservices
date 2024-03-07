package handlers

import (
	"net/http"

	"github.com/ruan-molinari/go-microservices/data"
)

// swagger:route PUT /products products updateProduct
// Update a product's details
//
// responses:
//  201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handlef PUT requests to update products
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("[DEBUG] updating record with id: ", prod.ID)

	err := data.UpdateProduct(prod)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in the database"}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
