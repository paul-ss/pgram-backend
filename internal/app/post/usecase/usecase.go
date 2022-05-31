package usecase

import (
	"context"
	"errors"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	postRepository "github.com/paul-ss/pgram-backend/internal/app/post/repository"
	staticRepository "github.com/paul-ss/pgram-backend/internal/app/static/repository"
	"github.com/paul-ss/pgram-backend/internal/pkg/pointers"
	"mime/multipart"
)

func NewUsecase() *Usecase {
	return &Usecase{
		postR: postRepository.NewRepository(),
		statR: staticRepository.NewRepository(),
	}
}

type Usecase struct {
	postR domain.PostRepository
	statR domain.StaticRepository
}

func (uc *Usecase) Create(ctx context.Context, req *domain.PostCreate, fh *multipart.FileHeader) (*domain.Post, error) {
	filePath, err := uc.statR.StoreFile(fh)
	if err != nil {
		return nil, err
	}

	post := domain.PostStoreR{
		UserId:  req.UserId,
		GroupId: pointers.New(*req.GroupId),
		Content: pointers.New(*req.Content),
		Created: req.Created,
		Image:   pointers.New(filePath),
	}

	res, err := uc.postR.Create(ctx, &post)
	if err != nil {
		if err = uc.statR.DeleteFile(filePath); err != nil {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

func (uc *Usecase) GetFeed(ctx context.Context, req *domain.FeedGet) ([]domain.Post, error) {
	reqUC := &domain.FeedGetRepo{
		Since: req.Since,
		Limit: req.Limit,
		Desc:  req.Desc,
	}

	switch req.Sort {
	case "created":
		return uc.postR.GetFeedCreated(ctx, reqUC)
	default:
		return nil, errors.New("")
	}
}
