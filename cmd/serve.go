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
	manager.Use()

	// global middleware
	// globalMiddlewares := []middleware.Middleware{
	// 	middleware.Hudai,
	// 	middleware.Logger,
	// 	middleware.CorsWithPreflight,
	// }
	manager.Use(middleware.Hudai, middleware.Logger, middleware.CorsWithPreflight)

	wrappedMux := manager.WrapMux(
		globalMiddlewares,
		mux,
	)

	initRoutes(mux, manager)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", wrappedMux) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
