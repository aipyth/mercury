package dbutils

import (
	// "context"
	"os"
	"log"
	"fmt"
	"database/sql"
	_ "github.com/jackc/pgx/v4/stdlib"
)

const db_host = "db"
const db_port = 5432
const db_name = "postgres"
const db_user = "postgres"
const db_password = "postgres"

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

