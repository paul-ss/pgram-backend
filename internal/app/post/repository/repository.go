package repository

import (
	"context"
	"github.com/jackc/pgx/v4"
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

func (r *Repository) Store(ctx context.Context, p *domain.Post) error {
	_, err := r.db.Exec(ctx, "")
	return err
}

func (r *Repository) Get(ctx context.Context, req *domain.PostGet) ([]domain.Post, error) {
	var post domain.Post
	var posts []domain.Post

	_, err := r.db.QueryFunc(ctx,
		`select id, user_id, group_id, content, created, image from posts
				where created > (select created from posts where id = $1)
				order by created
				limit $2`,
		[]interface{}{req.Since, req.Limit},
		[]interface{}{&post.Id, &post.UserId, &post.GroupId, &post.Content, &post.Created, &post.Image},
		func(row pgx.QueryFuncRow) error {
			posts = append(posts, post)
			return nil
		})

	if err != nil {
		return nil, err
	}

	return posts, nil
}
