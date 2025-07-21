package repository

import (
	"context"
	"database/sql"

	"github.com/thutasann/goops/interfaces/interface_demo/models"
)

type MySQLProductRepo struct {
	db *sql.DB
}

func NewMySQLProductRepo(db *sql.DB) *MySQLProductRepo {
	return &MySQLProductRepo{db: db}
}

func (r *MySQLProductRepo) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, price FROM products WHERE id = ?", id)
	p := &models.Product{}
	err := row.Scan(&p.ID, &p.Name, &p.Price)
	return p, err
}

func (r *MySQLProductRepo) FindAll(ctx context.Context) ([]*models.Product, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*models.Product
	for rows.Next() {
		p := &models.Product{}
		rows.Scan(&p.ID, &p.Name, &p.Price)
		products = append(products, p)
	}
	return products, nil
}

func (r *MySQLProductRepo) Save(ctx context.Context, p *models.Product) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO products (name, price) VALUES (?, ?)", p.Name, p.Price)
	return err
}
