package connectdb

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Collection {
	//create client options

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	//connect db
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DataBase is Connected")
	//Create a collection

	collection := client.Database("USER").Collection("USER_DETAILS")

	return collection
}

type ErrorResponce struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

func GetError(err error, w http.ResponseWriter) {

	//log.Fatal(err.Error())

	var res = ErrorResponce{
		StatusCode:   http.StatusInternalServerError,
		ErrorMessage: err.Error(),
	}

	message, _ := json.Marshal(res)
	w.WriteHeader(res.StatusCode)
	w.Write(message)

}
