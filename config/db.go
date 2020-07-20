package config

import (
	"log"
	"time"
	"database/sql" 
	_ "github.com/go-sql-driver/mysql"
)

// var DB *sql.DB
func InitDB() (db *sql.DB) {

	env := InitEnv()	
	var err error

	connection_string := env.DB_user + ":" + env.DB_password + "@tcp(127.0.0.1:" + env.DB_port + ")/" + env.DB_name
	db, err = sql.Open("mysql", connection_string)
	if err != nil {
		log.Panic(err)
	}

	db.SetConnMaxLifetime(120*time.Second)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}
	