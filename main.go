package main

import (
	"fmt"

	"github.com/rokoga/filas-backend/web"
)

func main() {

	done := make(chan string)
	go web.Run(done)

	fmt.Print(<-done)

	// Testes com conexao mongo
	// client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	// err = client.Connect(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // defer client.Disconnect(ctx)

	// collection := client.Database("bancoteste").Collection("cteste")

	// ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()

	// store := domain.Store{
	// 	ID:      "123112122",
	// 	Name:    "Lojinhaaaaaaaa",
	// 	URLName: "lojinhaaaaaaaa",
	// 	Queue:   nil,
	// }

	// result, err2 := collection.InsertOne(ctx, store)
	// if err2 != nil {
	// 	fmt.Printf("ERRO %v\n", err2)
	// 	// return nil
	// }
	// var resultfind domain.Store

	// fmt.Printf("RESULT %v\n", result)
	// fmt.Printf("RESULT %v\n", result.InsertedID)

	// // filter := bson.D{{Key: "name", Value: "Lojinhaaaaaaaa"}}
	// filter := bson.D{{Key: "_id", Value: result.InsertedID}}

	// err4 := collection.FindOne(ctx, filter).Decode(&resultfind)
	// if err4 != nil {
	// 	log.Fatal(err4)
	// }
	// fmt.Printf("FIND %v\n", resultfind)

	// cur, err := collection.Find(context.Background(), filter)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer cur.Close(context.Background())
	// for cur.Next(context.Background()) {
	// 	// To decode into a struct, use cursor.Decode()
	// 	result := struct {
	// 		Foo string
	// 		Bar int32
	// 	}{}
	// 	err := cur.Decode(&result)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	// do something with result...

	// 	// To get the raw bson bytes use cursor.Current
	// 	raw := cur.Current
	// 	// do something with raw...
	// 	fmt.Printf("RAW %v\n", raw)
	// }
	// if err := cur.Err(); err != nil {
	// 	fmt.Printf("ERRO %v\n", err)
	// }

}
