-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id UUID PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert 10 sample users with meaningful names and 'password' as password
INSERT INTO users (id, user_name, first_name, last_name, password)
VALUES
    (gen_random_uuid(), 'john_doe', 'John', 'Doe', 'password'),
    (gen_random_uuid(), 'jane_smith', 'Jane', 'Smith', 'password'),
    (gen_random_uuid(), 'michael_johnson', 'Michael', 'Johnson', 'password'),
    (gen_random_uuid(), 'emily_davis', 'Emily', 'Davis', 'password'),
    (gen_random_uuid(), 'david_miller', 'David', 'Miller', 'password'),
    (gen_random_uuid(), 'susan_wilson', 'Susan', 'Wilson', 'password'),
    (gen_random_uuid(), 'mark_brown', 'Mark', 'Brown', 'password'),
    (gen_random_uuid(), 'lucy_jones', 'Lucy', 'Jones', 'password'),
    (gen_random_uuid(), 'robert_white', 'Robert', 'White', 'password'),
    (gen_random_uuid(), 'lisa_clark', 'Lisa', 'Clark', 'password');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
