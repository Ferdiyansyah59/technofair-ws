package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func InitDatabase() *sql.DB {
	cfg := mysql.Config{
		User:   Config.Database.User,
        Passwd: Config.Database.Password,
        Net:    "tcp",
        Addr:   fmt.Sprintf("%s:%d", Config.Database.Host, Config.Database.Port),
        DBName: Config.Database.Dbname,
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error open db connection ", err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("Error ping db ", err)
		panic(err)
	}

	return db
}