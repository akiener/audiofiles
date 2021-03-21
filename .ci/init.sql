DROP TABLE IF EXISTS song;
DROP TABLE IF EXISTS dc_user;
DROP TABLE IF EXISTS artist;
DROP TABLE IF EXISTS album;

CREATE TABLE album
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE artist
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE dc_user
(
    id   BIGSERIAL PRIMARY KEY,
    name TEXT
);

CREATE TABLE song
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL,
    artist_id   BIGSERIAL REFERENCES artist ON DELETE CASCADE,
    album_id    BIGSERIAL REFERENCES album ON DELETE CASCADE,
    genre       TEXT NOT NULL,
    description TEXT NOT NULL,
    dc_user_id  BIGSERIAL REFERENCES dc_user ON DELETE CASCADE
);
