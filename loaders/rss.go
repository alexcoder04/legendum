package loaders

import (
	"github.com/alexcoder04/legendum/common"
	"github.com/mmcdole/gofeed"
)

func Rss(url string, name string) ([]common.Post, error) {
	posts := []common.Post{}

	parser := gofeed.NewParser()

	feed, err := parser.ParseURL(url)
	if err != nil {
		return posts, err
	}

	if name == "" {
		name = feed.Author.Name
	}

	for _, item := range feed.Items {
		posts = append(posts, common.Post{
			Id:    0,
			Title: item.Title,
			Text:  item.Content,
			Url:   item.Link,
			//ThumbnailUrl:   item.Image.URL,
			Deleted:        false,
			TimeCreated:    (*item.PublishedParsed).Unix(),
			TimeProcessing: (*item.PublishedParsed).Unix(),
			AuthorName:     name,
			AuthorUrl:      url,
		})
	}

	return posts, nil
}
