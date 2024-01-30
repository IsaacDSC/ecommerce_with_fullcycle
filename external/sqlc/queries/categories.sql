-- name: GetAllCategories :many
SELECT * FROM "categories" limit $1;

-- name: CreateCategory :exec
INSERT INTO "categories" ("id", "name", "active") VALUES($1, $2, $3);

-- name: GetCategoryByID :one
SELECT * FROM "categories" WHERE id = $1;