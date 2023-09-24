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
func (repo AuditRepository) Create(audit models.Audit) (error) {
	fmt.Println(repo.database)
	coll := repo.database.Collection("audit")
	// coll := client.Database("db").Collection("books")
	doc := Audit{
		UserId: audit.UserId,
		ObjectId: audit.ObjectId,
		ObjectType: audit.ObjectType,
		CreateTimeStamp: time.Now(),
		Action: audit.Action,
		Message: audit.Message,
		Key: audit.Key,
		Value: audit.Value,
	}
	// doc := []audit.UserId
	result, err := coll.InsertOne(context.TODO(), doc)

	if err != nil {
		log.Fatal(err)
	}
	// for _, id := range result.InsertedIDs {
	// 	fmt.Printf("Inserted document with _id: %v\n", id)
	// }
	newID := result.InsertedID
	fmt.Printf("Inserted document with _id: %v\n", newID)
	return nil
}