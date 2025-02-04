-- +goose Up
-- +goose StatementBegin
select 'up SQL query';

create table tile (
    id    text primary key,
    title text                                     not null,
    type  text check (type in ('link', 'display')) default 'link' not null,
    link  text,
    image text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
select 'down SQL query';


-- +goose StatementEnd