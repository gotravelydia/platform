// Copyright 2016 Travelydia, Inc. All rights reserved.

package database

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gotravelydia/platform/config"
	"github.com/gotravelydia/platform/log"

	"github.com/ziutek/mymysql/godrv"
)

type Database struct {
	DB *sql.DB
}

func New() (*Database, error) {
	db, err := Open()
	if err != nil {
		return nil, err
	}

	// Ping to verify connection.
	err = db.DB.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Open() (*Database, error) {
	config, err := config.ServiceConfig.GetRDSConfig()
	if err != nil {
		return nil, err
	}

	db, err := ConnectDB(config)
	if err != nil {
		return nil, err
	}

	return &Database{
		DB: db,
	}, nil
}

func ConnectDB(config *config.RDSConfig) (*sql.DB, error) {
	// Override utf8.
	godrv.Register("SET NAMES utf8mb4")

	// Set timezone to UTC.
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal("Failed to set timezone.")
		return nil, err
	}
	godrv.SetLocation(loc)

	conn := fmt.Sprintf("tcp:%s:%d,timeout=5s*%s/%s/%s", config.Host, config.Port, config.Database, config.Username, config.Password)
	db, err := sql.Open("mymysql", conn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	db.SetMaxIdleConns(config.PoolSize)

	return db, nil
}

func (db *Database) QueryRows(query string, args ...interface{}) (*sql.Rows, error) {
	row, err := db.DB.Query(query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return row, nil
}

func (db *Database) QuerySingleRow(query string, args ...interface{}) *sql.Row {
	row := db.DB.QueryRow(query, args...)
	return row
}

func (db *Database) Execute(query string, args ...interface{}) (sql.Result, error) {
	res, err := db.DB.Exec(query, args...)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	return res, nil
}

func (db *Database) Close() error {
	err := db.DB.Close()
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
