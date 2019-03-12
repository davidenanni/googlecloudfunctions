package main

import (
		"io"
        "encoding/json"
        "fmt"
        "html"
        "net/http"
        "math/rand"
        "math"
        "time"
        "strconv"
)


func montecarlo(w http.ResponseWriter, r *http.Request) {
	
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

        s := fmt.Sprintf("%f",float64(numLanciArea)/float64(numTotLanci))

        fmt.Fprintf(w, "%s", html.EscapeString(s))
}

var mux map[string]func(http.ResponseWriter, *http.Request)

func main() {
	server := http.Server{
		Addr:    ":8000",
		Handler: &myHandler{},
	}

	mux = make(map[string]func(http.ResponseWriter, *http.Request))
	mux["/"] = montecarlo

	server.ListenAndServe()
}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}

	io.WriteString(w, "My server: "+r.URL.String())
}