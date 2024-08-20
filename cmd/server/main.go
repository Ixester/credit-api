package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/usuario/credit-api/internal/credit"
)

func main() {
    r := mux.NewRouter()

    credit.RegisterRoutes(r)

    log.Println("Server is running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", r))
}
