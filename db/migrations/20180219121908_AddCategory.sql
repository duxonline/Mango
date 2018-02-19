
-- +goose Up

CREATE TABLE category (
    id integer NOT NULL
    --name varchar NOT NULL,
    --parent_id integer,
    --updated_at timestamp with time zone,
    --guid_id uuid,
    --PRIMARY KEY (id),
);


-- +goose Down

DROP TABLE category;