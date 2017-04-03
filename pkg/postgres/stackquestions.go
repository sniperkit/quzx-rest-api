package postgres

import "log"
import (
	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx.StackService
type StackService struct {
}

func (s *StackService) GetStackTags() ([]*quzx.StackTag, error) {

	selectQuery := `SELECT Classification, Unreaded
	                FROM StackTags
	                WHERE Unreaded > 0 and Hidden = 0`

	result := []*quzx.StackTag{}
	rows, err := db.Query(selectQuery)

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
	selectQuery := `SELECT Title, Link, QuestionId, Tags, CreationDate, Favorite, Classified
			FROM StackQuestions WHERE Id = $1`
	err := db.Get(&item, selectQuery, id)
	return &item, err
}

func (s *StackService) GetSecondTagByClassification(classification string)  (interface{}, error) {

	type Result struct {
		Details string `json:"details"`
		Count int `json:"count"`
	}

	result := []*Result{}
	selectQuery := `SELECT Details, COUNT(Id)
	                FROM StackQuestions
	                WHERE Classification = $1 and READED = 0
	                GROUP BY Details`

	err := db.Select(&result, selectQuery, classification)
	return result, err
}

func (s *StackService) GetStackQuestionsByClassification(classification string) ([]*quzx.StackQuestion, error) {

	result := []*quzx.StackQuestion{}
	selectQuery := `SELECT Id, Title, Link, QuestionId, Tags, CreationDate, Classification, Details,
			       Favorite, Classified
			FROM StackQuestions
			WHERE Classification = $1 and Readed = 0
			ORDER BY CreationDate DESC
			LIMIT 15`

	rows, err := db.Query(selectQuery, classification)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx.StackQuestion{}
			rows.Scan(&q.Id,
				&q.Title,
				&q.Link,
				&q.QuestionId,
				&q.Tags,
				&q.CreationDate,
				&q.Classification,
				&q.Details,
				&q.Favorite,
				&q.Classified)
			result = append(result, &q)
		}
	}

	return result, err
}

func (s *StackService) GetStackQuestionsByClassificationAndDetails(classification string, details string) ([]*quzx.StackQuestion, error) {

	result := []*quzx.StackQuestion{}
	selectQuery := `SELECT Id, Title, Link, QuestionId, Tags, CreationDate, Classification, Details,
			       Favorite, Classified
			FROM StackQuestions
			WHERE Classification = $1 AND Details = $2 AND Readed = 0
			ORDER BY CreationDate DESC
			LIMIT 15`

	rows, err := db.Query(selectQuery, classification, details)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			q := quzx.StackQuestion{}
			rows.Scan(&q.Id,
				&q.Title,
				&q.Link,
				&q.QuestionId,
				&q.Tags,
				&q.CreationDate,
				&q.Classification,
				&q.Details,
				&q.Favorite,
				&q.Classified)
			result = append(result, &q)
		}
	}

	return result, err
}

func (s *StackService) SetStackQuestionAsReaded(question_id int) {

	updateQuery := `UPDATE StackQuestions SET READED = 1 WHERE QuestionId = $1`

	tx := db.MustBegin()
	_, err := tx.Exec(updateQuery, question_id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}


func (s *StackService) SetStackQuestionsAsReadedByClassification(classification string) {

	updateQuery := `UPDATE StackQuestions SET READED = 1 WHERE Classification = $1`

	tx := db.MustBegin()
	_, err := tx.Exec(updateQuery, classification)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *StackService) SetStackQuestionsAsReadedByClassificationFromTime(classification string, t int64) {

	updateQuery := `UPDATE StackQuestions
	                SET READED = 1
	                WHERE Classification = $1 AND CreationDate < $2`

	tx := db.MustBegin()
	_, err := tx.Exec(updateQuery, classification, t)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
