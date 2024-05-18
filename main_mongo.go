package main

import (
    "context"
    "fmt"
    "log"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    // Set client options
    clientOptions := options.Client().ApplyURI("<connection_string>")

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

    // Get a handle for your collection
    collection := client.Database("tazapay-go").Collection("clients")

    // Insert a document
    insertResult, err := collection.InsertOne(context.TODO(), bson.D{
        {Key: "client_name", Value: "Jane Doe"},
        {Key: "contact_details", Value: "jane.doe@example.com"},
        {Key: "client_id", Value: 4},
        {Key: "account_balance", Value: 1000.40},
    })
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Inserted a single document: ", insertResult.InsertedID)

    // Close the connection once no longer needed
    err = client.Disconnect(context.TODO())
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Println("Connection to MongoDB closed.")
    }
}