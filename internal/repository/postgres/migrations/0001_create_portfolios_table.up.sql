CREATE TABLE portfolios(
    id uuid PRIMARY KEY,
    name varchar(20) UNIQUE NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL
);