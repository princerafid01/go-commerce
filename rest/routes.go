package rest

import (
	"ecommerce/rest/handlers"
	"ecommerce/rest/middlewares"
	"net/http"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /api/products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /api/products", manager.With(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /api/products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("PUT /api/products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /api/products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))

	mux.Handle("POST /api/users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /api/users/login", manager.With(http.HandlerFunc(handlers.Login)))

}
