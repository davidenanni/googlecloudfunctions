package main

import (
	"encoding/csv"
	"context"
	"log"
	"cloud.google.com/go/storage"
	"os"
	"bufio"

)

func main() {

	projectID := "tesidaviden"
	bucket := os.Args[1]

	ctx := context.Background()

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	createBucket( projectID, bucket )
}

// [END storage_quickstart]