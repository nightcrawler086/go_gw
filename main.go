package main

import (
	"fmt"
	"log"
	"net/http"
  "context"
  "encoding/json"

  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func getByTitle(w http.ResponseWriter, r *http.Request) {
	uri := "mongodb://localhost:27017"
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()
	coll := client.Database("db").Collection("app")
  title := r.URL.Query().Get("title")
  fmt.Println("Getting Movie with Title: {}", title)
  if title != "" {
    var result bson.M
    err = coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
    if err == mongo.ErrNoDocuments {
      fmt.Printf("No document was found with the title %s\n", title)
      return
    }
    if err != nil {
      panic(err)
    }
    jsonData, err := json.MarshalIndent(result, "", "    ")
    if err != nil {
      panic(err)
    }
    fmt.Printf("%s\n", jsonData)
    json.NewEncoder(w).Encode(result)

  }
}

func main() {

	//title := "Back to the Future"

	http.HandleFunc("/getByTitle", getByTitle) // Register the handler function for the root route
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server on port 8080
}

