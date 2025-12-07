CREATE TABLE IF NOT EXISTS accounts (
    account_no INT PRIMARY KEY,
    account_name VARCHAR(255) NOT NULL,
    account_group INT NOT NULL CHECK (account_group >= 1 AND account_group <= 8),
    tax_standard VARCHAR(50),
    type VARCHAR(10) NOT NULL CHECK (type IN ('P&L', 'BS')),
    standard_side VARCHAR(10) NOT NULL CHECK (standard_side IN ('Debit', 'Credit')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_accounts_group ON accounts(account_group);
CREATE INDEX idx_accounts_type ON accounts(type);
