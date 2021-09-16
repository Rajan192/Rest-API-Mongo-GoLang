package handlerfunc

import (
	db "USER/ConnectDb"
	Userstruct "USER/UserStruct"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	//	"os/user"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"golang.org/x/text/collate"
	// u "User/UserStruct"
)

func ValidateInput(user *Userstruct.User, w http.ResponseWriter) {

	var validateError []string
	if user.FirstName == "" {
		validateError = append(validateError, fmt.Errorf("First_Name should not be null ").Error())

	}

	if user.LastName == "" {
		validateError = append(validateError, fmt.Errorf("LastName should not be null ").Error())

	}
	if len(validateError) > 0 {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&validateError)
		return
	}
}



func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var oneUser Userstruct.User
	json.NewDecoder(r.Body).Decode(&oneUser)

	//for validation
	ValidateInput(&oneUser, w)

	collection := db.ConnectDB()
	//oneUser.Active = true

	result, err := collection.InsertOne(context.TODO(), oneUser)

	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}







func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var allUser []Userstruct.User

	collection := db.ConnectDB()
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		db.GetError(err, w)
	}

	defer cur.Close(context.TODO())

	//this will give use daqument one by one
	for cur.Next(context.TODO()) {
		var oneUser Userstruct.User
		err := cur.Decode(&oneUser)
		if err != nil {
			log.Fatal("error in decoding under getUser function ", err)
		}
		if oneUser.Active {
			allUser = append(allUser, oneUser)
		}
		//allUser = append(allUser, oneUser)

	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(allUser)
}







//get user by id

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var user1 Userstruct.User
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}

	filter := bson.M{"_id": id}
	collection := db.ConnectDB()
	err := collection.FindOne(context.TODO(), filter).Decode(&user1)

	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(user1)
}

// func UpdateUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	var params = mux.Vars(r)

// 	id, _ := primitive.ObjectIDFromHex(params["id"])

// 	var user1 Userstruct.User

// 	// Create filter
// 	filter := bson.M{"_id": id}

// 	// Read update model from body request
// 	 json.NewDecoder(r.Body).Decode(&user1)

// 	// prepare update model.
// 	update := bson.D{
// 		{"$set", bson.D{
// 			{"firstname", user1.FirstName},
// 			{"lastname",user1.LastName},
// 			{"age", bson.D{
// 				{"value", user1.Age.Value},
// 				{"interval", user1.Age.Interval},
// 			}},
// 		}},
// 	}
//      //create collection
// 	 collection:=db.ConnectDB()
// 	  err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&user1)

// 	if err != nil {
// 		db.GetError(err, w)
// 		return
// 	}

// 	user1.ID = id

// 	json.NewEncoder(w).Encode(user1)
// }






func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)
	//var user2 Userstruct.User

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])
    if err!=nil{
		log.Fatal("error in ObjectIDFromHex(params) ",err)
	}

	collection := db.ConnectDB()
	filter := bson.M{"_id": bson.M{"$eq": id}}
	update := bson.M{"$set": bson.M{"active": false}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		db.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}
