-- 000001_create_pictograms_table.up.sql
CREATE TABLE IF NOT EXISTS pictograms (
    id          INTEGER PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    categories  TEXT[] NOT NULL,
    tags        TEXT[] NOT NULL,
    image_url   VARCHAR(255) NOT NULL,
    schematic   BOOLEAN NOT NULL DEFAULT FALSE,
    violence    BOOLEAN NOT NULL DEFAULT FALSE,
    sex         BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
