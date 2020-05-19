package mssql

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type MsSql struct {
	Database *sqlx.DB
}

func NewMSSQL(addr, database, user, password, port string) (*MsSql, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		addr, user, password, port, database)

	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}

	return &MsSql{Database: db}, nil
}
