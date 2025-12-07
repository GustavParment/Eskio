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
