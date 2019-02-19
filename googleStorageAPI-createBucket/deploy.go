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


	csvFile, _ := os.Open("./bucket/"+bucket+".csv")
    reader := csv.NewReader(bufio.NewReader(csvFile))

    bucketName, _ := reader.Read()
    class, _ := reader.Read()
    zone, _ := reader.Read()    

	createBucket( projectID, client, bucketName[1], class[1], zone[1] )
}

// [END storage_quickstart]