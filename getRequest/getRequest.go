package main

import (
    _"io/ioutil"
    "net/http"
    "log"
    "encoding/json"
    "bytes"
    "fmt"
    "os"
    "time"
    "strconv"
    "crypto/tls"
)

func main() {

    file, err := os.Create("res.csv")
    if err != nil {
        return
    }
    defer file.Close()

    numLaunch := os.Args[1]
    numTest, _ := strconv.ParseInt(os.Args[2], 10, 64)

    var i int64

    for i=0; i<numTest; i++{

        latency, exeTime := MakeRequest(numLaunch)        

        file.WriteString(fmt.Sprintf("%f",latency)+","+fmt.Sprintf("%f",exeTime)+"\n")
    }
}


func MakeRequest( numLaunch string )(float64, float64) {

    message := map[string]interface{}{
        "numb": numLaunch}

    bytesRepresentation, err := json.Marshal(message)
    if err != nil {
        log.Fatalln(err)
    }

    timestampRequest := time.Now()


    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    resp, err := http.Post("https://34.76.90.247:443", "application/json", bytes.NewBuffer(bytesRepresentation))
    if err != nil {
        log.Fatalln(err)
    }

    timestampResponse := time.Now()
    latency := timestampResponse.Sub(timestampRequest)
    latency_ms := float64(latency/time.Millisecond)

    defer resp.Body.Close()

    var r struct {
        Result       string
        ExecutionTime string
    }


    if  resp.StatusCode == http.StatusOK {

        if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
                log.Println("[ERROR]  ---  INVALID INPUT")
        }
        
    }

    exeTime_ms, _ := strconv.ParseFloat(r.ExecutionTime, 32)

    return latency_ms, exeTime_ms

    //var result map[string]interface{}

    //json.NewDecoder(resp.Body).Decode(&result)

    //log.Println(result)
    //log.Println(result["data"])
}