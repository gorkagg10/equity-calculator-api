CREATE TABLE positions(
    id uuid PRIMARY KEY,
    portfolio_id uuid references portfolios(id),
    asset_id uuid references assets(id),
    quantity numeric(18,8) NOT NULL,
    currency varchar(8) NOT NULL,
    created_at timestamp NOT NULL,
    updated_at timestamp NOT NULL,
    unique(portfolio_id,asset_id)
)