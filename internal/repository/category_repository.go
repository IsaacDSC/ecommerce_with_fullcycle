package repository

import (
	"context"
	"database/sql"

	sqlc "github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/sqlc/generated"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type CategoryRepository struct {
	db *sqlc.Queries
}

func NewCategoryRepository(db *sqlc.Queries) *CategoryRepository {
	return &CategoryRepository{db}
}

func (cc *CategoryRepository) CreateCategory(ctx context.Context, input entity.Category) (err error) {
	err = cc.db.CreateCategory(ctx, sqlc.CreateCategoryParams{
		ID:     input.ID,
		Name:   input.Name,
		Active: sql.NullBool{Bool: true, Valid: true},
	})
	return
}

func (cc *CategoryRepository) GetAllCategories(ctx context.Context) (output []entity.Category, err error) {
	categories, err := cc.db.GetAllCategories(ctx, 100)
	if err != nil {
		return
	}
	for i := range categories {
		output = append(output, entity.Category{
			ID:   categories[i].ID,
			Name: categories[i].Name,
		})
	}
	return
}

func (cc *CategoryRepository) GetCategoryByID(ctx context.Context, categoryID string) (output entity.Category, err error) {
	category, err := cc.db.GetCategoryByID(ctx, categoryID)
	if err != nil {
		return
	}
	output = entity.Category{
		ID:   category.ID,
		Name: category.Name,
	}
	return
}
