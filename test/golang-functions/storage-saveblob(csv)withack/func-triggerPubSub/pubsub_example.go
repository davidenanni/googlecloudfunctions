// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "context"
        _"cloud.google.com/go/pubsub"
        "cloud.google.com/go/storage"
        "log"
        _"os"
        _"io"
        _"strings"
        _"strconv"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
        Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
    msgData  := string(m.Data)  

	client, err := storage.NewClient(ctx)

	if err != nil {
		log.Fatal(err)
	}


	fileName := "demo-filetest-go"

	bucket := client.Bucket("davnan-storage-buck")

	wc := bucket.Object(fileName).NewWriter(ctx)

	wc.ContentType = "text/plain"
	wc.Metadata = map[string]string{
		"x-goog-meta-foo": "foo",
		"x-goog-meta-bar": "bar",
	}

	//client.cleanUp = append(client.cleanUp, fileName)


	wc.Write([]byte(msgData+"\n"))

	wc.Close();

    return nil
}