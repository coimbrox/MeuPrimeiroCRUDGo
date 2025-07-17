package main

import (
	"log"

	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/database/mongodb"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/configuration/logger"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/routes"
	controller "github.com/coimbrox/MeuPrimeiroCRUDGo/src/controller/user"
	"github.com/coimbrox/MeuPrimeiroCRUDGo/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start the application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}


	mongodb.InitConnection()

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}