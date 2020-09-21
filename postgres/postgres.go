package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
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
			fmt.Printf("pg: connection to %s failed\n", addr)
			fmt.Println("try retry...")
			time.Sleep(1 * time.Second)
			continue
		}
		log.Info("connected to postgres")
		return db
	}
}
