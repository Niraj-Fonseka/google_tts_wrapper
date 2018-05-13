package models

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/datastore"
)

//DataStoreClient - global variable for the datastore client
var DataStoreClient *datastore.Client

func InitializeDatastore() {
	ctx := context.Background()
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := datastore.NewClient(ctx, projectID)

	//Initializing the local client with the Global variable
	DataStoreClient = client

	if err != nil {
		fmt.Println("Error Creating the datastore client")
	}
}

func CreateCredentials() {

}
