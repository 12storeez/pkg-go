package sitedb

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/zhs/loggr"
)

type SiteDB struct {
	Database *sqlx.DB
}

func NewSiteDB(addr string) (*SiteDB, error) {
	log := loggr.WithContext(context.Background())
	db, err := sqlx.Connect("mysql", addr)
	if err != nil {
		return nil, fmt.Errorf("sitedb connection failed: %v", err)
	}
	log.Info("connected to site db")
	return &SiteDB{Database: db}, nil
}
