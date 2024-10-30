-- +goose Up
-- +goose StatementBegin

CREATE TABLE recipes (
    uid uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    ccal INT NOT NULL,
		cooking_time INTERVAL NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
		img_path TEXT NOT NULl
);

CREATE TABLE recipes_products (
	recipe_uid uuid NOT NULL,
  product_uid uuid NOT NULL,
  PRIMARY KEY (recipe_uid, product_uid)
);

CREATE TABLE recipes_steps (
	recipe_uid uuid PRIMARY KEY,
	step INT NOT NULL,
	description TEXT NOT NULL,
	img_path TEXT NOT NULl
);

CREATE INDEX ix_recipes_products_recipe_uid ON recipes_products (recipe_uid);
CREATE INDEX ix_recipes_products_product_uid ON recipes_products (product_uid);


INSERT INTO recipes (uid, name, ccal, created_at, updated_at, img_path, cooking_time) VALUES
('cee5e047-84de-4d89-996f-d756a7d35c6c', 'Масло блять', 200, NOW()::TIMESTAMP, NOW()::TIMESTAMP, '', '1 hour'::interval);

INSERT INTO recipes_products (recipe_uid, product_uid) VALUES 
('cee5e047-84de-4d89-996f-d756a7d35c6c', '38e27c8f-570c-4d50-bf33-a0bd03afb511');

INSERT INTO recipes_steps (recipe_uid, step, description, img_path) VALUES 
('cee5e047-84de-4d89-996f-d756a7d35c6c', 1, 'Добавь масло', '');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS recipes_products;
DROP TABLE IF EXISTS recipes_steps;

DROP INDEX ix_recipes_products_recipe_uid;
DROP INDEX ix_recipes_products_product_uid;

-- +goose StatementEnd
