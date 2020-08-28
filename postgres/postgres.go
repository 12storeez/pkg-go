package postgres

import (
	"context"
	"github.com/go-pg/pg"
	"github.com/zhs/loggr"
	"time"
)

func NewPostgres(addr string, database string, user string, password string) *pg.DB {
	log := loggr.WithContext(context.Background())
	opts := &pg.Options{
		Addr:     addr,
		Database: database,
		User:     user,
		Password: password,
	}
	for {
		db := pg.Connect(opts)
		_, err := db.Exec("SELECT 1")
		if err != nil {
			log.Warnf("can't connect to postgres: %v", err)
			log.Info("try again")
			time.Sleep(1 * time.Second)
			continue
		}
		log.Info("connected to postgres")
		return db
	}
}
