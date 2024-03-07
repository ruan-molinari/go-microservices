// Package classification of Product API
//
// Documentation for Product API
//
//	Schemes: http
//	BasePath: /
//	Version: 0.1.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handlers

import "github.com/ruan-molinari/go-microservices/data"

//
// NOTE: Types defined here are used for documentation pourposes only
// those shall never be imported/used anywhere besides that

// Generic error message returned as string
// swagger:response errorResponse
type errorResponseWrapper struct {
	// Description of the error
	// in: body
	Body GenericError
}

// Validation errors defined as an array of strings
// swagger:response errorValidation
type errorValidationWrapper struct {
	// Collection of errors
	// in: body
	Body ValidationError
}

// A list of products
// swagger:response productsResponse
type productsResponseWrapper struct {
	// Newly created product
	// in: body
	Body []data.Product
}

// A single product
// swagger:response productResponse
type productResponseWrapper struct {
	// A single product
	// in: body
	Body data.Product
}

// No tontent is return by this API endpoint
// swagger:response noContentResponse
type noContentResponseWrapper struct{}

type productParamsWrapper struct {
	// Product data structure to Update or Create
	// Note: the id field is ignored by update and create operations
	// in: body
	// required: true
	Body data.Product
}

// swagger:parameters deleteProduct
type productIDParameterWrapper struct {
	// The id of the product to delete from database
	// in: path
	// required: true
	ID int `json:"id"`
}
