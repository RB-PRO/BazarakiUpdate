package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(DataDB string) (*DB, error) {

	// Поднять соединение с БД
	dbConn, ErrorOpenDB := sql.Open("mysql", DataDB)
	if ErrorOpenDB != nil {
		return &DB{}, ErrorOpenDB
	}

	ErrorPing := dbConn.Ping()
	if ErrorPing != nil {
		return &DB{}, ErrorPing
	}

	return &DB{dbConn}, nil
}

// func (db *DB) Close() error {
// 	return db.Close()
// }
