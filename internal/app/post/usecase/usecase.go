package usecase

import (
	"context"
	"errors"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	postRepository "github.com/paul-ss/pgram-backend/internal/app/post/repository"
	staticRepository "github.com/paul-ss/pgram-backend/internal/app/static/repository"
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

func (uc *Usecase) Store(ctx context.Context, req *domain.PostStoreUC, fh *multipart.FileHeader) (*domain.Post, error) {
	filePath, err := uc.statR.StoreFile(fh)
	if err != nil {
		return nil, err
	}

	post := domain.PostStoreR{
		UserId:  req.UserId,
		GroupId: req.GroupId,
		Content: req.Content,
		Created: req.Created,
		Image:   filePath,
	}

	res, err := uc.postR.Store(ctx, &post)
	if err != nil {
		if err = uc.statR.DeleteFile(filePath); err != nil {
			return nil, err
		}
		return nil, err
	}

	return res, nil
}

func (uc *Usecase) Get(ctx context.Context, req *domain.PostGetUC) ([]domain.Post, error) {
	reqUC := &domain.PostGetR{
		Since: req.Since,
		Limit: req.Limit,
		Desc:  req.Desc,
	}

	switch req.Sort {
	case "created":
		return uc.postR.GetSortCreated(ctx, reqUC)
	default:
		return nil, errors.New("")
	}
}
