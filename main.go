package main

import (
	"go_study/home"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/ping", home.Ping)
	// This handler will match /user/john but will not match /user/ or /user
	router.GET("/user/:name", home.HelloName)
	// However, this one will match /user/john/ and also /user/john/send
	// If no other routers match /user/john, it will redirect to /user/john/
	router.GET("/user/:name/*action", home.HelloNameAction)
	router.GET("/excel", home.CreateExcel)

	authorized := router.Group("/")
	authorized.Use(AuthRequired()){
		
	}

	router.Run(":8080")
}
