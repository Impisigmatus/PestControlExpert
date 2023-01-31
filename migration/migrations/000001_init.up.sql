CREATE SCHEMA IF NOT EXISTS main;

CREATE TABLE IF NOT EXISTS main.subscribers (
    chat_id     INTEGER NOT NULL UNIQUE,
    username    TEXT NOT NULL,
    name        TEXT
);

CREATE TABLE IF NOT EXISTS main.callbacks (
    date                DATE NOT NULL,
    name                TEXT NOT NULL,
    phone VARCHAR(10)   NOT NULL,
    description         TEXT
);

CREATE TABLE IF NOT EXISTS main.price (
    name        TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL UNIQUE,
    standart    REAL NOT NULL,
    premium     REAL NOT NULL
);

CREATE OR REPLACE FUNCTION main.delete_expired() RETURNS TRIGGER
AS $$
BEGIN
    DELETE FROM main.callbacks WHERE date < NOW() - INTERVAL '1 month';
    RETURN NEW;
END
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER delete_expired_trigger
AFTER INSERT OR UPDATE ON main.callbacks
EXECUTE PROCEDURE main.delete_expired();
