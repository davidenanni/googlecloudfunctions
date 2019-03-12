// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "context"
        "cloud.google.com/go/pubsub"
        "log"
        "strings"
        "strconv"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
        Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
    msgData  := string(m.Data)
    
    values := strings.Split(msgData,",")
    value0 := values[0]
    value1 := values[1]
    value2 := values[2]
    value3 := values[3]     

    v1, _ := strconv.ParseInt(value1, 10, 64)
    v2, _ := strconv.ParseInt(value2, 10, 64)
    msg := strconv.FormatInt(v1+v2, 10)

    client, _ := pubsub.NewClient(ctx, "tesidaviden")

		topic := client.Topic("response")
  		defer topic.Stop()   


  		r := topic.Publish(ctx, &pubsub.Message{
      		Data: []byte(msg),
  		})
  
    r.Get(ctx)

    return nil
}