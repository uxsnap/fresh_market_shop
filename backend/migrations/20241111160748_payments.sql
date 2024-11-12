-- +goose Up
-- +goose StatementBegin

CREATE TABLE users_cards (
    uid uuid PRIMARY KEY,
    user_uid uuid,
    external_uid uuid,
    number varchar(4),
    expired varchar(5)
);

--- remove after mvp
CREATE TABLE full_cards (
    uid uuid PRIMARY KEY,
    user_uid uuid,
    number varchar(16),
    expired varchar(5),
    cvv varchar(3)
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
DROP TABLE test_card;
DROP TABLE users_cards;
-- +goose StatementEnd
