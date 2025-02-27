/**
*  Application specific database setup
*  Create sqlite databases
*  Initialize the tables we need
**/

package db

import (
	"log"
	"os"

	"farmers/files"
)

func CreateDB(db_name string) {
	var db_exists bool
	if db_exists = files.Exists(db_name); db_exists == true {
		log.Println("Database already exists!")
	} else {
		_, err := os.Create(db_name)
		if err != nil {
			return
		}
	}
	if !db_exists {
		CreateGroup()
		CreateFarmers()
		CreateCharges()
	}
	log.Println("Successfully created: " + db_name)

	return
}

func CreateGroup() {
	sql_group := `CREATE TABLE members(
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"membership" TEXT,
		"location" TEXT,
		"produce" TEXT,
		"gid" TEXT);`
	driver, err := OpenDatabase()
	files.CheckError(err)
	defer driver.Close()
	group_statement, err := driver.Prepare(sql_group)
	files.CheckError(err)
	group_statement.Exec()
}

func CreateFarmers() {
	sql_farmers := `CREATE TABLE farmers(
		"id" integer NOT NULL  PRIMARY KEY AUTOINCREMENT,
		"first" TEXT,
		"second" TEXT,
		"gid" TEXT
	);`
	driver, err := OpenDatabase()
	files.CheckError(err)
	defer driver.Close()
	farmers_statement, err := driver.Prepare(sql_farmers)
	files.CheckError(err)
	farmers_statement.Exec()
}

func CreateCharges() {
	sql_charges := `CREATE TABLE charges(
		"means" TEXT,
		"load" integer,
		"cost" integer
	);`
	driver, err := OpenDatabase()
	files.CheckError(err)
	defer driver.Close()
	charges_statement, err := driver.Prepare(sql_charges)
	files.CheckError(err)
	charges_statement.Exec()
}
