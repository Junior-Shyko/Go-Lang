package main

import (
	"github.com/gin-gonic/gin"
	"api/src/router/config"
	"github.com/gin-contrib/cors"
	"log"
	"api/src/router"
	"fmt"
)

func main() {
	fmt.Println("Rodando API")
	// config.ToLoad()
	fmt.Println(config.Port)
	// r := router.Gerar()
	//Obrigado instanciar
	routers := gin.Default()
	
	config := cors.DefaultConfig()
  	config.AllowOrigins = []string{"http://localhost:8080"}
	routers.Use(cors.New(config))
	//Passa um ponteiro de RouterGroup
	router.InitRoutes(&routers.RouterGroup)
	
	if err := routers.Run(":5000"); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(r)
	// log.Fatal(http.ListenAndServe(":5000", r))
}