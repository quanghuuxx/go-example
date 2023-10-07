package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"golang_example/pkg/db"
	"golang_example/src/authentication"
	"golang_example/src/home"
)

func main() {
	db.Open()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()
	
	r.GET("/home", home.Middleware(), home.GetHomeData)
	r.POST("/authorization", authentication.Authorization)

	err = r.Run(":9999")
	if err != nil {
		panic(err)
	}
}
