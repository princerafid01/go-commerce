package rest

import (
	"ecommerce/config"
	"ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	manager := middleware.NewManger()
	manager.Use(middleware.Preflight, middleware.Cors, middleware.Logger)

	// router
	mux := http.NewServeMux()
	wrappedMux := manager.WrapMux(mux)

	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort)
	fmt.Println("Server running on port ", addr)
	err := http.ListenAndServe(addr, wrappedMux) // "Failed to start server"
	if err != nil {
		fmt.Println("Error starting server ", err)
		os.Exit(1)
	}
}
