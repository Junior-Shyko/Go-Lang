package router

import (
	"api/src/router/rotas"
	"github.com/gorilla/mux"
)

//Gerar Rotas
func Gerar() *mux.Router {
	r := mux.NewRouter()

	return rotas.Configurar(r)
}