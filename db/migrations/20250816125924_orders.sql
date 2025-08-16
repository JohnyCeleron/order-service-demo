-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS orders(
    id VARCHAR(100) PRIMARY KEY,
    track_number VARCHAR(50) NOT NULL,
    entry VARCHAR(50) NOT NULL,
    delivery_id INT REFERENCES deliveries(Id),
    payment_id INT REFERENCES payments(Id),
    locale VARCHAR(10) NOT NULL,
    internal_signature VARCHAR(50) NOT NULL,
    customer_id VARCHAR(50) NOT NULL,
    delivery_service VARCHAR(50) NOT NULL,
    shardkey VARCHAR(50) NOT NULL,
    sm_id INT NOT NULL,
    date_created timestamptz NOT NULL,
    oof_shart VARCHAR(50) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE orders IF EXISTS;
-- +goose StatementEnd
