package handlers

import (
	"net/http"

	"github.com/ruan-molinari/go-microservices/data"
)

func (p *Products) Create(w http.ResponseWriter, r *http.Request) {
	// fetch the product from the context
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	p.l.Printf("[DEBUG] Inserting product: %#v\n", prod)
	data.AddProduct(&prod)
}
