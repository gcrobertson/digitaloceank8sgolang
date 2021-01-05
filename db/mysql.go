package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // blank because Go needs a 3rd party SQL Driver but then uses the internal package database/sql through interfacing...
	"gitlab.com/gregcrobertson/millicent/env"
)

var db *sql.DB

// OpenMySQL Connection ...
func OpenMySQL() error {

	if db == nil {

		dbName := "app"
		dbUser := env.GetMysqlUser()
		dbPassword := env.GetMysqlPassword()
		// mysql.connect("db-service")                                     [Inside-the-same-namespace]
		// mysql.connect("<service-name>.<namespace>.svc.cluster.local")   [Connect-to-another-namespace]
		s := fmt.Sprintf("%s:%s@tcp(mysql-service:3306)/%s", dbUser, dbPassword, dbName)
		var err error
		db, err = sql.Open("mysql", s)

		if err != nil {
			return err
		}
	}

	return nil
}

// CloseMySQL Connection ...
func CloseMySQL() {

	// defer the close till after the main function has finished executing
	db.Close()
}
