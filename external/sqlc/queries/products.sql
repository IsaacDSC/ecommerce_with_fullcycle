-- name: GetProducts :many
select products.*, categories.name as category_name from products join categories on products.category_id = categories.id limit $1;

-- name: CreateProduct :exec
INSERT INTO products ("id","code","name","image_url","price","description","active","category_id") VALUES($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetProductByID :one
SELECT products.*, categories.name as category_name FROM products join categories on products.category_id = categories.id WHERE products.id = $1;