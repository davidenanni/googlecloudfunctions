package main

import(
    "context"
	  "cloud.google.com/go/pubsub"
    "os"
    "log"
)


// HelloPubSub consumes a Pub/Sub message.


func publishSimpleMessageTopic( ctx context.Context, client *pubsub.Client, topicName, msg string){
	
	topic := client.Topic(topicName)
  	defer topic.Stop()   


  	r := topic.Publish(ctx, &pubsub.Message{
      	Data: []byte(msg),
  	})
  
  	id, _ := r.Get(ctx)

  	log.Println("[PUBLISHED MESSAGE]   ---   "+id+"   ---   "+msg)
}


func publishMessageTopic( ctx context.Context, client *pubsub.Client, topicName, msg string){
  
  topic := client.Topic(topicName)
    defer topic.Stop()   


    msg = "timestamp,1,2,"+msg
   

    r := topic.Publish(ctx, &pubsub.Message{
        Data: []byte(msg)})
  
    id, _ := r.Get(ctx)

    log.Println("[PUBLISHED MESSAGE]   ---   "+id)
}


func getClient( project string )(context.Context,*pubsub.Client){
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, project)
  	
  	if err != nil {
    	log.Fatalf("Could not create pubsub Client: %v", err)
  	}

  	return ctx,client
}


func main() {

  ctx,client := getClient("tesidaviden")
  
  publishMessageTopic( ctx, client, "exampleTopic", os.Args[1])
}