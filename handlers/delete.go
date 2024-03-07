package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/ruan-molinari/go-microservices/data"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Delete a product from the database
//
// responses:
//
//	201: noContent
//	404: errorResponse
//	501: errorResponse

// Deletes a product from the database
func (p *Products) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println(w, "Unable to convert ID", err)
		return
	}

	err = data.DeleteProduct(id)
	if err == data.ErrProductNotFound {
		p.l.Println("[ERROR] product not found")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}
	if err != nil {
		p.l.Println("[ERROR] deleting product")

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
