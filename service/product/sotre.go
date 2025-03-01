package product

import (
	"database/sql"

	"github.com/yikuanzz/ecom/model"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]model.Product, error) {
	rows, err := s.db.Query(`
	SELECT p.id, p.name, p.description, p.image, p.price, i.quantity, p.createdAt
	FROM products p
	LEFT JOIN inventory i ON p.id = i.product_id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		err := scanRowIntoProduct(rows, &product)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func scanRowIntoProduct(row *sql.Rows, product *model.Product) error {
	return row.Scan(
		&product.ID,
		&product.Name,
		&product.Description,
		&product.Price,
		&product.Image,
		&product.Quantity,
		&product.CreatedAt,
	)
}
