-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id CHAR(36) PRIMARY KEY, 
    user_name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- Insert 10 sample users with meaningful names and 'password' as password
INSERT INTO users (id, user_name, first_name, last_name, password)
VALUES
    (UUID(), 'john_doe', 'John', 'Doe', 'password'),
    (UUID(), 'jane_smith', 'Jane', 'Smith', 'password'),
    (UUID(), 'michael_johnson', 'Michael', 'Johnson', 'password'),
    (UUID(), 'emily_davis', 'Emily', 'Davis', 'password'),
    (UUID(), 'david_miller', 'David', 'Miller', 'password'),
    (UUID(), 'susan_wilson', 'Susan', 'Wilson', 'password'),
    (UUID(), 'mark_brown', 'Mark', 'Brown', 'password'),
    (UUID(), 'lucy_jones', 'Lucy', 'Jones', 'password'),
    (UUID(), 'robert_white', 'Robert', 'White', 'password'),
    (UUID(), 'lisa_clark', 'Lisa', 'Clark', 'password');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
