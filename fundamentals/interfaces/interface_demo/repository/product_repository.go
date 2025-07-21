package repository

import (
	"context"

	"github.com/thutasann/goops/interfaces/interface_demo/models"
)

type ProductRepository interface {
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	FindAll(ctx context.Context) ([]*models.Product, error)
	Save(ctx context.Context, product *models.Product)
}
