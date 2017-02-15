package services

import (
	"github.com/ChimeraCoder/anaconda"
	"net/url"
	"log"
	"os"
)

func GetFavoritesTwits(name string) ([]anaconda.Tweet, error) {

	consumer_key := os.Getenv("TWICONKEY")
	consumer_secret := os.Getenv("TWICONSEC")
	access_token := os.Getenv("TWIACCTOK")
	access_token_secret := os.Getenv("TWIACCTOKSEC")

	anaconda.SetConsumerKey(consumer_key)
	anaconda.SetConsumerSecret(consumer_secret)
	api := anaconda.NewTwitterApi(access_token, access_token_secret)

	v := url.Values{}
	v.Set("screen_name", "meditat0r")

	tweets, err := api.GetFavorites(v)
	if err != nil {
		log.Println("Get twitter favorites returned error: %s", err.Error())
	}

	return tweets, err
}