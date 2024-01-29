-- name: GetProducts :many
SELECT * FROM products;

-- name: CreateProduct :exec
INSERT INTO products
"id"
"code"
"name"
"image_url"
"price"
"description"
"active"
"category_id"
 VALUES($1, $2, $3, $4, $5, $6, $7, $8);