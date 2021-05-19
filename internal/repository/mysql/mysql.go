package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(
	host string,
	port int,
	user string,
	password string,
	dbname string,
) (db *sql.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s", user, password, host, port, dbname)
	db, err = sql.Open("mysql", dsn)
	return
}
