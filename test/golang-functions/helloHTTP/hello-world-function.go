// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
        "encoding/json"
        "fmt"
        "html"
        "net/http"
)

// HelloHTTP is an HTTP Cloud Function with a request parameter.
func HelloHTTPIT(w http.ResponseWriter, r *http.Request) {
        var d struct {
                Name string `json:"name"`
        }
        if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
                fmt.Fprint(w, "Ciao, Mondo!")
                return
        }
        if d.Name == "" {
                fmt.Fprint(w, "Ciao, Mondo!")
                return
        }
        fmt.Fprintf(w, "Ciao, %s!", html.EscapeString(d.Name))
}
