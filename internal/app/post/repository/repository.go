package repository

import (
	"context"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	postgres "github.com/paul-ss/pgram-backend/internal/pkg/database"
)

func NewRepository() *Repository {
	return &Repository{
		db: postgres.GetConn(),
	}
}

type Repository struct {
	db postgres.PgxConn
}

func (r *Repository) Store(post *domain.Post) error {
	_, err := r.db.Exec(context.Background(), "")
	return err
}
