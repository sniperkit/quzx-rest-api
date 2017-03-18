package postgres

import (
	"github.com/demas/cowl-services/pkg/quzx"
	"log"
)

// represent a PostgreSQL implementation of quzx.BookmarkRepository
type BookmarkRepository struct {
}

func (r *BookmarkRepository) InsertBookmark(bookmark quzx.Bookmark, tags []string) {

	insertBookmarkQuery := `INSERT INTO Bookmark(Url, Title, Description, ReadItLater)
				VALUES($1, $2, $3, $4)
				RETURNING Id`

	insertTagQuery := `INSERT INTO BookmarkTag(Name, BookmarkCount)
			   VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING Id`

	insertTagConnectionQuery := `INSERT INTO BookmarkTagConnnection(BookmarkId, TagId)
				     VALUES($1, $2)`

	tx := db.MustBegin()

	log.Println(bookmark.Title)

	// insert bookmark
	result, err := tx.Exec(insertBookmarkQuery,
			  	bookmark.Url,
		          	bookmark.Title,
			  	bookmark.Description,
				bookmark.ReadItLater)
	if err != nil {
		log.Println(err)

	}

	// and get bookmark_id
	bookmarkId, err := result.LastInsertId()
	if err != nil {
		log.Println("Getting last inserted id for bookmark : " + err.Error())
	}

	// check create tags
	for _, tag := range tags {

		result, err = tx.Exec(insertTagQuery, tag, 0)
		if err != nil {
			log.Println(err)

		}

		tagId, err := result.LastInsertId()
		if err != nil {
			log.Println(err)

		}

		result, err = tx.Exec(insertTagConnectionQuery, bookmarkId, tagId)
	}


	tx.Commit()
}

