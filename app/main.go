package main

import (
	"fmt"
    "log"
    "net/http"
    "encoding/json"
    "github.com/gorilla/mux"
)



func main() {
    r := buildRouter()
    log.Fatal(http.ListenAndServe(":8080", r))
}