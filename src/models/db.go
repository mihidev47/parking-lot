package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // MySQL Database driver
)

func Conn() (db *sql.DB) {

	dbDriver := "mysql"   // Database driver
	dbUser := "root"      // Mysql username
	dbPass := "" // Mysql password
	dbName := "parking"   // Mysql schema

	// Realize the connection with mysql driver
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	// If error stop the application
	if err != nil {
		panic(err.Error())
	}

	// Return db object to be used by other functions
	return db
}