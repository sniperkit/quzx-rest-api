package model

import "log"
import "github.com/demas/cowl-services/quzx"


func GetStackTags() ([]*quzx.StackTag, error) {

	result := []*quzx.StackTag{}
	rows, err := db.Query("SELECT Classification, Unreaded FROM StackTags WHERE Unreaded > 0")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx.StackTag{}
			rows.Scan(&q.Classification, &q.Unreaded)
			result = append(result, &q)
		}
	}

	return result, err
}

func GetStackQuestionsByClassification(classification string) ([]*quzx.StackQuestion, error) {

	result := []*quzx.StackQuestion{}
	rows, err := db.Query("SELECT Id, Title, Link, QuestionId, Tags, CreationDate FROM StackQuestions " +
		"WHERE Classification = $1 and Readed = 0 ORDER BY CreationDate DESC LIMIT 15", classification)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx.StackQuestion{}
			rows.Scan(&q.Id, &q.Title, &q.Link, &q.QuestionId, &q.Tags, &q.CreationDate)
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


func SetStackQuestionsAsReadedByClassification(classification string) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 WHERE Classification = $1", classification)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 " +
		"WHERE Classification = $1 AND CreationDate < $2", classification, t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
