package main

import (
    "net/http"
    "encoding/json"
    "fmt"
)

type API struct {
    Message string "json:message"
}

func main() {
    http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
        message := API.Message("Hello World!")

        output, err := json.Marsha1(message)

        if err != nil {
            fmt.Println("Something went wrong!")
        }

        fmt.Fprintf(w, string(output))
    })

    http.ListenAndServe(":8080", nil)
}
