package main

import (
	// "fmt"
	// html/template generate HTML output that is safe against code injection
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
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

	// instantiate static fike server
	fs := http.FileServer(http.Dir("assets"))

	mux := http.NewServeMux()

	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// no longer required as running fresh to auto build the server.
	// fmt.Println("runnung on port :1313")

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}
