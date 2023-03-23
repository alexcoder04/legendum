package processor

import (
	"log"
	"time"

	"github.com/alexcoder04/legendum/common"
	"github.com/alexcoder04/legendum/loaders"
)

func Daemon(channels []common.Channel) {
	for {
		cont := []common.Post{}
		for i, c := range channels {
			log.Printf("loading channel %d: %s", i+1, c.Url)
			switch c.Type {
			case "rss":
				ps, err := loaders.Rss(c.Url)
				if err != nil {
					log.Printf("failed to load %s", c.Url)
					break
				}
				cont = append(cont, ps...)
			}
		}
		ContentBuffer = cont
		time.Sleep(60 * time.Second)
	}
}
