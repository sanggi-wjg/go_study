package main

import (
	"go_study/home"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", home.Ping)
	router.GET("/user/:name", home.HelloName)
	router.GET("/excel", home.CreateExcel)

	router.Run(":8080")
}
