package repository

import (
	"golang_example/model"
	"golang_example/pkg/jwt"
)

type _AuthenticationRepositoryImpl struct{}

func (_AuthenticationRepositoryImpl) CreateTokenByRequest(request model.AuthorizationRequest) (string, error) {
	return jwt.GenerateAuthorizationToken(request)
}
