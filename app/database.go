package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDBProduction() *sql.DB {
	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api")
	helper.HandleErrorWithPanic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func NewDBTest() *sql.DB {
	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_restful_api_test")
	helper.HandleErrorWithPanic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}