-- +goose Up
-- +goose StatementBegin

CREATE TYPE delivery_status AS ENUM('calculated', 'new', 'in_progress', 'deliveried', 'failed');

-- todo: добавить статус для интеграции с доставкой
CREATE TABLE delivery (
    uid uuid PRIMARY KEY,
    order_uid uuid,
    from_longitude NUMERIC(10, 8) NOT NULL,
    from_latitude NUMERIC(10, 8) NOT NULL,
    to_longitude NUMERIC(10, 8) NOT NULL,
    to_latitude NUMERIC(10, 8) NOT NULL,
    address VARCHAR(100),
    receiver VARCHAR(100),
    delivery_time interval,
    price INT,
    status delivery_status,
    created_at timestamp,
    updated_at timestamp
);

CREATE INDEX idx_delivery_order_uid ON delivery (order_uid);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX idx_delivery_order_uid;
DROP TABLE delivery;
DROP TYPE delivery_status;

-- +goose StatementEnd
