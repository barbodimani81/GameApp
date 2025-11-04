package mysql

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDB struct {
	db *sql.DB
}

func New() *MySQLDB {
    db, err := sql.Open("mysql", "gameuser:gamepass@tcp(localhost:3308)/game?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci")
	if err != nil {
		panic(fmt.Errorf("can't open mysql db: %v", err))
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

    // Ensure the database is reachable (useful when running via docker-compose)
    var pingErr error
    for i := 0; i < 30; i++ { // ~30s total
        pingErr = db.Ping()
        if pingErr == nil {
            break
        }
        time.Sleep(time.Second)
    }
    if pingErr != nil {
        panic(fmt.Errorf("mysql not reachable: %v", pingErr))
    }

	return &MySQLDB{
		db: db,
	}
}
