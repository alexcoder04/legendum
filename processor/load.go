package processor

import (
	"log"

	"github.com/alexcoder04/legendum/common"
)

func Load(startTime int) []common.Post {
	posts := []common.Post{}

	rows, err := DB.Query("SELECT id, title, text, url, thumbnail_url, deleted, time_created, time_processing, author_name, author_url FROM posts WHERE time_processing > ? ORDER BY time_processing ASC LIMIT 10", startTime)
	if err != nil {
		log.Printf("failed to query posts: %s", err.Error())
		return posts
	}
	defer rows.Close()

	for rows.Next() {
		post := common.Post{}
		err := rows.Scan(&post.Id, &post.Title, &post.Text, &post.Url, &post.ThumbnailUrl, &post.Deleted, &post.TimeCreated, &post.TimeProcessing, &post.AuthorName, &post.AuthorUrl)
		if err != nil {
			continue
		}
		posts = append(posts, post)
	}

	return posts
}
