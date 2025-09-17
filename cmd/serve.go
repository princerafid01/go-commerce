package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// router
	mux := http.NewServeMux()

	mux.Handle("GET /route", middleware.Logger(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /api/products", middleware.Logger(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("POST /api/products", middleware.Logger(http.HandlerFunc(handlers.CreateProducts)))
	mux.Handle("GET /api/products/{id}", middleware.Logger(http.HandlerFunc(handlers.GetProductById)))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", globalRouter) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
