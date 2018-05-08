package main

import (
	"net/http"
	"os"

	"github.com/adal-io/htmltopdf/pkg/server"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/convert", server.Convert).Methods("POST")
	h := cors.Default().Handler(r)
	http.ListenAndServe(":"+os.Getenv("HTML_TO_PDF_PORT"), h)
}
