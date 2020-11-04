package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"simplescraper/data"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PriceHandler handles the prices endpoint
func PriceHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.TODO()

		var cred options.Credential
		// cred.AuthSource = YourAuthSource
		cred.Username = "admin"
		cred.Password = "asd123"
		clientOpt := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred)

		client, e := mongo.Connect(ctx, clientOpt)
		if nil != e {
			log.Println("ERRPR", e)
		}
		defer client.Disconnect(ctx)
		var response []data.Entry
		for i := 0; i < 5; i++ {
			collection := client.Database("scraper-db").Collection("dolarhoy")
			var elem data.Entry
			collection.FindOne(context.TODO(), bson.D{}).Decode(&elem)
			response = append(response, elem)
		}
		fmt.Fprintln(w, response)
	})
}

//InitAPI initializes the API of the project
func InitAPI() {
	router := mux.NewRouter().StrictSlash(true)
	// router.HandleFunc("/lastfive", getData)
	router.Handle("/prices", PriceHandler()).Methods("GET")
	log.Println("Server running")
	log.Fatal(http.ListenAndServe(":8080", router))
}
