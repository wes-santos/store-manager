package services

import (
	"store-manager/models"
)

type Product struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Quantity uint `json:"quantity"`
}

func GetProducts() ([]Product, error) {
	var products []Product

	rows, err := models.GetProducts()
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var product Product

		err := rows.Scan(
			&product.ID,
			&product.Name,
			&product.Quantity,
		)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}