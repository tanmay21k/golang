package controller

import (
	"context"
	"fmt"
	"log"

	models "github.com/tanmaykulkarni2112/go/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

const connectionString = ""
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

// connecting with database
func init() {
	// client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.Background(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Success in connection")
	collection = client.Database(dbName).Collection(colName)
}

// mongo Helpers are placed in another folder namely Helpers
func insertOneMovie(movie models.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted the movie with id", inserted.InsertedID)
}

// updating one record
func updateOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId) // string to objectId conversion
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Modified count: ", result.ModifiedCount)
}

// deleting one record
func deleteOneRecord(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(deleteCount)
	}
	fmt.Println("Movie got deleted with delete count", deleteCount, "and the id was", id)
}

// Delete all records from mongoDB
func deleteAll() {
	filter := bson.D{{}}
	deleteCount, err := collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of deleted records ", deleteCount)
}
