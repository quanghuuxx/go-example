package home

import (
	"net/http"

	"golang_example/pkg/jwt"
	"golang_example/src/home/controller"
	"golang_example/utils"

	"github.com/gin-gonic/gin"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		e := jwt.ValidateRequest(c.Request)

		if e != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, utils.WriteError(e))
			return
		}

		c.Next()
	}
}

func GetHomeData(c *gin.Context) {
	result, e := controller.GetHomeData(c)

	if e != nil {
		c.JSON(http.StatusNotFound, utils.WriteError(e))
		return
	}

	c.JSON(http.StatusOK, utils.WriteResponse(result))
}