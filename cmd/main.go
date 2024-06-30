package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/asefatesfay/ecom-go/cmd/api"
	"github.com/asefatesfay/ecom-go/config"
	"github.com/asefatesfay/ecom-go/db"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMySQLDatabase(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal(err)
	}

	initDB(db)

	server := api.NewAPIServer(fmt.Sprintf(":%s", config.Envs.Port), db)
	server.Run()
}

func initDB(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB: successfully connected")
}
