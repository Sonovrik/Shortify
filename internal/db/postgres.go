package db

import (
	"database/sql"
	"fmt"
)

type DB struct {
	options Options
	db      sql.DB
}

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SslMode  string
}

func New(o *Options) *DB {
	// TODO check porting
	connectionString := fmt.Sprintf(""+
		"user=%s "+
		"dbname=%s "+
		"password=%s "+
		"host=%s "+
		"port=%s "+
		"sslmode=%s",
		o.User, o.DBName, o.Password, o.Host, o.Port, o.SslMode)
	sqlDb, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil
	}
	//defer db.Close() // ???

	if err = sqlDb.Ping(); err != nil {
		return nil
	}

	//retDB := DB{
	//	options: o
	//	db: sqlDb
	//}

	//return db
}
