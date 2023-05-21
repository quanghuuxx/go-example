package storage

import (
	"golang_example/pkg/db"
	"golang_example/src/home/models"
)

type HomeStorage interface {
	FetchPackages() ([]models.PackageViewModel, error)
}

func New() HomeStorage {
	return _HomeStorageImpl{
		db: db.Source(),
	}
}
