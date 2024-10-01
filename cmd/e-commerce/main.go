package main

import (
	"log"
	"os"

	registry "github.com/e-commerce/container"
	"github.com/e-commerce/system/config"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var mongoService *config.MongoService
var redisService *redis.Client
var container *registry.Container

func init() {
	var err error

	config.LoadEnv()

	mongoService, err = config.MongoInit()
	if err != nil {
		log.Fatalf("Error initializing MongoService: %v", err)
		panic(err)
	}

	redisService = config.RedisInit()

}

func main() {
	PORT := os.Getenv("PORT")

	defer mongoService.Disconnect()
	app := gin.Default()

	container = registry.InitContainer(mongoService, redisService)

	container.UserRoutes.RegisterRoutes(app)
	container.OtpRoutes.RegisterRoutes(app)

	if err := app.Run(":" + PORT); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
