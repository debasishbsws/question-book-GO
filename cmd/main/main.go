package main

import (
	"log"
	"net/http"

	"github.com/debasishbsws/question-book/config"
	"github.com/debasishbsws/question-book/internal/api"
	"github.com/debasishbsws/question-book/internal/db"
	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	r := mux.NewRouter()
	api.Router(r)
	http.Handle("/", r)

	if err := db.InitializeDatabase(); err != nil {
		log.Println("Error initializing the Database:", err)
	}
	port := ":" + config.PORT
	if err := http.ListenAndServe(port, r); err != nil {
		panic(err)
	}
}
