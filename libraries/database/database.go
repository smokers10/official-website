package database

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	config "github.com/smokers10/official-website.git/libraries"
)

type databaseConnetion struct{}

func Databases() *databaseConnetion {
	return &databaseConnetion{}
}

func (dbc *databaseConnetion) MYSQLConnection() (*sql.DB, error) {
	config := config.ReadConfig().Database

	db, err := sql.Open("mysql", config.URI)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxConnection)
	db.SetMaxIdleConns(config.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(config.MaxConnectionLifeTime * int(time.Minute)))
	db.SetConnMaxIdleTime(time.Duration(config.MaxIdleConnectionLifeTime * int(time.Minute)))

	return db, nil
}
