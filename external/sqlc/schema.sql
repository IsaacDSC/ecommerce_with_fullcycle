CREATE TABLE IF NOT EXISTS "categories" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "name" VARCHAR(255) NOT NULL,
  "active" BOOLEAN DEFAULT true,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "products" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "code" VARCHAR(255),
  "name" VARCHAR(255) NOT NULL,
  "image_url" VARCHAR(255),
  "price" INTEGER NOT NULL,
  "description" VARCHAR(255),
  "active" BOOLEAN DEFAULT true,
  "category_id" VARCHAR(255),
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);

CREATE TABLE IF NOT EXISTS "deliveries" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "CEP" VARCHAR(255) NOT NULL,
  "address" VARCHAR(255),
  "number" VARCHAR(255),
  "country" VARCHAR(255),
  "district" VARCHAR(255),
  "city" VARCHAR(255),
  "delivery_price" INTEGER NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "orders" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "total_price" INTEGER NOT NULL,
  "delivery_id" VARCHAR(255) NOT NULL,
  "status" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL,
  FOREIGN KEY (delivery_id) REFERENCES deliveries(id)
);

CREATE TABLE IF NOT EXISTS "order_items" (
  "id" VARCHAR(255) PRIMARY KEY NOT NULL,
  "product_id" VARCHAR(255) NOT NULL,
  "quantity" INTEGER NOT NULL,
  "product_price" INTEGER NOT NULL,
  "order_id" VARCHAR(255) NOT NULL,
  "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  "deleted_at" TIMESTAMPTZ DEFAULT NULL,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_id) REFERENCES products(id)
);