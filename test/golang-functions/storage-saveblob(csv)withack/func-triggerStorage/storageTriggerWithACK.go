// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "context"
        "fmt"
        "log"
        "time"

        "cloud.google.com/go/pubsub"
        "cloud.google.com/go/functions/metadata"
)

// GCSEvent is the payload of a GCS event. Please refer to the docs for
// additional information regarding GCS events.
type GCSEvent struct {
        Bucket         string    `json:"bucket"`
        Name           string    `json:"name"`
        Metageneration string    `json:"metageneration"`
        ResourceState  string    `json:"resourceState"`
        TimeCreated    time.Time `json:"timeCreated"`
        Updated        time.Time `json:"updated"`
}

// HelloGCSInfo prints information about a GCS event.
func HelloGCSInfo(ctx context.Context, e GCSEvent) error {
        meta, err := metadata.FromContext(ctx)
        if err != nil {
                return fmt.Errorf("metadata.FromContext: %v", err)
        }
        log.Printf("Event ID: %v\n", meta.EventID)
        log.Printf("Event type: %v\n", meta.EventType)
        log.Printf("Bucket: %v\n", e.Bucket)
        log.Printf("File: %v\n", e.Name)
        log.Printf("Metageneration: %v\n", e.Metageneration)
        log.Printf("Created: %v\n", e.TimeCreated)
        log.Printf("Updated: %v\n", e.Updated)

        timestamp := e.TimeCreated
        updated := e.Updated


        msg := "<"+meta.EventID+">,<"+meta.EventType+">,"+
        	   "<"+e.Bucket+">,<"+e.Name+">,"+
        	   "<"+timestamp.String()+">,<"+updated.String()+">"

    client, _ := pubsub.NewClient(ctx, "tesidaviden")

	topic := client.Topic("ackBucket")
  	defer topic.Stop()   


  	r := topic.Publish(ctx, &pubsub.Message{
      	Data: []byte(msg),
  	})
  
    r.Get(ctx)


    return nil
}