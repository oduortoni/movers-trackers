/**
*  package db
*  helpers.go
*  Contains helper functions to check and open database
**/

package db

import (
	"database/sql"
	"errors"
	"path"
	
	_ "github.com/mattn/go-sqlite3"
)

func OpenDatabase() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path.Join("storage/", "vms.db"))
	if err != nil {
		return nil, errors.New("Failed to open database")
	}
	return db, nil
}



