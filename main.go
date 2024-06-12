package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/prateekaphale/todo-api/config"
    "github.com/prateekaphale/todo-api/internal/todo"
)

func main() {
    config.InitScylla()
    defer config.Session.Close()

    r := mux.NewRouter()

    r.HandleFunc("/todos", todo.CreateTodoHandler).Methods("POST")
    // Add routes for other handlers (Read, Update, Delete, List)

    log.Fatal(http.ListenAndServe(":8080", r))
}
