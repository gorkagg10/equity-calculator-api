CREATE TABLE assets(
    id uuid PRIMARY KEY,
    name text NOT NULL,
    symbol varchar(10) UNIQUE NOT NULL,
    currency varchar(10) NOT NULL,
    exchange varchar(10) NOT NULL,
    price numeric(18,4) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);