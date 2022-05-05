package domain

import (
	"context"
	"mime/multipart"
	"time"
)

type Post struct {
	Id      int64
	UserId  int
	GroupId int
	Content string
	Created time.Time
	Image   string
}

type PostStore struct {
	Id      int64
	UserId  int
	GroupId int
	Content string
	Created time.Time
}

type PostGet struct {
	Limit int64
	Since int64
	Sort  string
	Desc  bool
}

type PostDelivery interface {
}

type PostUsecase interface {
	Store(ctx context.Context, req *PostStore, fh *multipart.FileHeader) error
	Get(ctx context.Context, req *PostGet) ([]Post, error)
}

type PostRepository interface {
	Store(ctx context.Context, p *Post) error
	Get(ctx context.Context, req *PostGet) ([]Post, error)
}
