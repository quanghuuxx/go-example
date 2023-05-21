package main

import (
	"github.com/gin-gonic/gin"

	"golang_example/pkg/db"
	"golang_example/src/authentication"
	"golang_example/src/home"
)

func main() {
	db.Open()

	r := gin.Default()

	r.GET("/home", home.Middleware(), home.GetHomeData)
	r.POST("/authorization", authentication.Authorization)

	err := r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
