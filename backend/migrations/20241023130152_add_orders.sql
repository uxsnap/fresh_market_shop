-- +goose Up
-- +goose StatementBegin

CREATE TABLE orders (
    uid uuid PRIMARY KEY,
    num SERIAL NOT NULL,
    sum INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE orders_products (
  order_uid uuid NOT NULL,
  product_uid uuid NOT NULL,
  PRIMARY KEY (order_uid, product_uid)
);

CREATE INDEX ix_orders_products_order_uid ON orders_products (order_uid);
CREATE INDEX ix_orders_products_product_uid ON orders_products (product_uid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX ix_orders_products_order_uid;
DROP INDEX ix_orders_products_product_uid;

DROP TABLE orders;
DROP TABLE orders_products;

-- +goose StatementEnd
