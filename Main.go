package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/GullitPen", servePen)

	fs := http.FileServer(http.Dir("./Static"))
	http.Handle("/", http.StripPrefix("/", fs))

	fmt.Println("Server is running on http://localhost:8080")

	http.ListenAndServe(":8080", nil)
}

// servePen handles requests to /GullitPen by serving seven.html
func servePen(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./Static/seven.html")
}
