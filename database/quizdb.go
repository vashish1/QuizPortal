package database

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Quizz contains the data for events.
type Quizz struct {
	Event       string
	Title       string
	Description string
	Answer      string
}

//Insertintoquizdb inserts the data into the database
func Insertintoquizdb(usercollection *mongo.Collection, q Quizz) {

	insertResult, err := usercollection.InsertOne(context.TODO(), q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//Findfromquizdb finds the required data
func Findfromquizdb(collection *mongo.Collection, st string) []Quizz {
	// Pass these options to the Find method
	findOptions := options.Find()
	fmt.Println("st:", st)

	// Here's an array in which you can store the decoded documents
	var results []Quizz

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{"event", st}}, findOptions)
	if err != nil {
		log.Fatal("the error is:", err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Quizz
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal("decoding error:", err)
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal("cursor error", err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Println("questions slice:", results)

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n")
	return results
}

//Deletequiz deletes the corresponding event from database
func Deletequiz(collection *mongo.Collection, st string) {

	filter := bson.D{primitive.E{Key: "", Value: st}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)
}
