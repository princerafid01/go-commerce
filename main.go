package main

import (
	"ecommerce/cmd"
	"ecommerce/config"
)

func main() {
	cnf := config.GetConfig()

	cmd.Serve()
}
