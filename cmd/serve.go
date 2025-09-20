package cmd

import (
	"ecommerce/global_router"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	// router
	mux := http.NewServeMux()
	manager := middleware.NewManger()
	manager.Use(middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)

	globalRouter := global_router.GlobalRouter(mux)

	fmt.Println("Server running on port 5000")
	err := http.ListenAndServe(":5000", globalRouter) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
	}
}
