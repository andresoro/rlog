-- +migrate Up
CREATE TABLE sites (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50)
);

-- +migrate Down
DROP TABLE IF EXISTS sites;