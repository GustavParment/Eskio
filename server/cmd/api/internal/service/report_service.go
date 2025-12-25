package service

import (
	"cmd/api/internal/domain"
	"cmd/api/internal/repository"
	"errors"
	"fmt"
	"time"
)

type ReportService struct {
	repository repository.ReportRepository
}

func NewReportService(repo repository.ReportRepository) *ReportService {
	return &ReportService{
		repository: repo,
	}
}

func (s *ReportService) GetIncomeStatement(fromDate, toDate string) (*domain.IncomeStatement, error) {
	// Validate dates
	if fromDate == "" || toDate == "" {
		return nil, errors.New("from_date and to_date are required")
	}

	// Parse and validate date format
	_, err := time.Parse("2006-01-02", fromDate)
	if err != nil {
		return nil, fmt.Errorf("invalid from_date format, expected YYYY-MM-DD: %w", err)
	}

	_, err = time.Parse("2006-01-02", toDate)
	if err != nil {
		return nil, fmt.Errorf("invalid to_date format, expected YYYY-MM-DD: %w", err)
	}

	// Validate that from_date is before to_date
	from, _ := time.Parse("2006-01-02", fromDate)
	to, _ := time.Parse("2006-01-02", toDate)
	if from.After(to) {
		return nil, errors.New("from_date must be before or equal to to_date")
	}

	// Get income statement from repository
	statement, err := s.repository.GetIncomeStatement(fromDate, toDate)
	if err != nil {
		return nil, fmt.Errorf("failed to generate income statement: %w", err)
	}

	return statement, nil
}
