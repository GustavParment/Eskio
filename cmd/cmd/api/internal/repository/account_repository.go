package repository

import (
	"cmd/api/internal/domain"
	"database/sql"
	"fmt"
)

type AccountRepository interface {
	CreateAccount(account *domain.Account) error
	GetAccountByNo(accountNo int) (*domain.Account, error)
	GetAllAccounts() ([]*domain.Account, error)
	GetAccountsByGroup(accountGroup int) ([]*domain.Account, error)
	UpdateAccount(account *domain.Account) error
	DeleteAccount(accountNo int) error
}

type accountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) CreateAccount(account *domain.Account) error {
	query := `
		INSERT INTO accounts (account_no, account_name, account_group, tax_standard, type, standard_side)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	_, err := r.db.Exec(query,
		account.AccountNo,
		account.AccountName,
		account.AccountGroup,
		account.TaxStandard,
		account.Type,
		account.StandardSide,
	)
	if err != nil {
		return fmt.Errorf("failed to create account: %w", err)
	}

	return nil
}

func (r *accountRepository) GetAccountByNo(accountNo int) (*domain.Account, error) {
	query := `
		SELECT account_no, account_name, account_group, tax_standard, type, standard_side
		FROM accounts
		WHERE account_no = ?
	`
	account := &domain.Account{}
	err := r.db.QueryRow(query, accountNo).Scan(
		&account.AccountNo,
		&account.AccountName,
		&account.AccountGroup,
		&account.TaxStandard,
		&account.Type,
		&account.StandardSide,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("account not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	return account, nil
}

func (r *accountRepository) GetAllAccounts() ([]*domain.Account, error) {
	query := `
		SELECT account_no, account_name, account_group, tax_standard, type, standard_side
		FROM accounts
		ORDER BY account_no
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts: %w", err)
	}
	defer rows.Close()

	var accounts []*domain.Account
	for rows.Next() {
		account := &domain.Account{}
		err := rows.Scan(
			&account.AccountNo,
			&account.AccountName,
			&account.AccountGroup,
			&account.TaxStandard,
			&account.Type,
			&account.StandardSide,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan account: %w", err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *accountRepository) GetAccountsByGroup(accountGroup int) ([]*domain.Account, error) {
	query := `
		SELECT account_no, account_name, account_group, tax_standard, type, standard_side
		FROM accounts
		WHERE account_group = ?
		ORDER BY account_no
	`
	rows, err := r.db.Query(query, accountGroup)
	if err != nil {
		return nil, fmt.Errorf("failed to get accounts by group: %w", err)
	}
	defer rows.Close()

	var accounts []*domain.Account
	for rows.Next() {
		account := &domain.Account{}
		err := rows.Scan(
			&account.AccountNo,
			&account.AccountName,
			&account.AccountGroup,
			&account.TaxStandard,
			&account.Type,
			&account.StandardSide,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan account: %w", err)
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (r *accountRepository) UpdateAccount(account *domain.Account) error {
	query := `
		UPDATE accounts
		SET account_name = ?, account_group = ?, tax_standard = ?, type = ?, standard_side = ?
		WHERE account_no = ?
	`
	_, err := r.db.Exec(query,
		account.AccountName,
		account.AccountGroup,
		account.TaxStandard,
		account.Type,
		account.StandardSide,
		account.AccountNo,
	)
	if err != nil {
		return fmt.Errorf("failed to update account: %w", err)
	}

	return nil
}

func (r *accountRepository) DeleteAccount(accountNo int) error {
	query := `DELETE FROM accounts WHERE account_no = ?`
	_, err := r.db.Exec(query, accountNo)
	if err != nil {
		return fmt.Errorf("failed to delete account: %w", err)
	}

	return nil
}