package repository

import (
	"cmd/api/internal/domain"
	"database/sql"
	"fmt"
)

type ReportRepository interface {
	GetIncomeStatement(fromDate, toDate string) (*domain.IncomeStatement, error)
}

type reportRepository struct {
	db *sql.DB
}

func NewReportRepository(db *sql.DB) ReportRepository {
	return &reportRepository{db: db}
}

func (r *reportRepository) GetIncomeStatement(fromDate, toDate string) (*domain.IncomeStatement, error) {
	query := `
		SELECT
			a.account_no,
			a.account_name,
			a.type,
			SUM(l.debit_amount - l.credit_amount) as balance
		FROM line_items l
		INNER JOIN vouchers v ON l.voucher_id = v.voucher_id
		INNER JOIN accounts a ON l.account_no = a.account_no
		WHERE v.date >= $1
		  AND v.date <= $2
		  AND v.corrected_by_voucher_id IS NULL
		  AND a.type = 'P&L'
		GROUP BY a.account_no, a.account_name, a.type
		HAVING SUM(l.debit_amount - l.credit_amount) != 0
		ORDER BY a.account_no
	`

	rows, err := r.db.Query(query, fromDate, toDate)
	if err != nil {
		return nil, fmt.Errorf("failed to query income statement: %w", err)
	}
	defer rows.Close()

	statement := &domain.IncomeStatement{
		Income:   make([]domain.IncomeStatementEntry, 0),
		Expenses: make([]domain.IncomeStatementEntry, 0),
	}
	statement.Period.FromDate = fromDate
	statement.Period.ToDate = toDate

	for rows.Next() {
		var accountNo int
		var accountName string
		var accountType string
		var balance float64

		err := rows.Scan(&accountNo, &accountName, &accountType, &balance)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		entry := domain.IncomeStatementEntry{
			AccountNo:   accountNo,
			AccountName: accountName,
			Balance:     balance,
		}

		// Income accounts (3000-3999): positive balance means income
		// Expense accounts (4000-8999): negative balance means expense
		if accountNo >= 3000 && accountNo < 4000 {
			statement.Income = append(statement.Income, entry)
			statement.TotalIncome += balance
		} else if accountNo >= 4000 && accountNo < 9000 {
			statement.Expenses = append(statement.Expenses, entry)
			statement.TotalExpenses += balance
		}
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating rows: %w", err)
	}

	// Calculate net result
	statement.NetResult = statement.TotalIncome + statement.TotalExpenses

	return statement, nil
}
