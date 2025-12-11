package repository

import (
	"cmd/api/internal/domain"
	"database/sql"
	"fmt"
)

type VoucherRepository interface {
	CreateVoucher(voucher *domain.Voucher) error
	CreateCorrectionVoucher(voucher *domain.Voucher, originalVoucherID int) error
	GetVoucherByID(voucherID int) (*domain.Voucher, error)
	GetAllVouchers() ([]*domain.Voucher, error)
	GetVouchersByPeriod(period string) ([]*domain.Voucher, error)
	GetVouchersByCreatedBy(userID int) ([]*domain.Voucher, error)
	UpdateVoucher(voucher *domain.Voucher) error
	DeleteVoucher(voucherID int) error
	MarkVoucherAsCorrected(voucherID int, correctedByID int) error
}

type voucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepository{db: db}
}

func (r *voucherRepository) CreateVoucher(voucher *domain.Voucher) error {
	query := `
		INSERT INTO vouchers (date, description, reference, total_amount, period, created_by)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING voucher_id, voucher_number
	`
	err := r.db.QueryRow(query,
		voucher.Date.Time,
		voucher.Description,
		voucher.Reference,
		voucher.TotalAmount,
		voucher.Period,
		voucher.CreatedBy,
	).Scan(&voucher.VoucherID, &voucher.VoucherNumber)
	if err != nil {
		return fmt.Errorf("failed to create voucher: %w", err)
	}

	return nil
}

func (r *voucherRepository) CreateCorrectionVoucher(voucher *domain.Voucher, originalVoucherID int) error {
	query := `
		INSERT INTO vouchers (date, description, reference, total_amount, period, created_by, corrects_voucher_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING voucher_id, voucher_number
	`
	err := r.db.QueryRow(query,
		voucher.Date.Time,
		voucher.Description,
		voucher.Reference,
		voucher.TotalAmount,
		voucher.Period,
		voucher.CreatedBy,
		originalVoucherID,
	).Scan(&voucher.VoucherID, &voucher.VoucherNumber)
	if err != nil {
		return fmt.Errorf("failed to create correction voucher: %w", err)
	}
	voucher.CorrectsVoucherID = &originalVoucherID

	return nil
}

func (r *voucherRepository) MarkVoucherAsCorrected(voucherID int, correctedByID int) error {
	query := `UPDATE vouchers SET corrected_by_voucher_id = $1 WHERE voucher_id = $2`
	_, err := r.db.Exec(query, correctedByID, voucherID)
	if err != nil {
		return fmt.Errorf("failed to mark voucher as corrected: %w", err)
	}
	return nil
}

func (r *voucherRepository) GetVoucherByID(voucherID int) (*domain.Voucher, error) {
	query := `
		SELECT voucher_id, voucher_number, date, description, reference, total_amount, period, created_by, corrects_voucher_id, corrected_by_voucher_id
		FROM vouchers
		WHERE voucher_id = $1
	`
	voucher := &domain.Voucher{}
	err := r.db.QueryRow(query, voucherID).Scan(
		&voucher.VoucherID,
		&voucher.VoucherNumber,
		&voucher.Date.Time,
		&voucher.Description,
		&voucher.Reference,
		&voucher.TotalAmount,
		&voucher.Period,
		&voucher.CreatedBy,
		&voucher.CorrectsVoucherID,
		&voucher.CorrectedByVoucherID,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("voucher not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get voucher: %w", err)
	}

	return voucher, nil
}

func (r *voucherRepository) GetAllVouchers() ([]*domain.Voucher, error) {
	query := `
		SELECT voucher_id, voucher_number, date, description, reference, total_amount, period, created_by, corrects_voucher_id, corrected_by_voucher_id
		FROM vouchers
		ORDER BY voucher_number DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers: %w", err)
	}
	defer rows.Close()

	vouchers := make([]*domain.Voucher, 0)
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.VoucherNumber,
			&voucher.Date.Time,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
			&voucher.CorrectsVoucherID,
			&voucher.CorrectedByVoucherID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan voucher: %w", err)
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}

func (r *voucherRepository) GetVouchersByPeriod(period string) ([]*domain.Voucher, error) {
	query := `
		SELECT voucher_id, voucher_number, date, description, reference, total_amount, period, created_by, corrects_voucher_id, corrected_by_voucher_id
		FROM vouchers
		WHERE period = $1
		ORDER BY voucher_number DESC
	`
	rows, err := r.db.Query(query, period)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers by period: %w", err)
	}
	defer rows.Close()

	vouchers := make([]*domain.Voucher, 0)
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.VoucherNumber,
			&voucher.Date.Time,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
			&voucher.CorrectsVoucherID,
			&voucher.CorrectedByVoucherID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan voucher: %w", err)
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}

func (r *voucherRepository) GetVouchersByCreatedBy(userID int) ([]*domain.Voucher, error) {
	query := `
		SELECT voucher_id, voucher_number, date, description, reference, total_amount, period, created_by, corrects_voucher_id, corrected_by_voucher_id
		FROM vouchers
		WHERE created_by = $1
		ORDER BY voucher_number DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers by user: %w", err)
	}
	defer rows.Close()

	vouchers := make([]*domain.Voucher, 0)
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.VoucherNumber,
			&voucher.Date.Time,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
			&voucher.CorrectsVoucherID,
			&voucher.CorrectedByVoucherID,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan voucher: %w", err)
		}
		vouchers = append(vouchers, voucher)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating vouchers: %w", err)
	}

	return vouchers, nil
}

func (r *voucherRepository) UpdateVoucher(voucher *domain.Voucher) error {
	query := `
		UPDATE vouchers
		SET date = $1, description = $2, reference = $3, total_amount = $4, period = $5, created_by = $6
		WHERE voucher_id = $7
	`
	_, err := r.db.Exec(query,
		voucher.Date.Time,
		voucher.Description,
		voucher.Reference,
		voucher.TotalAmount,
		voucher.Period,
		voucher.CreatedBy,
		voucher.VoucherID,
	)
	if err != nil {
		return fmt.Errorf("failed to update voucher: %w", err)
	}

	return nil
}

func (r *voucherRepository) DeleteVoucher(voucherID int) error {
	query := `DELETE FROM vouchers WHERE voucher_id = $1`
	_, err := r.db.Exec(query, voucherID)
	if err != nil {
		return fmt.Errorf("failed to delete voucher: %w", err)
	}

	return nil
}
