package model

import (
	"log"
	"fmt"
)

type Tag struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Total int `json:"total"`
	Unreaded int `json:"unreaded"`
}

type TaggedItem struct {
	Id int `json:"id"`
	TagId int `json:"tagid"`
	Title string `json:"title"`
	Summary string `json:"summary"`
	Content string `json:"content"`
	Link string `json:"link"`
	Date int64 `json:"date"`
	Source int `json:"source"`  // 1 stack
}

func GetTags() ([]*Tag, error) {

	result := []*Tag{}
	rows, err := db.Query("SELECT Id, Title, Total, Unreaded FROM Tags")

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			t := Tag{}
			rows.Scan(&t.Id, &t.Title, &t.Total, &t.Unreaded)
			result = append(result, &t)
		}
	}

	return result, err
}

func GetTaggedItemsByTagId(tagId int) ([]*TaggedItem, error) {

	result := []*TaggedItem{}
	rows, err := db.Query("SELECT Id, TagId, Title, Summary, Content, Link, Date, Source " +
		"FROM TaggedItems WHERE TagId = $1", tagId)

	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			i := TaggedItem{}
			rows.Scan(&i.Id, &i.TagId, &i.Title, &i.Summary, &i.Content, &i.Link, &i.Date, &i.Source)
			result = append(result, &i)
		}
	}

	return result, err
}


func InsertTaggedItemFromStockItem(questionId int, tagId int) {

	var item StackQuestion
	err := db.Get(&item, fmt.Sprintf("SELECT Title, Link, QuestionId, Tags, CreationDate " +
		"FROM StackQuestions WHERE Id = '%d'", questionId))
	if err != nil {
		log.Println(questionId)
		log.Fatal(err)
	}

	tx := db.MustBegin()
	_, err = tx.Exec("INSERT INTO TaggedItems(TagId, Title, Summary, Content, Link, Date, Source) " +
		"VALUES ($1, $2, $3, $4, $5, $6, $7)", tagId, item.Title, "", "", item.Link, item.CreationDate, 1)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}

func InsertTaggedItemFromRss(rssItemId int, tagId int) {

	var item RssItem
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


func DeleteTaggedItem(id int) {

	tx := db.MustBegin()
	_, err := tx.Exec("DELETE FROM TaggedItems WHERE Id = $1", id)
	if err != nil {
		log.Println(err)
	}
	tx.Commit()
}
