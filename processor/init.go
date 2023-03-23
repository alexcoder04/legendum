package processor

import (
	"os"

	"github.com/alexcoder04/legendum/common"
	"gopkg.in/yaml.v3"
)

func loadConfig() []common.Channel {
	data, err := os.ReadFile("./config.yml")
	if err != nil {
		panic("failed to read config")
	}
	conf := common.Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		panic("failed to parse config")
	}
	channels := []common.Channel{}
	for _, item := range conf.Sources.RSS {
		channels = append(channels, common.Channel{
			Type: "rss",
			Url:  item.Url,
		})
	}
	return channels
}

func init() {
	go Daemon(loadConfig())
}
