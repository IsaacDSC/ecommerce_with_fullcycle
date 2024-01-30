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
}

func NewContainerRepository() *ContainerRepository {
	conn := sqlc.New(config.DbAnaLyticsConn())
	return &ContainerRepository{
		ProductRepository:  repository.NewProductRepository(conn),
		CategoryRepository: repository.NewCategoryRepository(conn),
	}
}
