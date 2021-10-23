package main

import (
	"github.com/borischen0203/litclock-service/config"
	"github.com/borischen0203/litclock-service/logger"
	"github.com/borischen0203/litclock-service/router"

	_ "github.com/joho/godotenv/autoload"
)

func Setup() {
	logger.Setup()
	config.Setup()
	router.Setup()
}

func main() {
	Setup()
	router.Router.Run()
}
