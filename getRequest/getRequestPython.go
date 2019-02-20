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
        "workload": numLaunch}

    bytesRepresentation, err := json.Marshal(message)
    if err != nil {
        log.Fatalln(err)
    }

    timestampRequest := time.Now()


    http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}

    resp, err := http.Post("https://europe-west1-tesidaviden.cloudfunctions.net/FuncStorage-06", "application/json", bytes.NewBuffer(bytesRepresentation))
    if err != nil {
        log.Fatalln(err)
    }

    timestampResponse := time.Now()
    latency := timestampResponse.Sub(timestampRequest)
    latency_ms := float64(latency/time.Millisecond)


    defer resp.Body.Close()

    exeTime_ms := 0.0

    if resp.StatusCode == http.StatusOK {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        bodyString := string(bodyBytes)
        exeTime_ms, _ = strconv.ParseFloat(bodyString, 64)
    }

    return latency_ms, exeTime_ms

    //var result map[string]interface{}

    //json.NewDecoder(resp.Body).Decode(&result)

    //log.Println(result)
    //log.Println(result["data"])
}