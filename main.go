package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	svc := gin.Default()

	svc.GET("/", GetserverStatus)
	svc.Run()
}

func GetserverStatus(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "I am up and running")
}
