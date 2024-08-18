package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/mattn/go-sqlite3"
    "github.com/TgkCapture/Schedulo/config"
)

var DB *sql.DB

func InitDB() error {
    var err error

    DB, err = sql.Open(config.Cfg.DBDriver, config.Cfg.DBDSN)
    if err != nil {
        return fmt.Errorf("error opening database: %w", err)
    }

    err = DB.Ping()
    if err != nil {
        return fmt.Errorf("error connecting to database: %w", err)
    }

    log.Println("Database connection established")
    return nil
}

func CreateTables() error {
    createScheduleTable := `
    CREATE TABLE IF NOT EXISTS schedules (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        title TEXT NOT NULL,
        time_slot TEXT NOT NULL,
        channel TEXT NOT NULL
    );`

    _, err := DB.Exec(createScheduleTable)
    if err != nil {
        return fmt.Errorf("error creating schedules table: %w", err)
    }

    log.Println("Schedules table created or already exists")
    return nil
}

func CloseDB() {
    if DB != nil {
        err := DB.Close()
        if err != nil {
            log.Printf("Error closing database: %v\n", err)
        } else {
            log.Println("Database connection closed")
        }
    }
}
