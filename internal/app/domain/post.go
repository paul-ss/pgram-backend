package domain

import (
	"context"
	"mime/multipart"
	"time"
)

type Post struct {
	Id      int64     `json:"id"`
	UserId  int       `json:"user_id"`
	GroupId int       `json:"group_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Image   string    `json:"image"`
}

type PostStoreUC struct {
	UserId  int       `json:"user_id"`
	GroupId int       `json:"group_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
}

type PostStoreR struct {
	UserId  int       `json:"user_id"`
	GroupId int       `json:"group_id"`
	Content string    `json:"content"`
	Created time.Time `json:"created"`
	Image   string    `json:"image"`
}

type PostGetUC struct {
	Limit int64  `json:"limit"`
	Since int64  `json:"since"`
	Sort  string `json:"sort"`
	Desc  bool   `json:"desc"`
}

type PostGetR struct {
	Limit int64 `json:"limit"`
	Since int64 `json:"since"`
	Desc  bool  `json:"desc"`
}

type PostDelivery interface {
}

type PostUsecase interface {
	Store(ctx context.Context, req *PostStoreUC, fh *multipart.FileHeader) (*Post, error)
	Get(ctx context.Context, req *PostGetUC) ([]Post, error)
}

type PostRepository interface {
	Store(ctx context.Context, p *PostStoreR) (*Post, error)
	GetSortCreated(ctx context.Context, req *PostGetR) ([]Post, error)
}
