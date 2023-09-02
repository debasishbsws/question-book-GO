package db

import (
	"database/sql"
	"sync"

	"github.com/debasishbsws/question-book/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	DbPool      *sql.DB
	poolMutex   sync.Mutex
	initialized bool
)

func InitializeDatabase() error {
	poolMutex.Lock()
	defer poolMutex.Unlock()

	if initialized {
		return nil
	}

	// Create the database connection pool)
	var err error
	DbPool, err = sql.Open("mysql", config.DATABASE_URI)
	if err != nil {
		return err
	}
	initialized = true
	return nil
}

func TestConnection() error {
	InitializeDatabase()
	err := DbPool.Ping()
	if err != nil {
		return err
	}
	return nil
}
