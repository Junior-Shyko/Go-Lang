package controllers

import (
	"net/http"
)

//Cria registro de auditoria
func CriarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Inserindo auditoria"))
}
//Buscando dados de auditoria
func BuscarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando dados de  auditoria"))
}