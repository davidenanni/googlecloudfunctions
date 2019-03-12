// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "context"
        "log"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
        Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSubIT(ctx context.Context, m PubSubMessage) error {
        name := string(m.Data)
        if name == "" {
                name = "Mondo"
        }
        log.Printf("Ciao, %s!", name)
        return nil
}