package main

import (
	"database/sql"
	"log"

	"github.com/TonyCross23/goecom/cmd/api"
	"github.com/TonyCross23/goecom/config"
	"github.com/TonyCross23/goecom/database"
	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := database.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPassword,
		Addr:                 config.Envs.DBAddress,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	InitStorage(db)

	server := api.NewApiServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func InitStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}
	log.Println("Connected to the database successfully")
}
