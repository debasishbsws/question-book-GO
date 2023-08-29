package api

import (
	"encoding/json"
	// "log"
	"net/http"

	"github.com/debasishbsws/question-book/internal/db"
	"github.com/debasishbsws/question-book/internal/model"
	"github.com/gorilla/mux"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World from Home"))
}

func CheckDBHandler(w http.ResponseWriter, r *http.Request) {
	if err := db.TestConnection(); err == nil {
		w.Write([]byte("Database connection test passed."))
	} else {
		http.Error(w, "Database connection test failed.", http.StatusInternalServerError)
	}
}

func GetInstituteHandler(w http.ResponseWriter, r *http.Request) {

	institutes, err := model.GetAllInstitutes()
	if err != nil {
		http.Error(w, "Error fetching institutes.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(institutes)

}

func InstituteByIdHandler(w http.ResponseWriter, r *http.Request) {

	instituteID := mux.Vars(r)["instituteId"]

	institute, err := model.GetInstituteWithSubjectsByID(instituteID)
	if err != nil {
		http.Error(w, "Error fetching institute.", http.StatusInternalServerError)
		return
	}
	if institute == nil {
		http.Error(w, "No institute found with the specified ID name: "+instituteID, http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(institute)
}

func SubjectsHandler(w http.ResponseWriter, r *http.Request) {

	instituteID := mux.Vars(r)["instituteId"]
	courses, err := model.GetSubjectsByInstituteId(instituteID)
	if err != nil {
		http.Error(w, "Error fetching courses.", http.StatusInternalServerError)
		return
	}
	if courses == nil {
		http.Error(w, "No courses found with the specified Institute ID.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courses)

}

// // Add your Question Papers handler similarly
