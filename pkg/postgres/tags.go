package postgres

import (
	"log"
	"fmt"
	"github.com/demas/cowl-services/pkg/quzx"
)

// represent a PostgreSQL implementation of quzx.TagsService
type TagsService struct {
}

func (s *TagsService) GetTags() ([]*quzx.Tag, error) {

	result := []*quzx.Tag{}
	err := db.Select(&result, "SELECT * FROM Tags")
	return result, err
}

func (s *TagsService) GetTaggedItemsByTagId(tagId int) ([]*quzx.TaggedItem, error) {

	result := []*quzx.TaggedItem{}
	err := db.Select(&result, "SELECT * FROM TaggedItems WHERE TagId = $1", tagId)
	return result, err
}


func (s *TagsService) InsertTaggedItemFromStockItem(questionId int, tagId int) {

	item, err := (&StackService{}).GetStackQuestionById(questionId)
	if err != nil {
		log.Println(questionId)
		log.Fatal(err)
	}

	insertQuestion := `INSERT INTO TaggedItems(TagId, Title, Summary, Content, Link, Date, Source)
			   VALUES ($1, $2, $3, $4, $5, $6, $7)`

	tx := db.MustBegin()
	_, err = tx.Exec(insertQuestion,
		tagId,
		item.Title,
		"",
		"",
		item.Link,
		item.CreationDate,
		1)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func (s *TagsService) InsertTaggedItemFromRss(rssItemId int, tagId int) {

	var item quzx.RssItem
	err := db.Get(&item,
		fmt.Sprintf("SELECT Id, FeedId, Title, Summary, Content, Link, Date FROM RssItem WHERE Id = %d", rssItemId))
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	_, err = tx.Exec("INSERT INTO TaggedItems(TagId, Title, Summary, Content, Link, Date, Source) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)", tagId, item.Title, item.Summary, item.Content, item.Link, item.Date, 2)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}


func (s *TagsService) DeleteTaggedItem(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("DELETE FROM TaggedItems WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
