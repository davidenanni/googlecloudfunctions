package main

import (
	"net/http"
	"log"
        "encoding/json"
        "fmt"
        _"html"
        "math/rand"
        "math"
        "time"
        "strconv"
)


type Resp struct {
  ExecutionTime string
}


func Montecarlo(w http.ResponseWriter, req *http.Request) {


        var d struct {
                NumbLanci string `json:"numb"`
        }
        if err := json.NewDecoder(req.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
        }
        if d.NumbLanci == "" {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
        }


        numTotLanci, _ := strconv.ParseInt(d.NumbLanci, 10, 64)
        numLanciArea := 0
        var i int64

        rand.Seed(time.Now().UnixNano())
        timestamp := time.Now()

        for i=0; i<numTotLanci; i++{

                x := rand.Float64()
                y := rand.Float64()                

                if ( math.Pow(x,float64(2)) + math.Pow(y,float64(2)) <= 1){
                        numLanciArea++
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
    http.HandleFunc("/", Montecarlo)
    err := http.ListenAndServeTLS(":443", "server.crt", "server.key", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}