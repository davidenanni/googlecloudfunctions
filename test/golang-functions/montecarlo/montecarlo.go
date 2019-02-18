// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "encoding/json"
        "fmt"
        "net/http"
        "math/rand"
        "math"
        "time"
        "strconv"
)


type Resp struct {
  Result       string
  ExecutionTime string
}

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func Montecarlo(w http.ResponseWriter, r *http.Request) {
        var d struct {
                NumbLanci string `json:"numb"`
        }
        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
        }
        if d.NumbLanci == "" {
                fmt.Fprint(w, "[ERROR]  ---  INVALID INPUT")
                return
        }

        rand.Seed(time.Now().UnixNano())
        
        timestamp := time.Now()

        numTotLanci, _ := strconv.ParseInt(d.NumbLanci, 10, 64)
        numLanciArea := 0
        var i int64

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

}
