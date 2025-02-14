-- +goose Up
-- +goose StatementBegin

CREATE TABLE recipes (
  uid uuid PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  ccal INT NOT NULL,
  cooking_time INTERVAL NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

CREATE TABLE recipes_products (
	recipe_uid uuid NOT NULL,
  product_uid uuid NOT NULL,
  PRIMARY KEY (recipe_uid, product_uid)
);

CREATE TABLE recipes_steps (
	recipe_uid uuid NOT NULL,
	step INT NOT NULL,
	description TEXT NOT NULL
);

CREATE INDEX ix_recipes_products_recipe_uid ON recipes_products (recipe_uid);
CREATE INDEX ix_recipes_steps_step on recipes_steps (step); 
CREATE INDEX ix_recipes_products_product_uid ON recipes_products (product_uid);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS recipes_products;
DROP TABLE IF EXISTS recipes_steps;

DROP INDEX ix_recipes_products_recipe_uid;
DROP INDEX ix_recipes_products_product_uid;

-- +goose StatementEnd
