package main

import (
	"api/src/router/config"
	"net/http"
	"log"
	"api/src/router"
	"fmt"
)

func main() {
	fmt.Println("Rodando API")
	config.ToLoad()
	fmt.Println(config.Port)
	r := router.Gerar()
	// fmt.Println(r)
	log.Fatal(http.ListenAndServe(":5000", r))
}