-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS payments (
    Id SERIAL PRIMARY KEY,
    transaction VARCHAR(100) NOT NULL UNIQUE,
    request_id VARCHAR(100),
    currency VARCHAR(10) NOT NULL,
    provider VARCHAR(50) NOT NULL,
    amount INT NOT NULL CHECK (amount >= 0),
    payment_dt INT NOT NULL,
    bank VARCHAR(50) NOT NULL,
    delivery_cost INT NOT NULL CHECK (delivery_cost >= 0),
    goods_total INT NOT NULL CHECK (goods_total >= 0),
    custom_fee INT NOT NULL CHECK (custom_fee >= 0)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payments;
-- +goose StatementEnd
