package pizzas

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/explerror"
)

// createPizza creates a new pizza.
//
// This endpoint receives form-data in the body.
// The file-field has an image of Content-Type image/*
// The data-field has an object as Content-Type: application/json
//
// @Summary Create a new pizza
// @Description Create a new pizza with image and data
// @Tags Pizzas
// @Accept multipart/form-data
// @Param file formData file true "Image file (Content-Type: image/*)"
// @Param data formData string true "JSON data (Content-Type: application/json)"
// @Produce json
// @Success 201 {object} pizzas.Pizza
// @Failure 400 {object} main.ErrorResponse
// @Failure 500 {object} main.ErrorResponse
// @Router /pizzas [post]
func createPizza(w http.ResponseWriter, r *http.Request) {
	explerror.NotImplemented(w, errors.New("create pizza handler is not implemented yet"))
}

// @Summary Get pizzas
// @Description Retrieves a list of all pizzas.
// @Tags Pizzas
// @Produce json
// @Success 200 {array} pizzas.Pizza
// @Failure 404 {object} main.ErrorResponse
// @Router /pizzas [get]
func getPizzas(w http.ResponseWriter, r *http.Request) {
	explerror.NotImplemented(w, errors.New("get pizzas handler is not implemented yet"))
}

// @Summary Get pizzas
// @Description Retrieves a list of all pizzas.
// @Tags Pizzas
// @Param pizzaId path string true "example id"
// @Produce json
// @Success 200 {object} pizzas.Pizza
// @Failure 404 {object} main.ErrorResponse
// @Router /pizzas/{pizzaId} [get]
func getPizza(w http.ResponseWriter, r *http.Request) {
	pizzaId := chi.URLParam(r, "pizzaId")
	_ = pizzaId
	explerror.NotImplemented(w, errors.New("get pizza by id handler is not implemented yet"))
}

// updatePizza updates a pizza.
//
// This endpoint receives form-data in the body.
// The file-field has an image of Content-Type image/*
// The data-field has an object as Content-Type: application/json
//
// @Summary Updates a pizza
// @Description Updates a pizza with image and/or data by it's id
// @Tags Pizzas
// @Accept multipart/form-data
// @Param pizzaId path string true "example id"
// @Param file formData file false "Image file (Content-Type: image/*)"
// @Param data formData string false "JSON data (Content-Type: application/json)"
// @Produce json
// @Success 204
// @Failure 400 {object} main.ErrorResponse
// @Failure 404 {object} main.ErrorResponse
// @Failure 500 {object} main.ErrorResponse
// @Router /pizzas [post]
func updatePizza(w http.ResponseWriter, r *http.Request) {
	pizzaId := chi.URLParam(r, "pizzaId")
	_ = pizzaId
	explerror.NotImplemented(w, errors.New("update pizza handler is not implemented yet"))
}

// deletePizza deletes a pizza.
//
// This endpoint receives form-data in the body.
// The file-field has an image of Content-Type image/*
// The data-field has an object as Content-Type: application/json
//
// @Summary Deletes a pizza
// @Description Deletes a pizza by it's id
// @Tags Pizzas
// @Param pizzaId path string true "example id"
// @Produce json
// @Success 204
// @Failure 400 {object} main.ErrorResponse
// @Failure 404 {object} main.ErrorResponse
// @Failure 500 {object} main.ErrorResponse
// @Router /pizzas [delete]
func deletePizza(w http.ResponseWriter, r *http.Request) {
	pizzaId := chi.URLParam(r, "pizzaId")
	_ = pizzaId
	explerror.NotImplemented(w, errors.New("delete pizza is not implemented yet"))
}
