package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./assets")
	r.StaticFS("/static", http.Dir("static"))
	r.StaticFile("/favicon.ico", "./favivon.ico")


	r.Run(":9999")

}
