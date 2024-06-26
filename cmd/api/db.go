package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func (app *Application) connectToDB() *sql.DB {
	db, err := sql.Open("mysql", "u245175828_contract:SoftSwiftLock25852.@tcp(153.92.15.23)/u245175828_test")

	fmt.Println("Connecting . . . .")
	app.errHandlerNoti(err, "Can't Connect to database")
	return db
}
