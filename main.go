package main

import (
	"fmt"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Testing</h1>"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1313"
	}

	mux := http.NewServeMux()

	fmt.Println("runnung on port :1313")

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
