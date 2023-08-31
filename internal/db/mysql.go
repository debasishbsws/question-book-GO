package db

import (
	"database/sql"
	"fmt"
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

	// Create the database connection pool
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s", config.DB_USER_NAME, config.DB_PASSWORD, config.DB_HOST, config.DATABASE_NAME)
	var err error
	DbPool, err = sql.Open("mysql", connectionString)
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
