package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const (
	/*    本地      */
	dbLocalUser     = ``
	dbLocalPassword = ``
	dbLocalName     = ``
	dbLocalHost     = ``
	/*    線上      */
	dbRemoteUser     = ``
	dbRemotePassword = ``
	dbRemoteName     = ``
	dbRemoteHost     = ``
)

func (conn DBConnInfo) DbStartUp() (*gorm.DB, error) {
	connectString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true", conn.USERNAME, conn.USERPASSWORD, conn.DBHOST, conn.DBNAME)
	return gorm.Open("mysql", connectString)
}

func DbConn() (db *gorm.DB, err error) {
	return DBConnInfo{USERNAME: dbLocalUser, USERPASSWORD: dbLocalPassword, DBHOST: dbLocalHost, DBNAME: dbLocalName}.DbStartUp()
}
