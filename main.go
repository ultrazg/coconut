package main

import (
	S "coconut/server"
	U "coconut/util"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()

	router.GET("/", S.Test) // test router
	router.POST("/search", S.Search)

	U.PrintVersion()

	err := router.Run(":8080")
	if err != nil {
		panic("server start fail")

		return
	}
}
