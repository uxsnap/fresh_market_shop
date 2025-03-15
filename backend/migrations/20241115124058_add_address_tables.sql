-- +goose Up
-- +goose StatementBegin
CREATE TABLE cities(
  uid uuid PRIMARY KEY,
  name VARCHAR(100) NOT NULL
);

INSERT INTO cities (uid, name) 
VALUES ('59c0203e-a6da-45d6-98f3-021ac86adff6', 'Санкт-Петербург');

CREATE TABLE addresses (
  uid uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  city_uid uuid NOT NULL REFERENCES cities(uid),
  street TEXT,
  house_number TEXT,
  latitude NUMERIC(11, 8) NOT NULL,
  longitude NUMERIC(11, 8) NOT NULL
);

CREATE INDEX ix_addresses_city_uid on addresses (city_uid);
CREATE INDEX ix_addresses_street on addresses (street);
CREATE INDEX ix_addresses_house_number on addresses (house_number);

CREATE TABLE addresses_streets_vectors (
  address_uid uuid,
  street_vector tsvector
);

CREATE INDEX idx_gin_addr_street
ON addresses_streets_vectors
USING gin (street_vector);

BEGIN;

CREATE TEMP TABLE addresses_import (
  street TEXT,
  house_number TEXT,
  latitude NUMERIC(11, 8) NOT NULL,
  longitude NUMERIC(11, 8) NOT NULL
) ON COMMIT DROP;

COPY addresses_import(street,house_number,latitude,longitude)
FROM '/osm/ADDR-RU-SPE.csv'
DELIMITER ';'
QUOTE E'\b' 
NULL AS ''
CSV HEADER;

INSERT INTO addresses(city_uid, street, house_number, latitude, longitude)
  SELECT 
    '59c0203e-a6da-45d6-98f3-021ac86adff6', 
    trim(replace(ai.street, 'улица', '')), 
    replace(ai.house_number, ' ', ''),
    ai.latitude, 
    ai.longitude
  FROM addresses_import as ai;

COMMIT;

INSERT INTO addresses_streets_vectors (address_uid, street_vector) 
SELECT a.uid, to_tsvector('russian', a.street) from addresses a;

CREATE TABLE delivery_addresses (
  uid uuid PRIMARY KEY,
  user_uid uuid NOT NULL REFERENCES users (uid) ON DELETE CASCADE,
  address_uid uuid NOT NULL REFERENCES addresses (uid) ON DELETE CASCADE,
  apartment INT,
  entrance INT,
  floor INT,
  code INT,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,

  UNIQUE (address_uid, apartment, entrance, floor, code)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP INDEX ix_addresses_city_uid;
DROP INDEX ix_addresses_street;
DROP INDEX ix_addresses_house_number;
DROP INDEX idx_gin_addr_street;

DROP TABLE cities;
DROP TABLE addresses_streets_vectors;
DROP TABLE addresses;
DROP TABLE delivery_addresses;

-- +goose StatementEnd
