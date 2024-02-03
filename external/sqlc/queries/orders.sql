-- name: RegistryDelivery :exec
INSERT INTO deliveries (
"id",
"CEP",
"address",
"number",
"country",
"district",
"city",
"delivery_price"
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: CreateOrder :exec
INSERT INTO orders (
"id",
"total_price",
"delivery_id",
"status"
) VALUES ($1, $2, $3, $4);

-- name: CreateOrderItems :exec
INSERT INTO order_items (
"id",
"product_id",
"quantity",
"product_price",
"order_id"
) VALUES ($1, $2, $3, $4, $5);