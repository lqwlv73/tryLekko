package main

import (
    "os"
    "context"
    "fmt"
    "html"
    "log"
    "net/http"

    "lqwlv73/tryLekko/lekko"
)

func main() {
    env := os.Getenv("ENV")
    client := lekko.NewLekkoClient(context.TODO())

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    })

    http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request){
        fmt.Fprintf(w, client.Example.GetReturnText(env))
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
