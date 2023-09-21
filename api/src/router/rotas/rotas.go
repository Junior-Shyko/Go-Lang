package rotas

import (
	"net/http"
	"github.com/gorilla/mux"
)

type Rota struct {
	URI string
	Metodo string
	Funcao func(http.ResponseWriter, *http.Request)
	RequerAuth bool
}

//Congurar colocar todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasAudit

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
		// fmt.Println(rota.Metodo)

	}

	return r
}