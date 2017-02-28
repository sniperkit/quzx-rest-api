package model

import "log"

type StackQuestion struct {
	Title string `json:"title"`
	Link string `json:"link"`
	QuestionId int `json:"questionid"`
	Tags string `json:"tags"`
	CreationDate int64 `json:"creationdate"`
}

type StackTag struct {
	Classification string
	Unreaded int
}

func GetStackTags() ([]*StackTag, error) {

	result := []*StackTag{}
	rows, err := db.Query("SELECT Classification, Unreaded FROM StackTags WHERE Unreaded > 0")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := StackTag{}
			rows.Scan(&q.Classification, &q.Unreaded)
			result = append(result, &q)
		}
	}

	return result, err
}

func GetStackQuestionsByClassification(classification string) ([]*StackQuestion, error) {

	result := []*StackQuestion{}
	rows, err := db.Query("SELECT Title, Link, QuestionId, Tags, CreationDate FROM StackQuestions WHERE Classification = $1 and Readed = 0 ORDER BY CreationDate DESC LIMIT 15", classification)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := StackQuestion{}
			rows.Scan(&q.Title, &q.Link, &q.QuestionId, &q.Tags, &q.CreationDate)
			result = append(result, &q)
		}
	}

	return result, err
}


func SetStackQuestionAsReaded(question_id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 WHERE QuestionId = $1", question_id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
