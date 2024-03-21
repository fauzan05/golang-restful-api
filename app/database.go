package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDBProduction() *sql.DB {
	db_config := helper.GetDatabaseConfigProd()

	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	db, err := sql.Open(db_config.Type, db_config.Username + "@tcp(" + db_config.Host + ":" + db_config.Port + ")/" + db_config.Name)
	helper.HandleErrorWithPanic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func NewDBTest() *sql.DB {
	db_config := helper.GetDatabaseConfigTest()

	// user:password@tcp(localhost:5555)/dbname?tls=skip-verify&autocommit=true
	db, err := sql.Open(db_config.Type, db_config.Username + "@tcp(" + db_config.Host + ":" + db_config.Port + ")/" + db_config.Name)
	helper.HandleErrorWithPanic(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}