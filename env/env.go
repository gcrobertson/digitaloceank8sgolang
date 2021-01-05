package env

import (
	"os"
)

var mysqlUser string
var mysqlPassword string

// GetMysqlUser ...
func GetMysqlUser() string {

	if mysqlUser == "" {
		mysqlUser = os.Getenv("MYSQL_ROOT_USER")
	}

	return mysqlUser
}

// GetMysqlPassword ...
func GetMysqlPassword() string {

	if mysqlPassword == "" {
		mysqlPassword = os.Getenv("MYSQL_ROOT_PASSWORD")
	}

	return mysqlPassword
}

// @TODO: Get Email Password , TO: and FROM:
