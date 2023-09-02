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

	/**
	 * This route is used to get all questionPapers of the specified institute and subject.
	 * aditionally it can also filter the results by year, semester and examType.
	 * Example URL : http://<Host>/api/questions/1/1?year=2019&semester=spring&examType=endsem
	 */
	router.HandleFunc("/questions/{instituteId}/{subjectId}", QuestionPapersHandler).Methods("GET")
}
