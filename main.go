package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HomePageHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "hello world!"})
}

func main() {
	router := gin.Default()
	router.GET("/", HomePageHandler)
	err := router.Run()
	if err != nil {
		return
	}
}
