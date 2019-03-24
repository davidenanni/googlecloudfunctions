package main

import (
	"errors"
	"log"
	"encoding/json"
	"fmt"
	"math/rand"
    "time"
    "strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)



// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	var d struct{
	Workload string `json:"numb"`
	}

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	if err := json.Unmarshal([]byte(request.Body), &d); err != nil {
                return events.APIGatewayProxyResponse{}, ErrNameNotProvided
        }


	workload, _ := strconv.ParseInt(d.Workload, 10, 64)

    randomValues := make([]int64, workload)


    var i int64

    for i=0; i<workload; i++{
        randomValues[i] = rand.Int63n(workload)
    }

    workload = workload - 1

    swapped := true;

    rand.Seed(time.Now().UnixNano())
    timestamp := time.Now()


    for swapped==true {
        swapped = false
	
		for i=0; i<workload; i++{
        	prevValue := randomValues[i]
            succValue := randomValues[i+1]

            if prevValue > succValue{
               	temp := randomValues[i+1]
                randomValues[i+1] = randomValues[i]
                randomValues[i] = temp
                swapped = true
            }
        }
    }

    timestampCompleted := time.Now()

    executionTime := timestampCompleted.Sub(timestamp)

    executionTime_ms := float64(executionTime/time.Microsecond)/1000

    //s := fmt.Sprintf("%f",float64(numLanciArea)/float64(numTotLanci))
    exeTime := fmt.Sprintf("%f",executionTime_ms)

    return events.APIGatewayProxyResponse{
        Body: exeTime,
        StatusCode: 200,
    }, nil

}

func main() {
	lambda.Start(Handler)
}