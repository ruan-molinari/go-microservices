package handlers

import (
	"net/http"

	"github.com/ruan-molinari/go-microservices/data"
)

// swagger:route GET /products products listProducts
// Returns a list of products
//
// responses:
//
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request) {
	p.l.Println("[DEBUG] get all records")

	prods := data.GetProducts()

	err := data.ToJSON(prods, w)
	if err != nil {
		p.l.Println("[ERROR] serializing products")
	}
}

func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
	id := getProductID(r)

	p.l.Println("[DEBUG] get record of id: ", id)

	prod, err := data.GetProductById(id)

	switch err {
	case nil:

	case data.ErrProductNotFound:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return

	default:
		p.l.Println("[ERROR] fetching product", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	err = data.ToJSON(prod, w)
	if err != nil {
		p.l.Println("[ERROR] serializing product", err)
	}
}
