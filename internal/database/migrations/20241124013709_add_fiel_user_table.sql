-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN email VARCHAR(255) NOT NULL,
ADD COLUMN avatar VARCHAR(255) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN IF EXISTS email,
DROP COLUMN IF EXISTS avatar;
-- +goose StatementEnd
