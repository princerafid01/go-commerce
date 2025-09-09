package cmd

import (
	"ecommerce/global_router"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	// router
	mux := http.NewServeMux()

	mux.Handle("GET /api/products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /api/products", http.HandlerFunc(handlers.CreateProducts))
	mux.Handle("GET /api/products/{id}", http.HandlerFunc(handlers.GetProductById))

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", globalRouter) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
