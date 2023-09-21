package rotas

import (
	"api/src/router/controllers"
	"net/http"
)

var rotasAudit = []Rota{
	{
		URI: "/audit",
		Metodo: http.MethodPost,
		Funcao: controllers.CriarAudit,
		RequerAuth: false,
	},
	{
		URI: "/buscar-audit",
		Metodo: http.MethodGet,
		Funcao: controllers.BuscarAudit,
		RequerAuth: false,
	},
}