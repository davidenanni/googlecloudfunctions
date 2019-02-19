package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/storage"
)


func createBucket( projectID string, client *storage.Client, bucketName string, bucketClass string, zone string){

	ctx := context.Background()


	// Creates a Bucket instance.
	bucket := client.Bucket(bucketName)

	// Creates the new bucket.
	err := bucket.Create(ctx, projectID, &storage.BucketAttrs{
		StorageClass: bucketClass,
		Location:     zone,
	})

	if err != nil {
		log.Fatalf("Failed to create bucket: %v", err)
	}

	fmt.Printf("[CREATED][BUCKET]   ---	  %s \n", bucketName)
}