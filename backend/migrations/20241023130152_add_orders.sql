-- +goose Up
-- +goose StatementBegin

CREATE TYPE order_status AS ENUM('created','in_progress','done');

CREATE TABLE orders (
    uid uuid PRIMARY KEY,
    num SERIAL NOT NULL,
    user_uid uuid NOT NULL REFERENCES users(uid),
    status order_status not null default 'created',
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE order_products (
  order_uid uuid NOT NULL,
  product_uid uuid NOT NULL,
  count INT NOT NULL,
  PRIMARY KEY (order_uid, product_uid)
);

CREATE INDEX ix_orders_user_uid ON orders (user_uid);
CREATE INDEX ix_order_products_order_uid ON order_products (order_uid);
CREATE INDEX ix_order_products_product_uid ON order_products (product_uid);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TYPE order_status;

DROP INDEX ix_orders_user_uid;
DROP INDEX ix_order_products_order_uid;
DROP INDEX ix_order_products_product_uid;

DROP TABLE orders;
DROP TABLE order_products;

-- +goose StatementEnd
