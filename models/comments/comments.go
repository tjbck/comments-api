package comments

import (
	"comments-api/internal/db"

	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Comment struct {
	ID        string `json:"id" bson:"id"`
	ParentID  string `json:"parentId"  bson:"parentId"`
	UserID    string `json:"userId"  bson:"userId"`
	Content   string `json:"content"  bson:"content"`
	CreatedAt int    `json:"createdAt"  bson:"createdAt"`
	EditedAt  int    `json:"editedAt"  bson:"editedAt"`
}

var coll *mongo.Collection = db.GetCollection("comments")

func GetComments() []Comment {
	cursor, err := coll.Find(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var results []Comment
	if err := cursor.All(context.TODO(), &results); err != nil {
		log.Fatal(err)
	}

	return results
}

func GetCommentsById(id string) Comment {
	var result Comment
	err := coll.FindOne(context.TODO(), bson.D{{Key: "id", Value: id}}).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func AddComment(comment Comment) []Comment {
	res, err := coll.InsertOne(context.TODO(), comment)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("inserted document with ID %v\n", res.InsertedID)

	comment_list := GetComments()
	return comment_list
}
