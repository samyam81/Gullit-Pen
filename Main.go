package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/GullitPen", servePen)
    http.HandleFunc("/seven20", serveSeven20)

    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", http.StripPrefix("/", fs))

    fmt.Println("Server is running on http://localhost:8080")
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        fmt.Println("Error starting server:", err)
    }
}

func servePen(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven.html")
}

func serveSeven20(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven20.html")
}
