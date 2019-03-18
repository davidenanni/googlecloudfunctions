package main

import (
    _"io/ioutil"
    "net/http"
    "log"
    _"encoding/json"
    _"bytes"
    "fmt"
    "os"
    "time"
    "strconv"
    "crypto/tls"
)

func main() {

    numTest, _ := strconv.ParseInt(os.Args[1], 10, 64)
    cloudFunc := os.Args[2]


    file, err := os.Create("testPing.csv")
    if err != nil {
        return
    }
    defer file.Close()

    
    

    var i int64

    for i=0; i<numTest; i++{

        latency:= MakeRequest(cloudFunc)        

        file.WriteString(fmt.Sprintf("%f",latency)+"\n")
    }
}


func MakeRequest( cloudFunc string )(float64) {

    timestampRequest := time.Now()


    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    _, err := http.Get(cloudFunc)
    if err != nil {
        log.Fatalln(err)
    }

    timestampResponse := time.Now()
    latency := timestampResponse.Sub(timestampRequest)
    latency_ms := float64(latency/time.Microsecond)/1000

    return latency_ms
}