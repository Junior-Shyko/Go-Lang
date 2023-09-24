package controllers

import (
	"api/src/response"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "context"
	// "fmt"
	"api/src/repositories"
	"log"
	"api/src/database"
	"encoding/json"
	"api/src/router/models"
	"io/ioutil"
	"net/http"
)

//Cria registro de auditoria
// func CreateAudit(w http.ResponseWriter, r *http.Request) {
func CreateAudit(c *gin.Context) {
	
	bodyRequest, err := ioutil.ReadAll(c.Request.Body)
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
	// ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	
	// result, err := coll.InsertOne(context.Background(), audit)
	repo := repositories.NewAuditRepository(client)
	result, err := repo.Create(audit)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Print(result.InsertedID)
	// c.JSON(http.StatusOK, gin.H{"message" : "Inserido com sucesso " });
	response.JSON(c, http.StatusOK, result )
	
	// w.Write([]byte("Id inserido: %d", usuarioId))
	
}
//Buscando dados de auditoria
func BuscarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando dados de  auditoria"))
}