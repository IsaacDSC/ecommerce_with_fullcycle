package interfaces

import (
	"context"

	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type OrderRepository interface {
	GetProducts(ctx context.Context, productsID []string) (output []entity.Product, err error)
	CreateOrder(ctx context.Context, order entity.Order) (err error)
}
