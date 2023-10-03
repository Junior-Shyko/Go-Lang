package controllers

import (
	"api/src/response"
	"github.com/gin-gonic/gin"
	"api/src/repositories"
	"log"
	"api/src/database"
	"encoding/json"
	"api/src/router/models"
	"io/ioutil"
	"net/http"
)

//Cria registro de auditoria
func CreateAudit(c *gin.Context) {
	//corpo da requisição
	bodyRequest, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("error response: %v", err)
		log.Fatal(err)
	}
	var audit models.Audit//instancia com os valores recebidos

	if err = json.Unmarshal(bodyRequest, &audit); err != nil {
		log.Printf("error instancia: %v", err)
		log.Fatal(err)
	}
	//Conectando com banco
	client, err := database.ConectionDataBase()
	if err != nil {
		log.Fatal(err)
	}
	repo := repositories.NewAuditRepository(client)
	result, err := repo.Create(audit)
	if err != nil {
		log.Fatal(err)
	}
	//Retornando mensagem para API
	response.JSON(c, http.StatusOK, result )
}

//Buscando dados de auditoria
func BuscarAudit(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando dados de  auditoria"))
}