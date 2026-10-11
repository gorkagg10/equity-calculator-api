CREATE TABLE fx_rates (
    currency CHAR(3) PRIMARY KEY,
    symbol CHAR(10) UNIQUE,
    rate_to_eur NUMERIC(20, 10) NOT NULL
        CHECK (rate_to_eur > 0),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);