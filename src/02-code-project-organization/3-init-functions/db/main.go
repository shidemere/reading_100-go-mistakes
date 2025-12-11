package main

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB
// затрудняют тестирование
func init() {
	dataSourceName := os.Getenv("MYSQL_DATA_SOURCE_NAME")
	d, err := sql.Open("mysql", dataSourceName)
	// единственый способ обработать ошибку
	if err != nil {
		log.Panic(err)
	}
	err = d.Ping()
	if err != nil {
		log.Panic(err)
	}
	// использовать глобальные переменные не круто потому что их можно изменить откуда угодно
	db = d
}

func createClient(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// Единственный случай пока что мне известный для init Блока это конфигурация "статичского поля" 