package database

import (
	"context"
	"fmt"
	"log"

	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Event contains the data for events.
type Event struct {
	Eventsname       string
	Eventdescription string
	Startdate string
	Enddate string
	Stime string
	Etime string
	Starttime        time.Time
	Endtime          time.Time
	Timenow          time.Time
}

//Insertintoeventdb inserts the data into the database
func Insertintoeventdb(usercollection *mongo.Collection, e Event) {

	insertResult, err := usercollection.InsertOne(context.TODO(), e)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//Findfromeventdb finds the required data
func Findfromeventdb(collection *mongo.Collection) []Event {
	// Pass these options to the Find method
	findOptions := options.Find()

	// Here's an array in which you can store the decoded documents
	var results []Event

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal("the error is:", err)
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem Event
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

	// fmt.Printf("Found multiple documents (array of pointers): %+v\n")
	return results
}

//Deleteevent deletes the corresponding event from database
func Deleteevent(collection *mongo.Collection, st string) {

	filter := bson.D{primitive.E{Key: "eventsname", Value: st}}
	deleteResult, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the events collection\n", deleteResult.DeletedCount)
}

//Findevent finds the particular event from database
func Findevent(c *mongo.Collection, st string) Event {
	filter := bson.D{primitive.E{Key: "eventsname", Value: st}}
	var result Event

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}
