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
