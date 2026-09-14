CREATE TABLE transactions(
    id uuid PRIMARY KEY,
    symbol varchar(10),
    shares numeric(18,8),
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);