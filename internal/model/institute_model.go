package model

import (
	"encoding/json"
	// "fmt"
	"log"

	"github.com/debasishbsws/question-book/internal/db"
)

type Institute struct {
	Id         string   `json:"id"`
	Name       string   `json:"institute_name"`
	Alt_name   []string `json:"alt_name"`
	WebsiteURL string   `json:"website"`
	Country    string   `json:"country"`
	State      string   `json:"state"`
}

type InstituteWithSubjects struct {
	Institute Institute
	Subjects  []Subject
}

func GetAllInstitutes() (*[]Institute, error) {
	var institutes []Institute
	const query = "SELECT * FROM institute"

	result, err := db.DbPool.Query(query)
	if err != nil {
		log.Fatalln(err)
		return &institutes, err
	}
	defer result.Close()

	for result.Next() {
		var institute Institute
		var altNamesJson string
		err := result.Scan(&institute.Id, &institute.Name, &altNamesJson, &institute.WebsiteURL, &institute.Country, &institute.State)
		if err != nil {
			log.Fatalln(err)
			return &institutes, err
		}

		institute.Alt_name, err = strJsonArrayToSlice(altNamesJson)
		if err != nil {
			log.Fatalln(err)
			return nil, err
		}
		institutes = append(institutes, institute)
	}

	return &institutes, nil
}

func GetInstititeById(instituteID string) (*Institute, error) {
	const query = "SELECT * FROM institute WHERE id = ?"
	var institute Institute
	var altNamesJson string

	err := db.DbPool.QueryRow(query, instituteID).Scan(&institute.Id, &institute.Name, &altNamesJson, &institute.WebsiteURL, &institute.Country, &institute.State)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	institute.Alt_name, err = strJsonArrayToSlice(altNamesJson)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return &institute, nil
}

func strJsonArrayToSlice(strJsonArray string) ([]string, error) {
	var array []string
	err := json.Unmarshal([]byte(strJsonArray), &array)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return array, nil
}

func GetInstituteWithSubjectsByID(instituteID string) (*InstituteWithSubjects, error) {
	var instituteWithSubjects InstituteWithSubjects
	institute, err := GetInstititeById(instituteID)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	subjects, err := GetSubjectsByInstituteId(instituteID)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	instituteWithSubjects.Institute = *institute
	instituteWithSubjects.Subjects = *subjects

	return &instituteWithSubjects, nil
}
