/*
 * Open API Sample
 *
 * This is a sample code.
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
	errorHandler ErrorHandler
}

// DefaultApiOption for how the controller is set up.
type DefaultApiOption func(*DefaultApiController)

// WithDefaultApiErrorHandler inject ErrorHandler into controller
func WithDefaultApiErrorHandler(h ErrorHandler) DefaultApiOption {
	return func(c *DefaultApiController) {
		c.errorHandler = h
	}
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer, opts ...DefaultApiOption) Router {
	controller := &DefaultApiController{
		service:      s,
		errorHandler: DefaultErrorHandler,
	}

	for _, opt := range opts {
		opt(controller)
	}

	return controller
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"AddPet",
			strings.ToUpper("Post"),
			"/v1/pets",
			c.AddPet,
		},
		{
			"PetsGet",
			strings.ToUpper("Get"),
			"/v1/pets",
			c.PetsGet,
		},
		{
			"PetsIdDelete",
			strings.ToUpper("Delete"),
			"/v1/pets/{id}",
			c.PetsIdDelete,
		},
		{
			"PetsIdGet",
			strings.ToUpper("Get"),
			"/v1/pets/{id}",
			c.PetsIdGet,
		},
		{
			"TestGet",
			strings.ToUpper("Get"),
			"/v1/test",
			c.TestGet,
		},
	}
}

// AddPet - 
func (c *DefaultApiController) AddPet(w http.ResponseWriter, r *http.Request) {
	newPet := NewPet{}
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	if err := d.Decode(&newPet); err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}
	if err := AssertNewPetRequired(newPet); err != nil {
		c.errorHandler(w, r, err, nil)
		return
	}
	result, err := c.service.AddPet(r.Context(), newPet)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PetsGet - 
func (c *DefaultApiController) PetsGet(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	tags := strings.Split(query.Get("tags"), ",")
	result, err := c.service.PetsGet(r.Context(), tags)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PetsIdDelete - 
func (c *DefaultApiController) PetsIdDelete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := parseInt64Parameter(params["id"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.PetsIdDelete(r.Context(), id)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// PetsIdGet - 
func (c *DefaultApiController) PetsIdGet(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := parseInt64Parameter(params["id"], true)
	if err != nil {
		c.errorHandler(w, r, &ParsingError{Err: err}, nil)
		return
	}

	result, err := c.service.PetsIdGet(r.Context(), id)
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}

// TestGet - 
func (c *DefaultApiController) TestGet(w http.ResponseWriter, r *http.Request) {
	result, err := c.service.TestGet(r.Context())
	// If an error occurred, encode the error with the status code
	if err != nil {
		c.errorHandler(w, r, err, &result)
		return
	}
	// If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)

}
