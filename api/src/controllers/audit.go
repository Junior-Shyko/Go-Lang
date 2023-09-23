package controllers

import (
	// "go.mongodb.org/mongo-driver/bson"
	"context"
	"fmt"
	// "api/src/repositories"
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

	client, err := database.ConectionDataBase()
	if err != nil {
		log.Fatal(err)
	}
	coll := client.Database("admin").Collection("audit")
	result, err := coll.InsertOne(context.TODO(), audit)
	// repo := repositories.NewAuditRepository(coll)
	// err = repo.Create(audit)
	fmt.Print(err)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(result.InsertedID)
	// w.Write([]byte("Id inserido: %d", usuarioId))
	
}
//Buscando dados de auditoria
func BuscarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando dados de  auditoria"))
}