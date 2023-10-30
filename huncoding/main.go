package main

import (
	"log"

	"github.com/Junior-Shyko/Go-Lang/src/controller/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}
