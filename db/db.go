package db

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mzmico/toolkit/errors"
	"github.com/spf13/viper"
)

var (
	dbUser *sqlx.DB

	dbPool = make(map[string]*sqlx.DB)
)

func Use(name string) *sqlx.DB {

	return dbPool[name]
}
func Load() error {

	type DBConfig struct {
		DSN     string `toml:"dsn"`
		MaxOpen int    `toml:"max_open"`
		MaxIdle int    `toml:"max_idle"`
	}

	var (
		config = make(map[string]*DBConfig)
	)

	err := viper.UnmarshalKey("db", &config)

	fmt.Printf("%#v\n", config)

	if err != nil {
		return errors.By(err)
	}

	for name, c := range config {

		fmt.Printf("db %s load\n", name)

		db, err := sqlx.Connect(
			"mysql",
			c.DSN)

		if err != nil {
			return errors.By(err)
		}

		db.SetMaxOpenConns(c.MaxOpen)
		db.SetMaxIdleConns(c.MaxIdle)
		db.SetConnMaxLifetime(3 * time.Hour)

		dbPool[name] = db

	}

	return nil

}
