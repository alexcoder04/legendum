package loaders

import (
	"github.com/alexcoder04/legendum/common"
	"github.com/mmcdole/gofeed"
)

func rssToPosts(rss *gofeed.Feed) []common.Post {
	posts := []common.Post{}
	for _, item := range rss.Items {
		p := common.Post{
			TimeCreated:    *item.PublishedParsed,
			TimeProcessing: *item.PublishedParsed,
			Title:          item.Title,
			Text:           item.Content,
			Url:            item.Link,
			//ThumbnailUrl:   item.Image.URL,
			Channel: common.Channel{
				Type: "rss",
				Name: item.Author.Name,
				Url:  rss.FeedLink,
			},
			Deleted: false,
		}
		p.Id = Hash(p)
		posts = append(posts, p)
	}
	return posts
}

func Rss(url string) ([]common.Post, error) {
	parser := gofeed.NewParser()
	feed, err := parser.ParseURL(url)
	if err != nil {
		return []common.Post{}, err
	}
	return rssToPosts(feed), nil
}
