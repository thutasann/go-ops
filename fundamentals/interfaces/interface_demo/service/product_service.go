package service

import (
	"context"

	"github.com/thutasann/goops/interfaces/interface_demo/models"
	"github.com/thutasann/goops/interfaces/interface_demo/repository"
)

type ProductService struct {
	repo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(ctx context.Context, id int64) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *ProductService) ListProducts(ctx context.Context) ([]*models.Product, error) {
	return s.repo.FindAll(ctx)
}
