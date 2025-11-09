

-- Users table
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Rewards table
CREATE TABLE IF NOT EXISTS rewards (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    stock_symbol TEXT NOT NULL,
    shares NUMERIC(12,6) NOT NULL,
    rewarded_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Stock prices table
CREATE TABLE IF NOT EXISTS stock_prices (
    stock_symbol TEXT PRIMARY KEY,
    price_in_inr NUMERIC(12,2) NOT NULL,
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

--  Daily valuation table 
CREATE TABLE IF NOT EXISTS daily_stock_values (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    date DATE NOT NULL,
    total_value_in_inr NUMERIC(12,2) NOT NULL
);

-- inserting a  sample user
INSERT INTO users (name)
VALUES ('Divyanshu')
ON CONFLICT DO NOTHING;

