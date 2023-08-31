package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/debasishbsws/question-book/internal/db"
)

type Subject struct {
	ID       string   `json:"id"`
	Name     string   `json:"subject_name"`
	Synonyms []string `json:"synonyms"`
}

func GetSubjectsByInstituteId(instituteID string) (*[]Subject, error) {
	const query = `
	SELECT s.id, s.subject_name, s.synonyms
	FROM subject s
	LEFT JOIN institute_subject isub ON s.id = isub.subject_id
	LEFT JOIN institute i ON isub.institute_id = i.id
	WHERE i.id = ?;
	`
	var subjects []Subject
	result, err := db.DbPool.Query(query, instituteID)
	if err != nil {
		errMessage := fmt.Sprintf("No institute with the specified ID: %v", instituteID)
		return &subjects, errors.New(errMessage)
	}
	defer result.Close()

	for result.Next() {
		var subject Subject
		var synonymsJson string
		err := result.Scan(&subject.ID, &subject.Name, &synonymsJson)
		if err != nil {
			log.Println(err)
			return &subjects, err
		}

		subject.Synonyms, err = strJsonArrayToSlice(synonymsJson)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		subjects = append(subjects, subject)
	}

	return &subjects, nil
}
