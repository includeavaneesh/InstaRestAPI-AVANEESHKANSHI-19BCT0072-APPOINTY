package endpoints

import "go.mongodb.org/mongo-driver/bson/primitive"
import "sync"

var (
	lock sync.Mutex
)

type User struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name,omitempty" bson:"name,omitempty"`
	Email    string             `json:"email,omitempty" bson:"email,omitempty"`
	Password string             `json:"password,omitempty" bson:"password,omit"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Caption   string             `json:"caption,omitempty" bson:"caption,omitempty"`
	ImageURL  string             `json:"imageURL,omitempty" bson:"imageURL,omitempty"`
	Time      string             `json:"time,omitempty" bson:"time,omitempty"`
	UserID 	  string			 `json:"userID,omitempty" bson:"userID,omitempty"`
}