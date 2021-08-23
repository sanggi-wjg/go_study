package home

import (
	"go_study/myexcel"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func HelloName(c *gin.Context) {
	name := c.Param("name")

	c.String(http.StatusOK, "Hello %s", name)
}

func CreateExcel(c *gin.Context) {
	myexcel.CreateExcel()
}
