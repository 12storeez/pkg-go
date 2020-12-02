package postgres

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/zhs/loggr"
)

func NewPostgres(addr, port, database, user, password string) (*pg.DB, error) {
	log := loggr.WithContext(context.Background())
	connectionAddr := fmt.Sprintf("%s:%s", addr, port)
	log.Infof("connect to %s", connectionAddr)
	opts := &pg.Options{
		Addr:     connectionAddr,
		Database: database,
		User:     user,
		Password: password,
	}
	db := pg.Connect(opts)
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Errorf("pg: connection to %s failed\n: %v", connectionAddr, err)
		return nil, err
	}
	return db, nil
}
