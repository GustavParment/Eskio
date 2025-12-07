package repository

import (
	"cmd/api/internal/domain"
	"database/sql"
	"fmt"
)

type VoucherRepository interface {
	CreateVoucher(voucher *domain.Voucher) error
	GetVoucherByID(voucherID int) (*domain.Voucher, error)
	GetAllVouchers() ([]*domain.Voucher, error)
	GetVouchersByPeriod(period string) ([]*domain.Voucher, error)
	GetVouchersByCreatedBy(userID int) ([]*domain.Voucher, error)
	UpdateVoucher(voucher *domain.Voucher) error
	DeleteVoucher(voucherID int) error
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
		VALUES (?, ?, ?, ?, ?, ?)
	`
	result, err := r.db.Exec(query,
		voucher.Date,
		voucher.Description,
		voucher.Reference,
		voucher.TotalAmount,
		voucher.Period,
		voucher.CreatedBy,
	)
	if err != nil {
		return fmt.Errorf("failed to create voucher: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("failed to get last insert id: %w", err)
	}

	voucher.VoucherID = int(id)
	return nil
}

func (r *voucherRepository) GetVoucherByID(voucherID int) (*domain.Voucher, error) {
	query := `
		SELECT voucher_id, date, description, reference, total_amount, period, created_by
		FROM vouchers
		WHERE voucher_id = ?
	`
	voucher := &domain.Voucher{}
	err := r.db.QueryRow(query, voucherID).Scan(
		&voucher.VoucherID,
		&voucher.Date,
		&voucher.Description,
		&voucher.Reference,
		&voucher.TotalAmount,
		&voucher.Period,
		&voucher.CreatedBy,
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
		SELECT voucher_id, date, description, reference, total_amount, period, created_by
		FROM vouchers
		ORDER BY date DESC, voucher_id DESC
	`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers: %w", err)
	}
	defer rows.Close()

	var vouchers []*domain.Voucher
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.Date,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
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
		SELECT voucher_id, date, description, reference, total_amount, period, created_by
		FROM vouchers
		WHERE period = ?
		ORDER BY date DESC, voucher_id DESC
	`
	rows, err := r.db.Query(query, period)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers by period: %w", err)
	}
	defer rows.Close()

	var vouchers []*domain.Voucher
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.Date,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
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
		SELECT voucher_id, date, description, reference, total_amount, period, created_by
		FROM vouchers
		WHERE created_by = ?
		ORDER BY date DESC, voucher_id DESC
	`
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get vouchers by user: %w", err)
	}
	defer rows.Close()

	var vouchers []*domain.Voucher
	for rows.Next() {
		voucher := &domain.Voucher{}
		err := rows.Scan(
			&voucher.VoucherID,
			&voucher.Date,
			&voucher.Description,
			&voucher.Reference,
			&voucher.TotalAmount,
			&voucher.Period,
			&voucher.CreatedBy,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan voucher: %w", err)
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}

func (r *voucherRepository) UpdateVoucher(voucher *domain.Voucher) error {
	query := `
		UPDATE vouchers
		SET date = ?, description = ?, reference = ?, total_amount = ?, period = ?, created_by = ?
		WHERE voucher_id = ?
	`
	_, err := r.db.Exec(query,
		voucher.Date,
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
	query := `DELETE FROM vouchers WHERE voucher_id = ?`
	_, err := r.db.Exec(query, voucherID)
	if err != nil {
		return fmt.Errorf("failed to delete voucher: %w", err)
	}

	return nil
}
