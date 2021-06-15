package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/*action", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World !!!")
	})

	r.Run(":9999")
}
