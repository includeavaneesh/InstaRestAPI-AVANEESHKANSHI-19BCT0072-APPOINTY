package endpoints

// contains POSTS based functions http://localhost:8080/posts

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreatePostEndpoint(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	if r.Method != "POST" || r.URL.Path != "/posts" {
		ErrorHandler(w, r)
		return
	}
	w.Header().Set("content-type", "application/json")

	var post Post

	json.NewDecoder(r.Body).Decode(&post)
	collection := GetPostsClient(Client)
	result, _ := collection.InsertOne(context.TODO(), post)

	collection.FindOne(context.TODO(), bson.D{{"_id", result.InsertedID}}).Decode(&post)
	json.NewEncoder(w).Encode(post)
	defer lock.Unlock()
}

func GetPostEndpoint(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	var post Post
	id := r.URL.Path[7:]
	collection := GetPostsClient(Client)
	postId, _ := primitive.ObjectIDFromHex(id)

	collection.FindOne(context.TODO(), bson.D{{"_id", postId}}).Decode(&post)
	json.NewEncoder(w).Encode(post)
	defer lock.Unlock()
}

func GetPostsByUserEndpoint(w http.ResponseWriter, r *http.Request) {
	lock.Lock()
	var results []Post

	var perPage int64 = 2
	page := 1

	w.Header().Set("content-type", "application/json")
	path := strings.Split(r.URL.Path, "/")
	
	id := path[3]
	
	if len(path) == 5 {
		pageNo, _ := strconv.Atoi(path[4])
		page = pageNo
	}

	findOptions := options.Find()
	findOptions.SetLimit(perPage)
	findOptions.SetSkip(int64((page - 1) * 2))

	collection := GetPostsClient(Client)
	cur, err := collection.Find(context.TODO(), bson.D{{"userID", id}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.TODO()) {

		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	cur.Close(context.TODO())

	json.NewEncoder(w).Encode(results)

	defer lock.Unlock()
}
