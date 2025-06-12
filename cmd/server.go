package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

type student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}

func main() {
	// insert()
	// find()
	remove()
}

func connect() (*mongo.Database, error) {
	clientOptions := options.Client()
	clientOptions.ApplyURI("mongodb://root:secret@0.0.0.0:27017/?authSource=admin")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("connected to mongo_db")

	return client.Database("belajar_golang"), nil
}

// func insert() {
// 	db, err := connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	_, err = db.Collection("student").InsertOne(ctx, student{"Hadzami", 21})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	_, err = db.Collection("student").InsertOne(ctx, student{"Monkey D Luffy", 19})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	fmt.Println("insert success!")
// }

// func find() {
// 	db, err := connect()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	csr, err := db.get("")
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	defer csr.Close(ctx)

// 	result := make([]student, 0)
// 	for csr.Next(ctx) {
// 		var row student
// 		err := csr.Decode(&row)
// 		if err != nil {
// 			fmt.Println(err.Error())
// 			return
// 		}

// 		result = append(result, row)
// 	}

// 	if len(result) > 0 {
// 		fmt.Println("Name: ", result[0].Name)
// 		fmt.Println("Grade: ", result[0].Grade)
// 	}
// }

func remove() {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var selector = bson.M{"name": "Hadzami"}
	_, err = db.Collection("student").DeleteOne(ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println("successfully deleted")
}
