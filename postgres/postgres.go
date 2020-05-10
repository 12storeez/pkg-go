package postgres

import (
	"github.com/go-pg/pg"
	"time"
)

func NewPostgres(addr string, database string, user string, password string) *pg.DB {
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
			time.Sleep(1 * time.Second)
			continue
		}
		return db
	}
}
