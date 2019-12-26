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

//Createdb creates a database
func Createdb() (*mongo.Collection, *mongo.Collection, *mongo.Collection, *mongo.Collection, *mongo.Client) {

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	usercollection := client.Database("Quiz").Collection("User")
	organizercollection := client.Database("Quiz").Collection("organizer")
	eventcollection := client.Database("Quiz").Collection("event")
	quizcollection := client.Database("Quiz").Collection("quiz")
	return usercollection, organizercollection, eventcollection, quizcollection, client
}

//Insertintouserdb inserts the data into the database
func Insertintouserdb(usercollection *mongo.Collection, u User) {

	fmt.Println(u.Username)
	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//Findfromuserdb finds the required data
func Findfromuserdb(usercollection *mongo.Collection, st string) bool {
	filter := bson.D{primitive.E{Key: "username", Value: st}}
	var result User

	err := usercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

//Insertintoorganizerdb inserts the data into the database
func Insertintoorganizerdb(usercollection *mongo.Collection, u Organizer) {

	insertResult, err := usercollection.InsertOne(context.TODO(), u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
}

//Findfromorganizerdb finds the required data
func Findfromorganizerdb(organizercollection *mongo.Collection, st string) bool {
	filter := bson.D{primitive.E{Key: "username", Value: st}}
	var result Organizer

	err := organizercollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return false
	}
    return true

}

//err = client.Disconnect(context.TODO())

// if err != nil {
// 	log.Fatal(err)
// }
// fmt.Println("Connection to MongoDB closed.")

//Finddb finds the required database
func Finddb(c *mongo.Collection, s string) User {
	filter := bson.D{primitive.E{Key: "username", Value: s}}
	var result User

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}

//Findorgdb finds the required database
func Findorgdb(c *mongo.Collection, s string) Organizer {
	filter := bson.D{primitive.E{Key: "username", Value: s}}
	var result Organizer

	err := c.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return result
	}
	return result
}

//Updateorg updates the organizer database
func Updateorg(c *mongo.Collection, o string, s string) {

	filter := bson.D{{"username", o}}

	update := bson.M{
		"$push":bson.M{"events":s}}

	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
