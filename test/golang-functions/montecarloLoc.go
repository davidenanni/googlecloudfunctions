// Package helloworld provides a set of Cloud Function samples.
package main

import (
        _"encoding/json"
        "fmt"
        _"html"
        _"net/http"
        "math"
        "math/rand"
        "time"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func main() {

        rand.Seed(time.Now().UnixNano())
        
        numTotLanci := 100000
        numLanciArea := 0

        for i:=0; i<numTotLanci; i++{

                x := rand.Float64()
                y := rand.Float64()                

                if ( math.Pow(x,float64(2)) + math.Pow(y,float64(2)) <= 1){
                        numLanciArea++
                }

                fmt.Print(i)
                fmt.Print(" --- ")
                fmt.Println(numLanciArea)
        }

        s := fmt.Sprintf("%f",float64(numLanciArea)/float64(numTotLanci))

        fmt.Println(s)

        //fmt.Fprintf(w, "Hello, %s!", html.EscapeString(numLanciArea/numTotLanci))
}