package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/get", func(c *gin.Context) {
		c.String(http.StatusOK, "get")
	})
	r.POST("/post", func(c *gin.Context) {
		c.String(http.StatusOK, "post")
	})
	r.Handle("DELETE", "/delete", func(c *gin.Context) {
		c.String(http.StatusOK, "delete")
	})

	r.Any("/any", func(c *gin.Context) {
		c.String(http.StatusOK, "any")
	})

	r.Run(":9999")

}
