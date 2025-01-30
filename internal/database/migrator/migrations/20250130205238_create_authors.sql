-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE authors (
                         id   INTEGER PRIMARY KEY,
                         name text    NOT NULL,
                         bio  text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE authors
-- +goose StatementEnd
