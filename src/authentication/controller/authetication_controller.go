package controller

import (
	"net/http"

	"golang_example/model"
	"golang_example/src/authentication/repository"

	"github.com/gin-gonic/gin"
)

func Authorizating(c *gin.Context) (*model.AuthoriztionReponse, error) {
	var request model.AuthorizationRequest

	e := c.BindJSON(&request)

	if e != nil || request.Empty() {
		return nil, http.ErrBodyNotAllowed
	}

	token, e := repository.New().CreateTokenByRequest(request)

	if e != nil {
		return nil, e
	}

	return &model.AuthoriztionReponse{
		Token: token,
	}, nil
}
