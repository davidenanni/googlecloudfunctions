package main

import (
    "io/ioutil"
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

    cloudFunc := os.Args[1]
    numTest, _ := strconv.ParseInt(os.Args[2], 10, 64)
    workload := os.Args[3]

    file, err := os.Create("res-"+workload+".csv")
    if err != nil {
        return
    }
    defer file.Close()

    var i int64

    for i=0; i<numTest; i++{

        latency, exeTime := MakeRequest(workload, cloudFunc)        

        file.WriteString(fmt.Sprintf("%f",latency)+","+fmt.Sprintf("%f",exeTime)+"\n")
    }
}


func MakeRequest( workload, cloudFunc string )(float64, float64) {

    message := map[string]interface{}{
        "numb": workload}

    bytesRepresentation, err := json.Marshal(message)
    if err != nil {
        log.Fatalln(err)
    }

    timestampRequest := time.Now()


    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    resp, err := http.Post(cloudFunc, "application/json", bytes.NewBuffer(bytesRepresentation))
    if err != nil {
        log.Fatalln(err)
    }

    timestampResponse := time.Now()
    latency := timestampResponse.Sub(timestampRequest)
    latency_ms := float64(latency/time.Microsecond)/1000

    defer resp.Body.Close()

    exeTime_ms := 0.0


    if resp.StatusCode == http.StatusOK {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        //bodyString := string(bodyBytes)
        //log.Println(bodyString)
        bodyString := string(bodyBytes)
    
        exeTime_ms, err = strconv.ParseFloat(bodyString, 64)
        if err != nil{
            log.Fatalln(err)
        }
    }

    return latency_ms, exeTime_ms
}