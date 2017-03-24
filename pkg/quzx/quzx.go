package quzx

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type RssFeed struct {
	Id int `db:"id"`
	Title string `db:"title"`
	Description string `db:"description"`
	Link string `db:"link"`
	UpdateUrl string `db:"updateurl"`
	ImageTitle string `db:"imagetitle"`
	ImageUrl string `db:"imageurl"`
	ImageHeight int `db:"imageheight"`
	ImageWidth int `db:"imagewidth"`
	LastSyncTime int64 `db:"lastsynctime"`
	Total int `db:"total"`
	Unreaded int `db:"unreaded"`
	SyncInterval int `db:"syncinterval"`
	AlternativeName string `db:"alternativename"`
	RssType int `db:"rsstype"`
	ShowContent int `db:"showcontent"`
	ShowOrder int `db:"showorder"`
	Folder string `db:"folder"`
	LimitFull int `db:"limitfull"`
	LimitHeadersOnly int `db:"limitheadersonly"`
	Broken int `db:"broken"`
}

// Returns ORDER BY clause to get RssItems
func (rssFeed *RssFeed) OrderByClause() (string) {

	if rssFeed.ShowOrder == 0 {
		return " ORDER BY Date DESC"
	}

	 return " ORDER BY Date ASC"
}

func (rssFeed *RssFeed) Limit() (int) {

	if rssFeed.ShowContent == 1 {
		return rssFeed.LimitFull
	}

	return rssFeed.LimitHeadersOnly
}

type RssItem struct {
	Id int
	FeedId int
	Title string
	Summary string
	Content string
	Link string
	Date int64
	ItemId string
	Readed int
	Favorite int
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
	Favorite int
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
	Classification string `json:"classification"`
	Favorite int `json:"favorite"`
	Classified int `json:"classified"`
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

type BookmarkTag struct {
	Id int `json:"id"`
	Name string `json:"name"`
	BookmarkCount int `json:"cnt"`
}

type Bookmark struct {
	Id int `json:"id"`
	Url string `json:"url"`
	Title string `json:"title"`
	Description string `json:"description"`
	ReadItLater int `json:"readItLater"`
}

type BookmarkPOST struct {
	*Bookmark
	Tags []string `json:"tags"`
}

type BookmarkTagConnection struct {
	Id int `json:"id"`
	BookmarkId int `json:"bookmarkId"`
	TagId int `json:"tagId"`
}