-- +goose Up
-- +goose StatementBegin

CREATE TABLE payment_cards (
    uid uuid PRIMARY KEY,
    user_uid uuid,
    external_uid uuid,
    number varchar(4),
    expired varchar(5)
);

CREATE TABLE payments (
    uid uuid PRIMARY KEY,
    user_uid uuid,
    order_uid uuid,
    card_uid uuid,
    sum int,
    currency varchar(4),
    time timestamp
);



-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE payments;
DROP TABLE payment_cards;
-- +goose StatementEnd
