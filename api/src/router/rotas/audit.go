package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasAudit = []Rota{
	{
		URI: "/audit",
		Metodo: http.MethodPost,
		Funcao: controllers.CreateAudit,
		RequerAuth: false,
	},
	{
		URI: "/buscar-audit",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarAudit,
		RequerAuth: false,
	},
}