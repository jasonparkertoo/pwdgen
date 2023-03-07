package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/status", Status)
	r.POST("/api/v1/password", Pwd)
	r.SetTrustedProxies(nil)
	r.Run()
}
