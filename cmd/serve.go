package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// router
	mux := http.NewServeMux()
	manager := middleware.NewManger()
	manager.Use(middleware.Logger, middleware.Hudai, middleware.CorsWithPreflight)

	initRoutes(mux, manager)

	wrappedMux := manager.With(mux)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", wrappedMux) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
