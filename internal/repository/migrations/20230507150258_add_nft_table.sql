-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE nft
(
    id uuid default uuid_generate_v4() not null
        primary key,
    name varchar(255) not null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE nft
-- +goose StatementEnd
