-- +goose Up
CREATE TABLE "user" (
    id serial primary key,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    role int not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
DROP TABLE "user";
