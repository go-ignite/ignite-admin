package main

import (
	"ignite-admin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static/js", "./static/js")
	r.Static("/static/css", "./static/css")
	r.Static("/static/images", "./static/images")
	r.LoadHTMLGlob("static/*.html")

	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r)
}
