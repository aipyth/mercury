package dbutils

import (
	// "context"
	"os"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

var db_host = os.Getenv("DB_HOST")
var db_port = 5432
var db_name = os.Getenv("DB_NAME")
var db_user = os.Getenv("DB_USER")
var db_password = os.Getenv("DB_PASSWORD")

const INLINE_PAGINATION_LIMIT = 2

var DB *sql.DB = nil;

func Connect() *sql.DB {
	db_url := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s", db_user, db_password,
		db_host, db_port, db_name)
	db, err := sql.Open("pgx", db_url)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	DB = db
	return db
}

