package controllers

import (
	"api/src/repositories"
	"log"
	"api/src/database"
	"encoding/json"
	"api/src/router/models"
	"io/ioutil"
	"net/http"
)

//Cria registro de auditoria
func CreateAudit(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var audit models.Audit
	if err = json.Unmarshal(bodyRequest, &audit); err != nil {
		log.Fatal(err)
	}

	db, err := database.ConectionDataBase()
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.NewAuditRepository(db)
	message := repositories.Create(audit)
	w.Write([]byte(message))
	
}
//Buscando dados de auditoria
func BuscarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando dados de  auditoria"))
}