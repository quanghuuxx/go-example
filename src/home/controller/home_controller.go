package controller

import (
	"golang_example/src/home/repository"
	"golang_example/src/home/storage"

	"github.com/gin-gonic/gin"
)

func GetHomeData(c *gin.Context) ([]any, error) {
	re := repository.New(storage.New())

	packages, e := re.FetchPackages()

	if packages != nil {
		return packages, e
	}

	return nil, e
}
