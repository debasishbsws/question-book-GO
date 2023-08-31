package api

import (
	"encoding/json"
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
		http.Error(w, err.Error(), http.StatusNotFound)
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

// Question Paper Handlers

func QuestionPapersHandler(w http.ResponseWriter, r *http.Request) {
	instituteID := mux.Vars(r)["instituteId"]
	subjectID := mux.Vars(r)["subjectId"]

	var filters map[string]string = make(map[string]string)

	filters["year"] = r.URL.Query().Get("year")
	filters["semester"] = r.URL.Query().Get("semester")
	filters["examType"] = r.URL.Query().Get("examType")

	questionPapers, err := model.GetQuestionPapersByInstituteIdAndSubjectId(instituteID, subjectID, filters)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if questionPapers == nil {
		http.Error(w, "No question papers found with the specified Institute ID and Subject ID.", http.StatusNotFound)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questionPapers)
}
