package main

import (
	"net/http"
	"log"
        "encoding/json"
        "fmt"
        _"html"
        "math/rand"
        _"math"
        "time"
        "strconv"
)


type Resp struct {
  ExecutionTime string
}


func Bubblesort(w http.ResponseWriter, req *http.Request) {

        var d struct {
                Workload string `json:"numb"`
        }
        if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
        }
        if d.Workload == "" {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
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

        exeTime := fmt.Sprintf("%f",executionTime_ms)


        profile := Resp{exeTime}

        js, err := json.Marshal(profile)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }


        w.Header().Set("Content-Type", "application/json")
        w.Write(js)


}

func main() {
    http.HandleFunc("/", Bubblesort)
    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}