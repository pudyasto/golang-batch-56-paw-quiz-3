-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category(
    id SERIAL NOT NULL,
    name VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP NULL
);

CREATE TABLE book(
    id SERIAL NOT NULL,
    title VARCHAR(256),
    description VARCHAR(256),
    image_url TEXT,
    release_year INT,
    price VARCHAR(256),
    total_page INT,
    thickness VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP NULL,
    category_id INT
);
-- +migrate StatementEnd