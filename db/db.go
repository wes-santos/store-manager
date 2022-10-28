package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // MYSQL Driver Connection
)

// Connections opens connection with DB
func Connection() (*sql.DB, error) {
	connectionURL := "root:root@/StoreManager?charset=utf8&parseTime=True&loc=Local"

	db, err := sql.Open("mysql", connectionURL)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}