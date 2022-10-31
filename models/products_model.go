package models

import (
	"database/sql"
	"store-manager/db"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Quantity uint `json:"quantity"`
}

func GetProducts() (*sql.Rows, error) {
	db, err := db.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}

	return rows, nil
}

func GetProductById(ID uint64) (*sql.Rows, error) {
	db, err := db.Connection()

	if err != nil {
		return nil, err
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM products WHERE id = ?", ID)

	if err != nil {
		return nil, err
	}

	return rows, nil
}