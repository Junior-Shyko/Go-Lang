package database


import (
	"context"
	"fmt"
	"log"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
// Replace the placeholder with your Atlas connection string
const uri = "mongodb://localhost:27017"
var collection *mongo.Collection
var ctx = context.TODO()

func ConectionDataBase() (*mongo.Client, error) {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	// serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	// opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	// credential := options.Credential{
	// 	Username: "root",
	// 	Password: "MongoDB2019!",mongodb://localhost:27017
	//  }
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		log.Fatal(err)
	}
	

	return client, nil
	// Send a ping to confirm a successful connection
	// var result bson.M
	// _, err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result)
	// if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
	// 	panic(err)
	// }
	// return result.Collection.c , nil
	// fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	// return
}