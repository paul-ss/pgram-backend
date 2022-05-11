package repository

import (
	"github.com/google/uuid"
	"github.com/paul-ss/pgram-backend/internal/pkg/config"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

const subdirsNuber = 3

func NewRepository() *Repository {
	return &Repository{
		staticDir: config.C().Static.StaticDir,
	}
}

type Repository struct {
	staticDir string
}

func (r *Repository) generateRelPath() (string, error) {
	ud, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}

	splitedPath := strings.SplitN(ud.String(), "", subdirsNuber-1)
	return filepath.Join(splitedPath...), nil
}

// StoreFile saves file and returns relative filepath in static dir or error
func (r *Repository) StoreFile(fh *multipart.FileHeader) (relPath string, err error) {
	src, err := fh.Open()
	if err != nil {
		return
	}
	defer src.Close()

	relPath, err = r.generateRelPath()
	if err != nil {
		return
	}

	out, err := os.Create(filepath.Join(r.staticDir, relPath))
	if err != nil {
		return
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return
}

func (r *Repository) DeleteFile(relPath string) error {
	return os.Remove(filepath.Join(r.staticDir, relPath))
}
