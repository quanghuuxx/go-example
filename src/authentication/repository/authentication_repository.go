package repository

import "golang_example/model"

type AuthenticationRepository interface {
	CreateTokenByRequest(request model.AuthorizationRequest) (string, error)
}

func New() AuthenticationRepository {
	return _AuthenticationRepositoryImpl{}
}
