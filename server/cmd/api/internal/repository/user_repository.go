package repository

import (
	"cmd/api/internal/domain"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	CreateUser(user *domain.User) error
	GetUserByID(userID int) (*domain.User, error)
	GetUserByEmail(email string) (*domain.User, error)
	UpdateUser(user *domain.User) error
	DeleteUser(userID int) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *domain.User) error {
	query := `
		INSERT INTO users (name, email, password_hash, role)
		VALUES ($1, $2, $3, $4)
		RETURNING user_id
	`
	err := r.db.QueryRow(query, user.Name, user.Email, user.PasswordHash, user.Role).Scan(&user.UserID)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *userRepository) GetUserByID(userID int) (*domain.User, error) {
	query := `
		SELECT user_id, name, email, password_hash, role
		FROM users
		WHERE user_id = $1
	`
	user := &domain.User{}
	err := r.db.QueryRow(query, userID).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*domain.User, error) {
	query := `
		SELECT user_id, name, email, password_hash, role
		FROM users
		WHERE email = $1
	`
	user := &domain.User{}
	err := r.db.QueryRow(query, email).Scan(
		&user.UserID,
		&user.Name,
		&user.Email,
		&user.PasswordHash,
		&user.Role,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *userRepository) UpdateUser(user *domain.User) error {
	query := `
		UPDATE users
		SET name = $1, email = $2, password_hash = $3, role = $4
		WHERE user_id = $5
	`
	_, err := r.db.Exec(query, user.Name, user.Email, user.PasswordHash, user.Role, user.UserID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *userRepository) DeleteUser(userID int) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := r.db.Exec(query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
