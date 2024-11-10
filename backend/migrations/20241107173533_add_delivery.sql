-- +goose Up
-- +goose StatementBegin

-- todo: добавить статус для интеграции с доставкой
CREATE TABLE delivery (
    uid uuid PRIMARY KEY,
    order_uid uuid,
    from_longitude NUMERIC(10, 8) NOT NULL,
    from_latitude NUMERIC(10, 8) NOT NULL,
    to_longitude NUMERIC(10, 8) NOT NULL,
    to_latitude NUMERIC(10, 8) NOT NULL,
    address VARCHAR(50),
    receiver VARCHAR(50),
    delivery_time interval,
    price INT,
    created_at timestamp,
    updated_at timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE delivery IF EXISTS;
-- +goose StatementEnd
