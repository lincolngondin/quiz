package main

import (
	"encoding/csv"
	"errors"
	"os"
)

type Question struct {
	Statement string
	Answer    string
}

func NewQuestion(statement string, answer string) Question {
	return Question{
		Statement: statement,
		Answer:    answer,
	}
}

func NewQuestionFromRecord(record []string) (*Question, error) {
	if len(record) != 2 {
		return nil, errors.New("Invalid Parameter!")
	}

	return &Question{
		Statement: record[0],
		Answer:    record[1],
	}, nil
}

func GetAllQuestionsFromFile(filename string) ([]Question, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, errors.New("File not found!")
	}
    defer file.Close()
	reader := csv.NewReader(file)
    reader.FieldsPerRecord = 2
    reader.TrimLeadingSpace = true
    records, readErr := reader.ReadAll()
    if readErr != nil || len(records) == 0 {
        return nil, errors.New("Invalid CSV file: Each row in the file MUST have only two fields, the question and the answer!")
    }
    questions := make([]Question, 0, len(records))
    for _, record := range records {
        questions = append(questions, Question{
            Statement: record[0],
            Answer: record[1],
        })
    }
	return questions, nil
}
