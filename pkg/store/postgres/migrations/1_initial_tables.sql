-- +migrate Up
CREATE TABLE sites (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50)
);

CREATE TABLE events (
    id SERIAL PRIMARY KEY NOT NULL,
    site_id INTEGER,
    host VARCHAR(50),
    event_key VARCHAR(200),
    addr VARCHAR(50),
    request_date DATE,
    uniq BOOLEAN,
    FOREIGN KEY (site_id) REFERENCES sites (id)
);

-- +migrate Down
DROP TABLE IF EXISTS sites;
DROP TABLE IF EXISTS events;