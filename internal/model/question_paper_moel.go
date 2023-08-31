package model

import (
	"log"

	"github.com/debasishbsws/question-book/internal/db"
)

type QuestionPaper struct {
	ID          int    `json:"id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Year        string `json:"year"`
	Semester    string `json:"semester"`
	ExamType    string `json:"exam_type"`
	InstituteID string `json:"institute_id"`
	SubjectID   string `json:"subject_id"`
}

func GetQuestionPapersByInstituteIdAndSubjectId(instituteID string, subjectID string, filters map[string]string) (*[]QuestionPaper, error) {
	var query = `
	SELECT *
    FROM question_paper qp
    WHERE qp.institute_id = ? AND qp.subject_id = ?`
	addFilters(&query, filters)

	var questionPapers []QuestionPaper = []QuestionPaper{}
	result, err := db.DbPool.Query(query, instituteID, subjectID)
	if err != nil {
		log.Println(err)
		return &questionPapers, err
	}
	defer result.Close()

	for result.Next() {
		var questionPaper QuestionPaper
		err := result.Scan(&questionPaper.ID, &questionPaper.URL, &questionPaper.Title, &questionPaper.Year, &questionPaper.Semester, &questionPaper.ExamType, &questionPaper.InstituteID, &questionPaper.SubjectID)
		if err != nil {
			log.Println(err)
			return &questionPapers, err
		}
		questionPapers = append(questionPapers, questionPaper)
	}

	return &questionPapers, nil
}

func addFilters(query *string, filters map[string]string) {
	if filters["year"] != "" {
		*query += " AND qp.year = " + filters["year"]
	}
	if filters["semester"] != "" {
		*query += " AND qp.semester = " + filters["semester"]
	}
	if filters["examType"] != "" {
		*query += " AND qp.exam_type = " + filters["examType"]
	}
}
