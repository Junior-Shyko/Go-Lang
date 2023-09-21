package main

import (
	"net/http"
	"log"
	"api/src/router"
	"fmt"
)

func main() {
	fmt.Println("Rodando API")

	r := router.Gerar()
	// fmt.Println(r)
	log.Fatal(http.ListenAndServe(":5000", r))
}