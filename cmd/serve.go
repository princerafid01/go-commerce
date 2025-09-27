package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManger()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	// router
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", wrappedMux) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
