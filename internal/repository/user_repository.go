package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/thetramp22/rifflog/internal/models"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1, $2)
	`

	_, err := r.DB.Exec(
		context.Background(),
		query,
		user.Email,
		user.PasswordHash,
	)

	return err
}
