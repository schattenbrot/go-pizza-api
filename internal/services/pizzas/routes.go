package pizzas

import "github.com/go-chi/chi/v5"

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/", createPizza)
	r.Get("/", getPizzas)
	r.Get("/{pizzaId}", getPizza)
	r.Put("/{pizzaId}", updatePizza)
	r.Delete("/{pizzaId}", deletePizza)

	return r
}
