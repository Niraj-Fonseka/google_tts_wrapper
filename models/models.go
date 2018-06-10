package models

import (
	"fmt"
	"google_tts_wrapper/models/entities"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

//DataStoreClient - global variable for the datastore client
var DataStoreClient *datastore.Client
var gcloudContext context.Context

func InitializeDatastore() {

	gcloudContext = context.Background()
	projectID := os.Getenv("PROJECT_ID")

	// Creates a client.
	client, err := datastore.NewClient(gcloudContext, projectID)

	//Initializing the local client with the Global variable
	DataStoreClient = client

	if err != nil {
		fmt.Println("Error Creating the datastore client")
	}
}

func CreateCredentials(username string, password string) {
	//var creds entities.Credentials

	creds := &entities.Credentials{
		Username: username,
		Password: password,
		Created:  time.Now(),
	}
	key := datastore.IncompleteKey("Task", nil)
	DataStoreClient.Put(gcloudContext, key, creds)
}
