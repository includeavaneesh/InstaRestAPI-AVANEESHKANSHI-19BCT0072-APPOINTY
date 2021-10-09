package main

import (
	"net/http"

	"./endpoints"
)

/*	GENERAL FUNCTIONS

	These functions are generally used throughout the API and are not application specific

	1. ErrorHandler - to handle response errors
	2. encryptPassword - encrypts password by employing SHA256
	3. getUserClient - provides a connection for user collection
	4. getPostsClient - provides a connection for posts collection
*/

/*	API ENDPOINTS

	These are the API's endpoints

	A. User Handlers
		> CreateUser - creates a new user
		> GetUser - returns details of a user

	B. Post Handlers
		> CreatePost - creates a new post
		> GetPost - returns details of a post
		> GetPostsByUser - returns all the posts of a user
*/

func main() {
	endpoints.InitMongo()
	mux := http.NewServeMux()
	mux.HandleFunc("/", endpoints.HomePageEndpoint)
	mux.HandleFunc("/users", endpoints.CreateUserEndpoint)
	mux.HandleFunc("/users/", endpoints.GetUserEndpoint)
	mux.HandleFunc("/posts", endpoints.CreatePostEndpoint)
	mux.HandleFunc("/posts/", endpoints.GetPostEndpoint)
	mux.HandleFunc("/posts/users/", endpoints.GetPostsByUserEndpoint)

	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		panic(err)
	}
}
