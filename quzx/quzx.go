package quzx

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type RssFeed struct {
	Id int
	Title string
	Description string
	Link string
	LastSyncTime int64
	ImageUrl string
	AlternativeName string
	Total int
	Unreaded int
	SyncInterval int
	RssType int
	ShowContent int
	ShowOrder int
	Folder string
}

type RssItem struct {
	Id int
	FeedId int
	Title string
	Summary string
	Content string
	Link string
	Date int64
}

type HackerNews struct {
	Id int64
	By string
	Score int
	Time int64
	Title string
	Type string
	Url string
	Readed int
}

type StackTag struct {
	Classification string
	Unreaded int
}

type StackQuestion struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Link string `json:"link"`
	QuestionId int `json:"questionid"`
	Tags string `json:"tags"`
	CreationDate int64 `json:"creationdate"`
}

type TorrentFeed struct {
	Id int `json:"id"`
	TypeId int `json:"type_id"`
	Link string `json:"link"`
	Title string `json:"title"`
	Total int `json:"total"`
	Unread int `json:"unread"`
}
type TorrentNews struct {
	Id int `json:"id"`
	FeedId int `json:"feed_id"`
	Link string `json:"link"`
	Title string `json:"title"`
	Readed int `json:"readed"`
}

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

type FeedService interface {
	GetAllRssFeeds() ([]*RssFeed, error)
	GetUnreadRssFeeds(rssType int) ([]*RssFeed, error)
	GetRssFeedById(id int) (RssFeed, error)
	UpdateRssFeed(feed *RssFeed)
	InsertRssFeed(feed *RssFeed)
	GetRssItemsByFeedId(feed_id int) ([]*RssItem, error)
	SetRssItemAsReaded(id int)
	SetRssFeedAsReaded(feedId int)
	UnsubscribeRssFeed(feedId int)
}
