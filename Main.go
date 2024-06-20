package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/GullitPen", servePen)

    http.HandleFunc("/seven20", serveSeven20)
	http.HandleFunc("/seven2",serveSeven2)

	http.HandleFunc("/seven",serveSeven)
	http.HandleFunc("/seven0",serveSeven0)

	http.HandleFunc("/seven1",serveSeven1)
	http.HandleFunc("/seven10",serveSeven10)

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

func serveSeven2(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven2.html")
}

func serveSeven(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven.html")
}

func serveSeven0(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven0.html")
}

func serveSeven10(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven10.html")
}

func serveSeven1(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./static/seven1.html")
}