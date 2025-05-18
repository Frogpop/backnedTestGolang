package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

func WaitForPostgres(dsn string, maxRetries int, delay time.Duration) error {
	for i := 0; i < maxRetries; i++ {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			if err = db.Ping(); err == nil {
				return nil // Успешное соединение
			}
		}
		log.Printf("Postgres not ready (attempt %d): %v", i+1, err)
		time.Sleep(delay)
	}
	return fmt.Errorf("Postgres is not available after %d attempts", maxRetries)
}
