package shared

import (
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/config"
	sqlc "github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/sqlc/generated"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/interfaces"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/repository"
)

type ContainerRepository struct {
	ProductRepository  interfaces.ProductRepository
	CategoryRepository interfaces.CategoryRepository
	OrderRepository    interfaces.OrderRepository
}

func NewContainerRepository() *ContainerRepository {
	conn := config.DbAnaLyticsConn()
	db := sqlc.New(conn)
	return &ContainerRepository{
		ProductRepository:  repository.NewProductRepository(db),
		CategoryRepository: repository.NewCategoryRepository(db),
		OrderRepository:    repository.NewOrderRepository(conn, db),
	}
}
