-- +goose Up
CREATE TABLE "products" (
  "id" bigserial PRIMARY KEY,
  "name" varchar,
  "user_id" int,
  "price" int
);

CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "full_name" varchar,
  "created_at" timestamptz,
  "country_code" int
);

CREATE INDEX ON "products" ("id", "name");

ALTER TABLE "products" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");


-- +goose Down
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;

