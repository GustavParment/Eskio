-- Eskio Bookkeeping Database Schema
-- This file combines all migrations for initial database setup

-- Migration 001: Create users table
CREATE TABLE IF NOT EXISTS users (
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,
    role VARCHAR(50) NOT NULL DEFAULT 'Bookkeeper',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_users_email ON users(email);

-- Migration 002: Create accounts table
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

-- Migration 003: Create vouchers table
CREATE TABLE IF NOT EXISTS vouchers (
    voucher_id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    description TEXT NOT NULL,
    reference VARCHAR(255),
    total_amount DECIMAL(15, 2) NOT NULL,
    period VARCHAR(7) NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (created_by) REFERENCES users(user_id) ON DELETE RESTRICT
);

CREATE INDEX idx_vouchers_period ON vouchers(period);
CREATE INDEX idx_vouchers_created_by ON vouchers(created_by);
CREATE INDEX idx_vouchers_date ON vouchers(date);

-- Migration 004: Create line_items table
CREATE TABLE IF NOT EXISTS line_items (
    line_id SERIAL PRIMARY KEY,
    voucher_id INT NOT NULL,
    account_no INT NOT NULL,
    debit_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    credit_amount DECIMAL(15, 2) NOT NULL DEFAULT 0,
    tax_code INT,
    project_id INT,
    cost_center_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (voucher_id) REFERENCES vouchers(voucher_id) ON DELETE CASCADE,
    FOREIGN KEY (account_no) REFERENCES accounts(account_no) ON DELETE RESTRICT,
    CHECK (
        (debit_amount > 0 AND credit_amount = 0) OR
        (credit_amount > 0 AND debit_amount = 0)
    )
);

CREATE INDEX idx_line_items_voucher ON line_items(voucher_id);
CREATE INDEX idx_line_items_account ON line_items(account_no);

-- Insert some sample BAS accounts for testing
INSERT INTO accounts (account_no, account_name, account_group, tax_standard, type, standard_side) VALUES
(1510, 'Kundfordringar', 1, '0%', 'BS', 'Debit'),
(1930, 'Företagskonto / checkkonto', 1, '0%', 'BS', 'Debit'),
(2440, 'Leverantörsskulder', 2, '0%', 'BS', 'Credit'),
(2610, 'Utgående moms', 2, '25%', 'BS', 'Credit'),
(2640, 'Ingående moms', 2, '25%', 'BS', 'Debit'),
(3000, 'Försäljning varor och tjänster', 3, '25%', 'P&L', 'Credit'),
(4000, 'Inköp varor och material', 4, '25%', 'P&L', 'Debit'),
(6000, 'Lokalkostnader', 4, '25%', 'P&L', 'Debit')
ON CONFLICT (account_no) DO NOTHING;
