package repositories

import (
	"context"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"api/src/router/models"
	"log"
	"fmt"
	// "encoding/json"
	// "net/http"
)
//AuditRepository representa o repositorio de auditoria
type AuditRepository struct {
	database *mongo.Database
}

type Audit struct {
	UserId 			uint64 `json:"user_id"`
	ObjectId 		string `json:"object_id"`
	ObjectType		string `json:"object_type"`
	CreateTimeStamp	time.Time `json:"create_timestamp"`
	Action			string `json:"action"`
	Message			string `json:"message"`
	Key				string `json:"key"`
	Value			string `json:"value"`
}	

//NewAuditRepository cria um repo de audit
func NewAuditRepository(database *mongo.Database) *AuditRepository {
	
	return &AuditRepository{database}
}
//Create insere dados no banco de dados
func (repo AuditRepository) Create(audit models.Audit) (*interface{}, error) {
	fmt.Println(repo.database)
	//setando a collection
	coll := repo.database.Collection("audit")
	//populando dados para inserir
	aud := Audit{
		UserId: audit.UserId,
		ObjectId: audit.ObjectId,
		ObjectType: audit.ObjectType,
		CreateTimeStamp: time.Now(),
		Action: audit.Action,
		Message: audit.Message,
		Key: audit.Key,
		Value: audit.Value,
	}
	//metodo para inserir
	result, err := coll.InsertOne(context.TODO(), aud)

	if err != nil {
		log.Fatal(err)
	}
	//retornando ID
	return &result.InsertedID, nil
}