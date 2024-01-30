package sqlc

import "github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"

func (c *Category) BatchToDomain(categories []Category) (output []entity.Category) {
	for i := range categories {
		output = append(output, categories[i].ToDomain())
	}
	return
}

func (c *Category) ToDomain() (output entity.Category) {
	output = entity.Category{
		ID:   c.ID,
		Name: c.Name,
	}
	return
}

func (c *Product) BatchToDomain(products []Product) (output []entity.Product) {
	for i := range products {
		output = append(output, products[i].ToDomain())
	}
	return
}

func (p *Product) ToDomain() (output entity.Product) {
	output = entity.Product{
		ID:          p.ID,
		Code:        p.Code.String,
		Name:        p.Name,
		ImageURL:    p.ImageUrl.String,
		Price:       int(p.Price),
		Description: p.Description.String,
		Active:      p.Active.Bool,
		CategoryID:  p.CategoryID.String,
	}
	return
}
