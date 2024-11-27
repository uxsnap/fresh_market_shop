-- +goose Up
-- +goose StatementBegin

-- Темы обращений, справочник
CREATE TYPE support_tickets_topics (
    uid uuid PRIMARY KEY,
    name varchar,
    description varchar
);

INSERT INTO support_tickets_topics (uid, name, description) 
VALUES ('99c0203e-a6da-45d6-98f3-021ac86adff9', 'Другое', 'Все что не подходит под другие темы обращений');

INSERT INTO support_tickets_topics (uid, name, description) 
VALUES 
(gen_random_uuid(), 'Авторизация/регистрация', 'Проблемы с регистрацией/авторизацией'),
(gen_random_uuid(), 'Заказы', 'Проблемы с заказами'),
(gen_random_uuid(), 'Оплата', 'Проблемы с оплатой'),
(gen_random_uuid(), 'Доставка', 'Проблемы с доставкой'),
(gen_random_uuid(), 'Рецепты', 'Проблемы с рецептами. Неточности, проблемы при просмотре рецептов'),
(gen_random_uuid(), 'Продукты', 'Проблемы с продуктами. Неточности описания, проблемы отображения'),
(gen_random_uuid(), 'Рекомендации', 'Проблемы с рекомендациями'),
(gen_random_uuid(), 'Адреса доставки', 'Проблемы с адресами доставки');


CREATE TYPE support_ticket_status AS ENUM('created', 'in_process', 'solved', 'cant_solve');

CREATE TABLE support_tickets (
    uid uuid PRIMARY KEY,
    user_uid uuid,
    topic_uid uuid default '99c0203e-a6da-45d6-98f3-021ac86adff9',
    solver_uid uuid,
    from_email varchar,
    from_phone varchar,
    title varchar,
    description varchar,
    status support_ticket_status,
    created_at timestamp,
    updated_at timestamp
);

CREATE INDEX idx_support_tickets_topic ON support_tickets(topic_uid);

CREATE TABLE support_tickets_comment_messages (
    uid uuid PRIMARY KEY,
    ticket_uid uuid,
    sender_uid uuid,
    content varchar,
    is_imported_from_email boolean,
    created_at timestamp,
    updated_at timestamp
);

CREATE TABLE support_tickets_solutions (
    ticket_uid uuid REFERENCES support_tickets(uid) ON DELETE CASCADE,
    description varchar,
    email_text varchar,
    is_success boolean,
    created_at timestamp,
    updated_at timestamp,
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE support_tickets_solutions;
DROP TABLE support_tickets_comment_messages;
DROP INDEX idx_support_tickets_topic;
DROP TABLE support_tickets;
DROP TYPE support_ticket_status;
DROP TYPE support_tickets_topics;
-- +goose StatementEnd
