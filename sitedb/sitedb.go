package sitedb

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type SiteDB struct {
	Database *sqlx.DB
}

func NewSiteDB(addr string) (*SiteDB, error) {
	db, err := sqlx.Connect("mysql", addr)
	if err != nil {
		return nil, fmt.Errorf("sitedb connection failed: %v", err)
	}
	return &SiteDB{Database: db}, nil
}
