package model

import (
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
	"github.com/google/uuid"
)

func (c Category) FromDomain(category entity.Category) (output *Category) {
	output = &Category{
		ID:   category.ID,
		Name: category.Name,
	}
	return
}

func (c NewCategory) BatchToDomain(categories []NewCategory) (output []entity.Category) {
	for i := range categories {
		output = append(output, categories[i].ToDomain())
	}
	return
}

func (c NewCategory) ToDomain() (output entity.Category) {
	output = entity.Category{
		ID:   uuid.NewString(),
		Name: c.Name,
	}
	return
}

func (c Product) FromDomain(product entity.Product) (output *Product) {
	output = &Product{
		ID:          product.ID,
		Code:        &product.Code,
		Name:        product.Name,
		ImageURL:    product.ImageURL,
		Price:       product.Price,
		Description: &product.Description,
		Active:      product.Active,
		CategoryID:  product.CategoryID,
		Category: &Category{
			ID:   product.CategoryID,
			Name: product.CategoryName,
		},
	}
	return
}

func (c NewProduct) BatchToDomain(products []NewProduct) (output []entity.Product) {
	for i := range products {
		output = append(output, products[i].ToDomain())
	}
	return
}

func (p NewProduct) ToDomain() (output entity.Product) {
	var (
		code, desc string
	)
	if p.Code != nil {
		code = *p.Code
	}
	if p.Description != nil {
		desc = *p.Description
	}
	output = entity.Product{
		ID:          uuid.NewString(),
		Code:        code,
		Name:        p.Name,
		ImageURL:    p.ImageURL,
		Price:       p.Price,
		Description: desc,
		Active:      p.Active,
		CategoryID:  p.CategoryID,
	}
	return
}

func (p NewOrder) ToDomain(products []entity.Product) (output *entity.Order) {
	deliveryID := uuid.NewString()
	orderID := uuid.NewString()
	orderItems := make([]entity.OrderItem, len(products))

	for i := range products {
		orderItems[i] = entity.OrderItem{
			ID:           uuid.NewString(),
			ProductID:    products[i].ID,
			Quantity:     int32(p.Items[i].Quantity),
			ProductPrice: int32(products[i].Price),
			OrderID:      orderID,
			DeletedAt:    products[i].DeletedAt,
		}
	}

	output = &entity.Order{
		ID:         orderID,
		Status:     entity.ORDER_STATUS_CREATED,
		OrderItems: orderItems,
		DeliveryID: deliveryID,
		Delivery: entity.Delivery{
			ID:       deliveryID,
			CEP:      p.Delivery.Cep,
			Address:  p.Delivery.Address,
			Number:   p.Delivery.Number,
			Country:  p.Delivery.Country,
			District: p.Delivery.District,
			City:     p.Delivery.City,
		},
	}

	return
}
