package repository

import (
	"context"
	"database/sql"

	sqlc "github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/sqlc/generated"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type ProductRepository struct {
	db *sqlc.Queries
}

func NewProductRepository(db *sqlc.Queries) *ProductRepository {
	return &ProductRepository{db}
}

func (pr *ProductRepository) CreateProduct(ctx context.Context, input entity.Product) (output entity.Product, err error) {
	if err = pr.db.CreateProduct(ctx, sqlc.CreateProductParams{
		ID:          input.ID,
		Code:        sql.NullString{String: input.Code, Valid: true},
		Name:        input.Name,
		ImageUrl:    sql.NullString{String: input.ImageURL, Valid: true},
		Price:       int32(input.Price),
		Description: sql.NullString{String: input.Description, Valid: true},
		Active:      sql.NullBool{Bool: input.Active, Valid: true},
		CategoryID:  sql.NullString{String: input.CategoryID, Valid: true},
	}); err != nil {
		return
	}
	output = input
	return
}

func (pr *ProductRepository) GetProducts(ctx context.Context) (output []entity.Product, err error) {
	products, err := pr.db.GetProducts(ctx, 100)
	if err != nil {
		return
	}
	for i := range products {
		output = append(output, entity.Product{
			ID:           products[i].ID,
			Code:         products[i].Code.String,
			Name:         products[i].Name,
			ImageURL:     products[i].ImageUrl.String,
			Price:        int(products[i].Price),
			Description:  products[i].Description.String,
			Active:       products[i].Active.Bool,
			CategoryID:   products[i].CategoryID.String,
			CategoryName: products[i].CategoryName,
		})
	}
	return
}

func (pr *ProductRepository) GetProductByID(ctx context.Context, productID string) (output entity.Product, err error) {
	product, err := pr.db.GetProductByID(ctx, productID)
	if err != nil {
		return
	}
	output = entity.Product{
		ID:           product.ID,
		Code:         product.Code.String,
		Name:         product.Name,
		ImageURL:     product.ImageUrl.String,
		Price:        int(product.Price),
		Description:  product.Description.String,
		Active:       product.Active.Bool,
		CategoryID:   product.CategoryID.String,
		CategoryName: product.CategoryName,
	}
	return
}
