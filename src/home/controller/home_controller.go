package controller

import (
	"net/http"

	"golang_example/pkg/jwt"
	"golang_example/src/home/repository"
	"golang_example/src/home/storage"
)

func Handle(r *http.Request) (any, error) {
	e := jwt.ValidateRequest(r)

	if e != nil {
		return nil, e
	}

	switch r.Method {
	case http.MethodGet:
		return GetHomeData()
	case http.MethodPost:
		// TODO: do somthing here
		return nil, http.ErrAbortHandler
	default:
		return nil, http.ErrAbortHandler
	}
}

func GetHomeData() ([]any, error) {
	re := repository.New(storage.New())

	packages, e := re.FetchPackages()

	if packages != nil {
		return packages, e
	}

	return nil, e
}
