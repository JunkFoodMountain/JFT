-- +goose Up
-- +goose StatementBegin
SELECT 'Up models';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users
(
    id uuid DEFAULT uuid_generate_v4(),
    username varchar NOT NULL ,

    PRIMARY KEY (id)
);

CREATE TABLE publishers
(
    id uuid DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE games (
    id uuid DEFAULT uuid_generate_v4(),
    name varchar NOT NULL,
    publisher_id uuid NOT NULL,

    l2_tx varchar NULL,
    owner_user_id uuid NULL,

    PRIMARY KEY (id),
    CONSTRAINT fk_gamer_owner
        FOREIGN KEY (owner_user_id)
            REFERENCES users(id),
    CONSTRAINT fk_publisher
        FOREIGN KEY (publisher_id)
            REFERENCES publishers(id)
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'Down models';

DROP TABLE users;
DROP TABLE publishers;
DROP TABLE games;

-- +goose StatementEnd
