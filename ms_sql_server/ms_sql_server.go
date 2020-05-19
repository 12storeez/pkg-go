package ms_sql_server

import (
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
)

type SqlServer struct {
	Database *sqlx.DB
}

func NewMsSQLServerDB(addr, database, user, password, port string) (*SqlServer, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;",
		addr, user, password, port, database)

	db, err := sqlx.Open("sqlserver", connString)
	if err != nil {
		return nil, fmt.Errorf("error creating connection pool: %v", err)
	}

	return &SqlServer{Database: db}, nil
}
