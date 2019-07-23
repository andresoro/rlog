-- +migrate Up
CREATE TABLE sites (
    id SERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(50)
);

CREATE TABLE page_views (
    id SERIAL PRIMARY KEY NOT NULL,
    site_id INTEGER,
    duration INTEGER,
    host VARCHAR(50),
    url_path VARCHAR(200),
    referrer VARCHAR(200),
    request_date DATE,
    uniq BOOLEAN,
    FOREIGN KEY (site_id) REFERENCES sites (id)
);

-- +migrate Down
DROP TABLE IF EXISTS sites;
DROP TABLE IF EXISTS page_views;