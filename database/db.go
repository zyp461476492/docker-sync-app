package database

import (
	"github.com/asdine/storm"
	"github.com/zyp461476492/docker-app/types"
	"log"
)

func GetStorm(config *types.Config) (*storm.DB, error) {
	return storm.Open(config.FileLocation)
}

func CloseStorm(db *storm.DB) {
	err := db.Close()
	if err != nil {
		log.Fatalf("关闭数据库失败 %s", err.Error())
	}
}