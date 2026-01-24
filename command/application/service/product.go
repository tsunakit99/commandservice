package service

import (
	"context"

	"github.com/tsunakit99/commandservice/command/domain/models/products"
)

type ProductService interface {
	Add(ctx context.Context, product *products.Product) error
	Update(ctx context.Context, product *products.Product) error
	Delete(ctx context.Context, product *products.Product) error
}
