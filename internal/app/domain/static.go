package domain

import "mime/multipart"

type StaticRepository interface {
	StoreFile(fh *multipart.FileHeader) (filepath string, err error)
	DeleteFile(filepath string) error
}
