CREATE OR REPLACE FUNCTION update_balance()
RETURNS TRIGGER AS $$
DECLARE
    last_balance DOUBLE PRECISION := 0;
    has_balance BOOLEAN;
BEGIN
    SELECT EXISTS (
        SELECT 1 FROM balances WHERE wallet_id = NEW.wallet_id
    ) INTO has_balance;

    IF has_balance THEN
        SELECT balance INTO last_balance
        FROM balances
        WHERE wallet_id = NEW.wallet_id
        ORDER BY timestamp DESC
        LIMIT 1;
    END IF;

    IF NEW.operation_type = 'DEPOSIT' THEN
        last_balance := last_balance + NEW.amount;
    ELSIF NEW.operation_type = 'WITHDRAW' THEN
        last_balance := last_balance - NEW.amount;
		IF last_balance < 0 THEN
			RAISE EXCEPTION USING 
				MESSAGE = NEW.wallet_id,
				ERRCODE = 'P1001';
        END IF;
    END IF;

    INSERT INTO balances (wallet_id, balance, timestamp)
    VALUES (NEW.wallet_id, last_balance, NEW.timestamp);

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_update_balance
AFTER INSERT ON transactions
FOR EACH ROW
EXECUTE FUNCTION update_balance();