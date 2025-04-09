CREATE TABLE IF NOT EXISTS balances (
    wallet_id UUID NOT NULL,
    balance DOUBLE PRECISION NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    PRIMARY KEY(wallet_id, timestamp)
);