-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS order_items(
    id SERIAL PRIMARY KEY,
    chrt_id INT NOT NULL,
    track_number VARCHAR(50) NOT NULL,
    price INT NOT NULL,
    rid VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    sale INTEGER NOT NULL,
    size VARCHAR(10) NOT NULL,
    total_price INT NOT NULL,
    nm_id INT NOT NULL,
    brand VARCHAR(100) NOT NULL,
    status INTEGER NOT NULL,
    order_id VARCHAR(100) REFERENCES orders(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE order_items IF EXISTS;
-- +goose StatementEnd
