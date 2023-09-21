package config

import (
	"fmt"
	"os"
	"strconv"
	"log"
	"github.com/joho/godotenv"
)

var (
	//string q tera a conexao com o banco
	StringConnectionDB = ""
	//Guardara a porta para conexao com o banco
	Port = 0
)

//Carrega valores para conexao com o banco
func ToLoad() {
	var errorToLoad error

	if errorToLoad = godotenv.Load(); errorToLoad != nil {
		log.Fatal(errorToLoad)
	}

	Port, errorToLoad = strconv.Atoi(os.Getenv("DB_PORT"))
	if errorToLoad != nil {
		Port = 9000
	}

	StringConnectionDB = fmt.Sprintf("mongodb://%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
}