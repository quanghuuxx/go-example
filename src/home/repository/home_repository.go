package repository

import (
	"golang_example/src/home/storage"
)

type HomeRepository interface {
	FetchPackages() ([]any, error)
}

func New(storage storage.HomeStorage) HomeRepository {
	return _HomeRepositoryImpl{
		storage: storage,
	}
}
