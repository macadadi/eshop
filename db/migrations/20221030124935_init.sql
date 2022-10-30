-- +goose Up
CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "price" bigint NOT NULL
);

-- +goose Down
DROP TABLE IF EXISTS products

