package common

type Post struct {
	Id      int
	Deleted bool

	Title        string
	Text         string
	Url          string
	ThumbnailUrl string

	TimeCreated    int64
	TimeProcessing int64

	AuthorName string
	AuthorUrl  string
}

type Config struct {
	Sources Sources `yaml:"Sources"`
}

type Sources struct {
	RSS []RssSource `yaml:"RSS"`
}

type RssSource struct {
	Url  string `yaml:"Url"`
	Name string `yaml:"Name"`
}
