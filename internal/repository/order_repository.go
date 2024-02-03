package repository

import (
	"context"
	"database/sql"
	"fmt"

	sqlc "github.com/IsaacDSC/fullcycle_catalog_ecommerce/external/sqlc/generated"
	"github.com/IsaacDSC/fullcycle_catalog_ecommerce/internal/entity"
)

type OrderRepository struct {
	conn *sql.DB
	db   *sqlc.Queries
}

func NewOrderRepository(conn *sql.DB, db *sqlc.Queries) *OrderRepository {
	return &OrderRepository{conn, db}
}

func (or *OrderRepository) GetProducts(ctx context.Context, productsID []string) (output []entity.Product, err error) {
	productIDs := ""
	for i := range productsID {
		productIDs += fmt.Sprintf("%s,", productsID[i])
	}
	query := fmt.Sprintf(`
  SELECT 
  id,
  code,
  name,
  image_url,
  price,
  description,
  active,
  deleted_at
  FROM products WHERE id IN (%s);`, productIDs)
	rows, err := or.conn.Query(query)
	if err != nil {
		return
	}

	var prods []sqlc.Product
	for rows.Next() {
		var p sqlc.Product
		if err = rows.Scan(
			&p.ID,
			&p.Code,
			&p.Name,
			&p.ImageUrl,
			&p.Price,
			&p.Description,
			&p.Active,
			&p.DeletedAt,
		); err != nil {
			return
		}
		prods = append(prods, p)
	}

	for i := range prods {
		output = append(output, entity.Product{
			ID:          prods[i].ID,
			Code:        prods[i].Code.String,
			Name:        prods[i].Name,
			ImageURL:    prods[i].ImageUrl.String,
			Price:       int(prods[i].Price),
			Description: prods[i].Description.String,
			Active:      prods[i].Active.Bool,
			CategoryID:  prods[i].CategoryID.String,
		})
	}
	return
}

func (or *OrderRepository) CreateOrder(ctx context.Context, order entity.Order) (err error) {
	tx, _ := or.conn.Begin()
	defer tx.Rollback()
	qtx := or.db.WithTx(tx)

	if err = qtx.RegistryDelivery(ctx, sqlc.RegistryDeliveryParams{
		ID:            order.DeliveryID,
		CEP:           order.Delivery.CEP,
		Address:       sql.NullString{String: order.Delivery.Address, Valid: true},
		Number:        sql.NullString{String: order.Delivery.Number, Valid: true},
		Country:       sql.NullString{String: order.Delivery.Country, Valid: true},
		District:      sql.NullString{String: order.Delivery.District, Valid: true},
		City:          sql.NullString{String: order.Delivery.Address, Valid: true},
		DeliveryPrice: int32(order.Delivery.DeliveryPrice),
	}); err != nil {
		return
	}

	if err = qtx.CreateOrder(ctx, sqlc.CreateOrderParams{
		ID:         order.ID,
		TotalPrice: order.TotalPrice,
		DeliveryID: order.DeliveryID,
		Status:     order.Status,
	}); err != nil {
		return
	}

	for i := range order.OrderItems {
		if err = qtx.CreateOrderItems(ctx, sqlc.CreateOrderItemsParams{
			ID:           order.OrderItems[i].ID,
			ProductID:    order.OrderItems[i].ProductID,
			Quantity:     order.OrderItems[i].Quantity,
			ProductPrice: order.OrderItems[i].ProductPrice,
			OrderID:      order.ID,
		}); err != nil {
			return
		}
	}

	return tx.Commit()
}
