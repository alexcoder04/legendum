package processor

import (
	"database/sql"
	"log"
	"os"

	"github.com/alexcoder04/legendum/common"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

var DB *sql.DB

func loadConfig() common.Config {
	data, err := os.ReadFile("./config.yml")
	if err != nil {
		log.Fatalf("failed to read config: %s", err.Error())
	}

	conf := common.Config{}
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("failed to parse config: %s", err.Error())
	}

	return conf
}

func init() {
	var err error
	DB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("failed to open database: %s", err.Error())
	}

	_, err = DB.Exec("CREATE TABLE posts (id INTEGER PRIMARY KEY, title TEXT, text TEXT, url TEXT, thumbnail_url TEXT, deleted INTEGER, time_created TEXT, time_processing TEXT, author_name TEXT, author_url TEXT)")
	if err != nil {
		log.Fatalf("failed to create database table: %s", err.Error())
	}

	go Daemon(loadConfig())
}
