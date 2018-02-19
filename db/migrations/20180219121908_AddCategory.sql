
-- +goose Up

CREATE TABLE category (
    id integer NOT NULL,
    name varchar NOT NULL,
    parent_id integer,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    guid_id uuid,
    PRIMARY KEY (id)
);

CREATE SEQUENCE category_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE category_id_seq OWNED BY category.id;

-- +goose Down

DROP TABLE category;