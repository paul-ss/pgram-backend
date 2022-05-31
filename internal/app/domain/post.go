package domain

import (
	"context"
	"mime/multipart"
	"time"
)

// Post is a database model
type Post struct {
	Id      int64     `json:"Id"`
	UserId  int       `json:"UserId"`
	GroupId *int      `json:"GroupId,omitempty"`
	Content *string   `json:"Content,omitempty"`
	Created time.Time `json:"Created"`
	Image   *string   `json:"Image,omitempty"`
}

// PostCreate is api model
type PostCreate struct {
	UserId  int       `json:"UserId"`
	GroupId *int      `json:"GroupId" validate:"required"`
	Content *string   `json:"Content"`
	Created time.Time `json:"Created"`
}

// PostResponse is api model
type PostResponse struct {
	Post Post
}

type PostStoreR struct {
	UserId  int       `json:"UserId"`
	GroupId *int      `json:"GroupId"`
	Content *string   `json:"Content"`
	Created time.Time `json:"Created"`
	Image   *string   `json:"Image"`
}

type PostDelivery interface {
}

type PostUsecase interface {
	Create(ctx context.Context, req *PostCreate, fh *multipart.FileHeader) (*Post, error)
	GetFeed(ctx context.Context, req *FeedGet) ([]Post, error)
}

type PostRepository interface {
	Create(ctx context.Context, p *PostStoreR) (*Post, error)
	GetFeedCreated(ctx context.Context, req *FeedGetRepo) ([]Post, error)
}
