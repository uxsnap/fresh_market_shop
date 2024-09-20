-- +goose Up
-- SQL в этой секции будет выполнен для обновления БД

-- Создание таблицы "users"
CREATE TABLE users (
    uid uuid PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    phone VARCHAR(20),
    is_deleted BOOLEAN,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE cities_handbook (
    uid uuid PRIMARY KEY,
    name VARCHAR(20) NOT NULL
);

-- Создание таблицы "addresses" для хранения адресов доставки пользователей с широтой и долготой
CREATE TABLE delivery_addresses (
    uid uuid PRIMARY KEY,
    user_uid uuid NOT NULL REFERENCES users(uid) ON DELETE CASCADE,
    latitude NUMERIC(10, 8) NOT NULL,  -- Широта с точностью до 8 знаков после запятой
    longitude NUMERIC(11, 8) NOT NULL, -- Долгота с точностью до 8 знаков после запятой
    city_uid uuid NOT NULL,
    street_name VARCHAR(30) NOT NULL,
    house_number INT NOT NULL,
    building INT NOT NULL,
    floor INT NOT NULL,
    apartment INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Создание таблицы "payment_cards" для хранения карт оплаты
CREATE TABLE payment_cards (
    uid uuid PRIMARY KEY,
    user_uid uuid NOT NULL REFERENCES users(uid) ON DELETE CASCADE,
    card_number VARCHAR(16) NOT NULL,
    card_holder_name VARCHAR(100) NOT NULL,
    expiration_date DATE NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Создание таблицы "categories" для классификации товаров
CREATE TABLE categories (
    uid uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE,
    description TEXT
);

-- Создание таблицы "products" для хранения данных о товарах
CREATE TABLE products (
    uid uuid PRIMARY KEY,
    category_uid uuid NOT NULL REFERENCES categories(uid) ON DELETE SET NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    ccal INT,
    price INT NOT NULL,
    stock_quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE product_photos (
    product_uid uuid NOT NULL,
    photo_external_uid uuid NOT NULL,
    order_idx INT,
    created_at TIMESTAMP NOT NULL
);

-- Создание таблицы "recipes" для хранения рецептов, где используются продукты
CREATE TABLE recipes (
    uid uuid PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- Связь продуктов с рецептами (многие ко многим)
CREATE TABLE recipe_products (
    recipe_uid uuid NOT NULL REFERENCES recipes(uid) ON DELETE CASCADE,
    product_uid uuid NOT NULL,
    quantity NUMERIC(10, 2) NOT NULL,
    name VARCHAR(255) NOT NULL,
    PRIMARY KEY (recipe_uid, product_uid)
);

-- +goose Down
-- SQL в этой секции будет выполнен для отката изменений

-- Удаление всех созданных таблиц
DROP TABLE IF EXISTS recipe_products;
DROP TABLE IF EXISTS recipes;
DROP TABLE IF EXISTS product_photos;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS payment_cards;
DROP TABLE IF EXISTS delivery_addresses;
DROP TABLE IF EXISTS cities_handbook;
DROP TABLE IF EXISTS users;