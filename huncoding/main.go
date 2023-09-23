package main

import (
	"github.com/Junior-Shyko/Go-Lang/src/controller/routes"
	"log"  
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
  }