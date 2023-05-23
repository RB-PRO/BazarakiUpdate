package database

import (
	"database/sql"
	"errors"

	_ "github.com/go-sql-driver/mysql"
)

type DB struct {
	*sql.DB
}

func New(filename string) (*DB, error) {

	// Загрузить данные из файла
	DataDB, LoadFile := dataFile(filename)
	if LoadFile != nil {
		return &DB{}, LoadFile
	}

	// Поднять соединение с БД
	dbConn, ErrorOpenDB := sql.Open("mysql", DataDB)
	if ErrorOpenDB != nil {
		return &DB{}, errors.Join(errors.New("Open:"), ErrorOpenDB)
	}

	ErrorPing := dbConn.Ping()
	if ErrorPing != nil {
		return &DB{}, errors.Join(errors.New("Ping:"), ErrorPing)
	}

	return &DB{dbConn}, nil
}

// func (db *DB) Close() error {
// 	return db.Close()
// }
