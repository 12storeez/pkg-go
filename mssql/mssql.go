package mssql

import (
	"context"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/zhs/loggr"
)

type MsSql struct {
	Database *sqlx.DB
}

func NewMSSQL(addr, database, user, password, port string) (*MsSql, error) {
	log := loggr.WithContext(context.Background())
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		addr, user, password, port, database)
	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}
	log.Info("connected to mssql")
	return &MsSql{Database: db}, nil
}
