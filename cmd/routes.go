package cmd

import (
	"ecommerce/handlers"
	"ecommerce/middleware"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle(
		"GET /route",
		manager.With(
			http.HandlerFunc(handlers.Test),
		),
	)

	mux.Handle("GET /api/products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /api/products", manager.With(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /api/products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))

}
