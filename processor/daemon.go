package processor

import (
	"log"
	"time"

	"github.com/alexcoder04/legendum/common"
	"github.com/alexcoder04/legendum/loaders"
)

func Daemon(conf common.Config) {
	time.Sleep(time.Second)

	log.Println("starting content loader daemon")

	for {
		// RSS
		for i, c := range conf.Sources.RSS {
			log.Printf("loading rss channel %d: %s", i+1, c.Url)

			posts, err := loaders.Rss(c.Url, c.Name)
			if err != nil {
				log.Printf("failed to load rss '%s'", c.Url)
				continue
			}

			for _, p := range posts {
				_, err := DB.Exec(
					"INSERT INTO posts (title, text, url, thumbnail_url, deleted, time_created, time_processing, author_name, author_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)",
					p.Title, p.Text, p.Url, p.ThumbnailUrl, p.Deleted, p.TimeCreated, p.TimeProcessing, p.AuthorName, p.AuthorUrl,
				)
				if err != nil {
					log.Printf("failed to insert post into database: %s", err.Error())
				}
			}
		}

		time.Sleep(60 * time.Second)
	}
}
