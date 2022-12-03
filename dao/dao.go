package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:42424242@tcp(127.0.0.1:3306)/user")
	if err != nil {
		fmt.Println(err)
		return
	}
	DB = db
}
