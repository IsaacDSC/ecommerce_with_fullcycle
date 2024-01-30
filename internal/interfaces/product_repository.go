package interfaces

import (
	"context"

	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type ProductRepository interface {
	CreateProduct(context.Context, entity.Product) (entity.Product, error)
	GetProducts(context.Context) ([]entity.Product, error)
	GetProductByID(ctx context.Context, productID string) (output entity.Product, err error)
}
