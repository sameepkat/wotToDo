package db

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"os/user"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func getDBFilePath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", errors.New("error getting user: " + err.Error())
	}

	dbFilePath := filepath.Join(usr.HomeDir, ".config", "wottodo", "tasksDB.db")

	err = os.MkdirAll(filepath.Dir(dbFilePath), os.ModePerm)
	if err != nil {
		return "", errors.New("error creating directory: " + err.Error())
	}

	return dbFilePath, nil
}

func initDB(dbFilePath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFilePath)

	if err != nil {
		return nil, fmt.Errorf("error opening database: %w", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS Tasks(
	id INTEGER PRIMARY KEY,
	 Title TEXT(30),
	 Status TEXT CHECK(Status IN ('TODO', 'DONE')),
	 createdAt DATETIME DEFAULT CURRENT_TIMESTAMP)`)
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("error creating Tasks Table: %w", err)
	}
	return db, nil
}

func Exec() (*sql.DB, error) {
	dbFilePath, err := getDBFilePath()
	if err != nil {
		return nil, fmt.Errorf("error getting DB File Path: %w", err)
	}
	return initDB(dbFilePath)
}

func Add(db *sql.DB, title string, status string) (sql.Result, error) {
	query := `INSERT INTO Tasks (Title, Status) VALUES (?, ?)`

	res, err := db.Exec(query, title, status)
	if err != nil {
		return nil, fmt.Errorf("failed to add task: %w", err)
	}

	return res, nil
}
