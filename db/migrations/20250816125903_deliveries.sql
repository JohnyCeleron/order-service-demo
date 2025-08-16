-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS deliveries(
    Id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    phone VARCHAR(50) NOT NULL CHECK (phone ~ '^\+[0-9]{10,}$'),
    zip VARCHAR(50) NOT NULL,
    city VARCHAR(50) NOT NULL,
    address VARCHAR(50) NOT NULL,
    region VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL CHECK (email ~* '^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$')
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE deliveries IF EXISTS;
-- +goose StatementEnd
