package endpoints

// contains USERS based functions http://localhost:8080/users

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func HomePageEndpoint(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		ErrorHandler(w, r)
		return
	}
	w.Write([]byte("Welcome to Instagram"))

}

func CreateUserEndpoint(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method != "POST" || r.URL.Path != "/users" {
		ErrorHandler(w, r)
		return
	}

	w.Header().Set("content-type", "application/json")

	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	collection := GetUserClient(Client)

	newUser.Password = EncryptPassword(newUser.Password)
	result, _ := collection.InsertOne(context.TODO(), newUser)

	collection.FindOne(context.TODO(), bson.D{{"_id", result.InsertedID}}).Decode(&newUser)
	json.NewEncoder(w).Encode(newUser)
	defer lock.Unlock()
}

func GetUserEndpoint(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	id := r.URL.Path[7:]

	collection := GetUserClient(Client)
	userId, _ := primitive.ObjectIDFromHex(id)

	var user User
	collection.FindOne(context.TODO(), bson.D{{"_id", userId}}).Decode(&user)
	json.NewEncoder(w).Encode(user)
	defer lock.Unlock()
}
