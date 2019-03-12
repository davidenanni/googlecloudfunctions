// Package helloworld provides a set of Cloud Function samples.
package main

import (
        _"encoding/json"
        "fmt"
        _"net/http"
        "math/rand"
        _"math"
        "time"
        "strconv"
        "os"
)


type Resp struct {
  ExecutionTime string
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func Bubblesort(workload int64) {

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

        fmt.Println(executionTime_ms)     
}


func main(){
    workload, _ := strconv.ParseInt(os.Args[1], 10, 64)
    Bubblesort(workload)
}


/*
        rand.Seed(time.Now().UnixNano())
        
        timestamp := time.Now()


        numLanciArea := 0
        

        for i=0; i<numTotLanci; i++{

                x := rand.Float64()
                y := rand.Float64()                

                if ( math.Pow(x,float64(2)) + math.Pow(y,float64(2)) <= 1){
                        numLanciArea++
                }
        }

        timestampCompleted := time.Now()

        executionTime := timestampCompleted.Sub(timestamp)

        executionTime_ms := float64(executionTime/time.Millisecond)

        s := fmt.Sprintf("%f",float64(numLanciArea)/float64(numTotLanci))
        exeTime := fmt.Sprintf("%f",executionTime_ms)


        profile := Resp{s,exeTime}

        js, err := json.Marshal(profile)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }


        w.Header().Set("Content-Type", "application/json")
        w.Write(js)

}*/
