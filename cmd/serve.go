package cmd

import (
	"ecommerce/config"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	cnf := config.GetConfig()

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
