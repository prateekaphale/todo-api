package todo

import (
    "encoding/json"
    "net/http"

    "github.com/gocql/gocql"
    "github.com/prateekaphale/todo-api/config"
)

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
    var t Todo
    json.NewDecoder(r.Body).Decode(&t)
    t.ID = gocql.TimeUUID()
    err := config.Session.Query("INSERT INTO todos (id, title, description, completed) VALUES (?, ?, ?, ?)",
        t.ID, t.Title, t.Description, t.Completed).Exec()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(t)
}
