CREATE TYPE operation_type_enum AS ENUM ('DEPOSIT', 'WITHDRAW');

CREATE TABLE IF NOT EXISTS transactions (
    transaction_id UUID PRIMARY KEY,
    wallet_id UUID NOT NULL,
    operation_type operation_type_enum NOT NULL,
    amount DOUBLE PRECISION  NOT NULL,
    timestamp TIMESTAMPTZ NOT NULL DEFAULT NOW()
);