package interfaces

import (
	"context"

	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type CategoryRepository interface {
	CreateCategory(ctx context.Context, input entity.Category) (err error)
	GetAllCategories(ctx context.Context) (output []entity.Category, err error)
	GetCategoryByID(ctx context.Context, categoryID string) (output entity.Category, err error)
}
