package db

import (
	"Short/internal/config"
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

type DB struct {
	Options *Options
	DB      *sql.DB
}

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func openDB(o *Options) (*sql.DB, error) {
	connectionString := fmt.Sprintf(""+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"host=%s "+
		"port=%s "+
		"sslmode=%s",
		o.User, o.DBName, o.Password, o.Host, o.Port, o.SslMode)

	sqlDB, err := sql.Open("postgres", connectionString)
	if err != nil {
		sqlDB.Close()

		return nil, err
	}

	// check ping
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return sqlDB, nil
}

func New(config *config.AppConfig) *DB {
	opt := &Options{
		Host:     config.DB.Host,
		Port:     config.DB.Port,
		User:     config.DB.User,
		Password: config.DB.Password,
		DBName:   config.DB.DBName,
		SslMode:  config.DB.SslMode,
	}

	sqlDB, err := openDB(opt)
	if err != nil {
		logrus.Fatalf("error while connect to db: %s", err.Error())

		return nil
	}

	db := &DB{
		Options: opt,
		DB:      sqlDB,
	}

	return db
}
