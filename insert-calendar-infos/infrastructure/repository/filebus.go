package repository

import (
	"insert-calendar-infos/infrastructure"
	"insert-calendar-infos/infrastructure/filebus"
)

type filebusRepository struct {
	filebusHandler filebus.IFilebusHandler
}

// NewFilebusRepository filebus repository constructor
func NewFilebusRepository(infra *infrastructure.Infra) IFilebusRepository {
	return &filebusRepository{
		filebusHandler: infra.FilebusHandler,
	}
}

// DownloadFile download file from filebus use file ID
func (f *filebusRepository) DownloadFile(path, file string) ([]byte, error) {
	return f.filebusHandler.Download(path, file)
}
