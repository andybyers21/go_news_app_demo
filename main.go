package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Testing</h1>"))
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	// NOTE: I don"t think this is required by godotenv but will leave it here for the timebeing for troubleshooting."
	//	if port == "" {
	//		port = "1313"
	//	}

	mux := http.NewServeMux()

	fmt.Println("runnung on port :1313")

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
