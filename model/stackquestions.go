package model

import "log"

type StackQuestion struct {
	Title string `json:"title"`
	Link string `json:"link"`
}

func GetStackQuestionsByClassification(classification string) ([]*StackQuestion, error) {

	result := []*StackQuestion{}
	rows, err := db.Query("SELECT Title, Link FROM StackQuestions WHERE Classification = $1", classification)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := StackQuestion{}
			rows.Scan(&q.Title, &q.Link)
			result = append(result, &q)
		}
	}

	return result, err
}
