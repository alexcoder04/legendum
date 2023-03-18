package common

import "time"

type Post struct {
	TimeCreated    time.Time
	TimeProcessing time.Time
	Title          string
	Text           string
	Url            string
	ThumbnailUrl   string
	Id             string
	Channel        Channel
	Deleted        bool
}

type Channel struct {
	Type string
	Name string
	Url  string
}
