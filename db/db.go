package db

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mzmico/toolkit/errors"
)

var (
	dbUser *sqlx.DB
)

func init() {

	var (
		err error
	)
	dbUser, err = Load()

	if err != nil {
		panic(err)
	}
}

func Use(name string) *sqlx.DB {

	return dbUser
}
func Load() (*sqlx.DB, error) {

	db, err := sqlx.Connect("mysql", "user:password@localhost/db_user")

	db.SetMaxOpenConns(5)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(3 * time.Hour)

	if err != nil {
		return db, errors.By(err)
	}

	return db, nil
}
