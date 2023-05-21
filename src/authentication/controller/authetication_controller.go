package controller

import (
	"encoding/json"
	"net/http"

	"golang_example/model"
	"golang_example/src/authentication/repository"
)

func Handle(r *http.Request) (any, error) {
	var request model.AuthorizationRequest

	e := json.NewDecoder(r.Body).Decode(&request)

	if e != nil || request.Empty() {
		return nil, http.ErrBodyNotAllowed
	}

	switch r.Method {
	case http.MethodGet:
		// TODO: do somthing here
		return nil, http.ErrAbortHandler
	case http.MethodPost:
		return Authorizating(request)
	default:
		return nil, http.ErrAbortHandler
	}
}

func Authorizating(request model.AuthorizationRequest) (*model.AuthoriztionReponse, error) {
	token, e := repository.New().CreateTokenByRequest(request)

	if e != nil {
		return nil, e
	}

	return &model.AuthoriztionReponse{
		Token: token,
	}, nil
}



