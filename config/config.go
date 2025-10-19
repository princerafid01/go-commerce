package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
}

var configurations *Config

func loadConfig() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("failed to read the ENV variables:", err)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service Name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)

	if err != nil {
		fmt.Println("Port Must be a number")
	}

	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")

	if jwtSecretKey == "" {
		fmt.Println("JWT Secret Key is required")
		os.Exit(1)
	}

	configurations = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port),
		JwtSecretKey: jwtSecretKey,
	}

}

func GetConfig() *Config {
	if configurations == nil {
		// first time
		loadConfig()

	}

	return configurations
}
