package database

import (
	"context"
	"fmt"
	encryption "golang-react-to-do/server/encryption"
	"golang-react-to-do/server/models"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

var mongoClient *mongo.Client

func CreateDbInstance() {
	connectionString := "mongodb+srv://Shreya:Shreya@cluster0.wht2k.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"
	dbName := os.Getenv("DB_NAME")
	collName := os.Getenv("DB_COLLECTION")

	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	mongoClient = client
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to mongodb")
	collection = client.Database(dbName).Collection(collName)
	fmt.Println("collection instace created")
}
func InsertUser(user models.User) models.User {
	collection = mongoClient.Database("golang-db").Collection("users")
	HashedUserPassword, err := encryption.HashPassword(user.Password)
	user.Password = HashedUserPassword
	if err != nil {
		log.Fatal(err)
	}
	insertResult, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("USER CREATED ID", insertResult.InsertedID, user.Username)
	var result models.User
	var nullUser models.User
	err = collection.FindOne(context.Background(), bson.D{{"_id", insertResult.InsertedID}}).Decode(&result)
	fmt.Println("RESULT ", result)
	if err != nil {
		return nullUser
	}
	return result
}
func CheckUserLogin(user models.User) models.User {
	collection = mongoClient.Database("golang-db").Collection("users")
	var result models.User
	var nullUser models.User
	err := collection.FindOne(context.Background(), bson.D{{"username", user.Username}}).Decode(&result)
	if err != nil {
		return nullUser
	}
	var passwordMatch = encryption.CheckPasswordHash(user.Password, result.Password)
	if !passwordMatch {
		return nullUser
	}
	fmt.Println("Login Data", result.ID, result.Username, result.Password)
	return result
}
func InsertTask(task models.Task) models.Task {
	collection = mongoClient.Database("golang-db").Collection("notes")
	insertResult, err := collection.InsertOne(context.Background(), task)
	var result models.Task
	err = collection.FindOne(context.Background(), bson.D{{"_id", insertResult.InsertedID}}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("TASK CREATED", insertResult.InsertedID)
	return result
}
func GetTasksByUser(userIDHex string) []models.Task {
	userID, err := primitive.ObjectIDFromHex(userIDHex)
	collection = mongoClient.Database("golang-db").Collection("notes")
	findResult, err := collection.Find(context.TODO(), bson.M{"user._id": userID})
	var tasksList []models.Task
	if err != nil {
		log.Fatal(err)
		return tasksList
	}
	for findResult.Next(context.TODO()) {
		var task models.Task
		err := findResult.Decode(&task)
		if err != nil {
			log.Fatal(err)
		}
		tasksList = append(tasksList, task)
	}
	if err := findResult.Err(); err != nil {
		log.Fatal(err)
	}
	return tasksList

}
func DeleteTask(taskIdHex string) {
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectId", taskId)
	collection = mongoClient.Database("golang-db").Collection("notes")

	deleteResult, _ := collection.DeleteOne(context.TODO(), bson.M{"_id": taskId})
	if deleteResult.DeletedCount == 0 {
		log.Fatal("Error while deleting", err)
	}
	fmt.Println("Deleted task of Id ", taskIdHex, deleteResult)
}
func TaskStatus(task string) {
	collection = mongoClient.Database("golang-db").Collection("notes")
	id, _ := primitive.ObjectIDFromHex(task)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"status": true}}
	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("modified count: ", result.ModifiedCount)
}
func UpdateTask(taskIdHex string, updatedTask models.Task) {
	collection = mongoClient.Database("golang-db").Collection("tasks")
	taskId, err := primitive.ObjectIDFromHex(taskIdHex)
	if err != nil {
		panic(err)
	}
	fmt.Println("objectID", taskId)
	filter := bson.M{"_id": taskId}
	result, err := collection.ReplaceOne(context.TODO(), filter, updatedTask)
	fmt.Println("full task", updatedTask)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(
		"Note updated successfully", result.MatchedCount, result.ModifiedCount, result.UpsertedCount)
}
