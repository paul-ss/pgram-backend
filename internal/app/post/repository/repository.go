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

func (r *Repository) Store(ctx context.Context, p *domain.PostStoreR) (*domain.Post, error) {
	var res domain.Post
	err := r.db.QueryRow(ctx,
		`insert into posts (user_id, group_id, content, created, image)
				values ($1, $2, $3, $4, $5)
				returning id, user_id, group_id, content, created, image`,
		p.UserId, p.GroupId, p.Content, p.Created, p.Image,
	).Scan(&res.Id, &res.UserId, &res.GroupId, &res.Content, &res.Created, &res.Image)

	return &res, err
}

func (r *Repository) GetSortCreated(ctx context.Context, req *domain.PostGetR) ([]domain.Post, error) {
	desc := ""
	if req.Desc {
		desc = "desc "
	}

	if req.Since <= 0 {
		return r.get(ctx,
			"select id, user_id, group_id, content, created, image from posts "+
				"order by created "+desc+
				"limit $1 ",
			req.Limit)
	}

	return r.get(ctx,
		"select id, user_id, group_id, content, created, image from posts "+
			"where created > (select created from posts where id = $1) "+
			"order by created "+desc+
			"limit $2 ",
		req.Since, req.Limit)
}

func (r *Repository) get(ctx context.Context, query string, args ...interface{}) ([]domain.Post, error) {
	var post domain.Post
	var posts []domain.Post

	_, err := r.db.QueryFunc(ctx, query, args,
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
