package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Enable     bool
	Connection *sql.DB
}

type Shorts struct {
	ID    int64
	Long  string
	Short string
}

func (database *Database) Open(gconfig Config) {
	database.Enable = gconfig.MySQL.Enable

	if !database.Enable {
		return
	}

	udb := gconfig.MySQL
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", udb.DBUser, udb.DBPass, udb.DBHost, udb.DBPort, udb.DBName))

	if err != nil {
		log.Fatalf("%s.\n", err)
	}

	database.Connection = db
	database.Connection.SetConnMaxLifetime(2 * time.Second)
	database.Connection.SetMaxIdleConns(5)
	database.Connection.SetMaxOpenConns(30)

	if err := database.Connection.Ping(); err != nil {
		log.Fatalf("%s.\n", err)
	}

}

func (database *Database) Get(key, value string) (Shorts, error) {
	var shorts Shorts
	err := database.Connection.QueryRow(fmt.Sprintf("SELECT `id`, `long`, `short` FROM `shorts` WHERE `%s` = '%s'", key, value)).Scan(&shorts.ID, &shorts.Long, &shorts.Short)
	return shorts, err
}

func (database *Database) Set(short, long string) error {
	results, err := database.Connection.Query(fmt.Sprintf("INSERT INTO `shorts` (`long`, `short`) VALUES ('%s', '%s')", long, short))
	defer results.Close()
	return err
}

func (database *Database) Del(short string) error {
	results, err := database.Connection.Query(fmt.Sprintf("DELETE FROM `shorts` WHERE `short` = '%s'", short))
	defer results.Close()
	return err
}
