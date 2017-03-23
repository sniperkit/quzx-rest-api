package postgres

import "log"
import (
	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx.StackService
type StackService struct {
}

func (s *StackService) GetStackTags() ([]*quzx.StackTag, error) {

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

func (s *StackService) GetStackQuestionById(id int) (*quzx.StackQuestion, error) {

	var item quzx.StackQuestion

	selectQuery := `SELECT Title, Link, QuestionId, Tags, CreationDate FROM StackQuestions WHERE Id = $1`
	err := db.Get(&item, selectQuery, id)
	return &item, err
}

func (s *StackService) GetStackQuestionsByClassification(classification string) ([]*quzx.StackQuestion, error) {

	result := []*quzx.StackQuestion{}
	rows, err := db.Query("SELECT Id, Title, Link, QuestionId, Tags, CreationDate, Classification FROM StackQuestions " +
		"WHERE Classification = $1 and Readed = 0 ORDER BY CreationDate DESC LIMIT 15", classification)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx.StackQuestion{}
			rows.Scan(&q.Id, &q.Title, &q.Link, &q.QuestionId, &q.Tags, &q.CreationDate, &q.Classification)
			result = append(result, &q)
		}
	}

	return result, err
}


func (s *StackService) SetStackQuestionAsReaded(question_id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 WHERE QuestionId = $1", question_id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}


func (s *StackService) SetStackQuestionsAsReadedByClassification(classification string) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 WHERE Classification = $1", classification)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *StackService) SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	tx := db.MustBegin()
	_, err := tx.Exec("UPDATE StackQuestions SET READED = 1 " +
		"WHERE Classification = $1 AND CreationDate < $2", classification, t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
