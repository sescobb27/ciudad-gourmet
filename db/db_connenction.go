package db

import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "os"
)

const (
    connection_format string = "user=%s dbname=%s sslmode=disable password=%s host=%s"
    db_name                  = "ciudad_gourmet"
)

var (
    DB                *sql.DB
    user              string
    pass              string
    host              string
    connection_params string
    err               error
)

func init() {
    user = os.Getenv("POSTGRESQL_USER")
    pass = os.Getenv("POSTGRESQL_PASS")
    host = os.Getenv("PGHOST")
    connection_params = fmt.Sprintf(connection_format, user, db_name, pass, host)
    DB, err = sql.Open("postgres", connection_params)
    DB.SetMaxIdleConns(10)
    DB.SetMaxOpenConns(10)
    if err != nil {
        panic(err)
    }
}
