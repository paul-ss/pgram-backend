package usecase

import (
	"context"
	"github.com/paul-ss/pgram-backend/internal/app/domain"
	"mime/multipart"
)

type Usecase struct {
	postR domain.PostRepository
	statR domain.StaticRepository
}

func (uc *Usecase) Store(ctx context.Context, req *domain.PostStore, fh *multipart.FileHeader) error {
	filePath, err := uc.statR.StoreFile(fh)
	if err != nil {
		return err
	}

	post := domain.Post{
		Id:      req.Id,
		UserId:  req.UserId,
		GroupId: req.GroupId,
		Content: req.Content,
		Created: req.Created,
		Image:   filePath,
	}

	if err = uc.postR.Store(ctx, &post); err != nil {
		if err = uc.statR.DeleteFile(filePath); err != nil {
			return err
		}

		return err
	}

	return nil
}

func (uc *Usecase) Get(ctx context.Context, req *domain.PostGet) ([]domain.Post, error) {
	return uc.postR.Get(ctx, req)
}
