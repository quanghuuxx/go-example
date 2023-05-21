package authentication

import (
	"net/http"

	"golang_example/src/authentication/controller"
	"golang_example/utils"

	"github.com/gin-gonic/gin"
)

func Authorization(c *gin.Context) {
	result, e := controller.Authorizating(c)

	if e != nil {
		c.JSON(http.StatusNotFound, utils.WriteError(e))
		return
	}

	c.JSON(http.StatusOK, utils.WriteResponse(result))
}
