package api

import (
	"github.com/gorilla/mux"
)

func Router(router *mux.Router) {
	// Common routes
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/checkdb", CheckDBHandler)

	// Institute routes

	router.HandleFunc("/institute", GetInstituteHandler).Methods("GET")
	router.HandleFunc("/institute/{instituteId}", InstituteByIdHandler).Methods("GET")
	router.HandleFunc("/{instituteId}/subjects", SubjectsHandler).Methods("GET")

	// Question Papers routes
	// r.HandleFunc("/questionpapers/{courseId}", QuestionPapersHandler)
	// r.HandleFunc("/questionpapers/{courseId}/{questionPaperId}", QuestionPaperByIdHandler)

}
