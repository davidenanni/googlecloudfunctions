package main

import(
    "context"
		"cloud.google.com/go/pubsub"
		"fmt"
		_"encoding/json"
    _"os"
    "log"
    _"google.golang.org/api/iterator"
    _"google.golang.org/api/option"
)

type Req struct {
  Value1 string
  Value2 string
}

type Resp struct {
  Result string
}


// HelloPubSub consumes a Pub/Sub message.
func main() {

  //jsonPath := "~/gcloud-keys/tesidaviden-3e3a498fc74c.json"
	  //value1 := strconv.ParseInt(string(m.Value1), 10, 64)
    //value2 := strconv.ParseInt(string(m.Value2), 10, 64)
    
  ctx := context.Background()
  /*proj := os.Getenv("tesidaviden")
  if proj == "" {
    fmt.Fprintf(os.Stderr, "GOOGLE_CLOUD_PROJECT environment variable must be set.\n")
    os.Exit(1)
  }*/
  client, err := pubsub.NewClient(ctx, "tesidaviden")
  if err != nil {
    log.Fatalf("Could not create pubsub Client: %v", err)
  }


  // Sets the name for the new topic.
  topicName := "exampleTopic"

  // Creates the new topic.
  topic, err := client.CreateTopic(ctx, topicName)
  if err != nil {
    log.Fatalf("Failed to create topic: %v", err)
  }

  fmt.Printf("Topic %v created.\n", topic)



    /*value1 := 1
    value2 := 2

    sum := value1 + value2

    s := fmt.Sprintf("%d", sum)

    res := Resp{string(s)}

    js, _ := json.Marshal(res)


    client, _ := pubsub.NewClient(ctx, "tesidaviden")
    t := client.Topic("result")
    result := t.Publish(ctx, &pubsub.Message{
		Data: []byte(js),
	})

	result.Get(ctx)*/
}