package main

import (
    "database/sql"
    "fmt"
    "html/template"
    "log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var tmpl = template.Must(template.ParseFiles("templates/index.html"))

func main() {
    mysqlHost := getEnv("MYSQL_HOST", "localhost")
    mysqlUser := getEnv("MYSQL_USER", "default_user")
    mysqlPassword := getEnv("MYSQL_PASSWORD", "default_password")
    mysqlDB := getEnv("MYSQL_DB", "default_db")

    var err error
    db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s",
        mysqlUser,
        mysqlPassword,
        mysqlHost,
        mysqlDB,
    ))
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/", helloHandler)
    http.HandleFunc("/submit", submitHandler)

    log.Println("Server started on :5000")
    log.Fatal(http.ListenAndServe(":5000", nil))
}

func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    rows, err := db.Query("SELECT message FROM messages")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var messages []string
    for rows.Next() {
        var message string
        if err := rows.Scan(&message); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        messages = append(messages, message)
    }

    if err := tmpl.Execute(w, messages); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    newMessage := r.FormValue("new_message")
    _, err := db.Exec("INSERT INTO messages (message) VALUES (?)", newMessage)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
