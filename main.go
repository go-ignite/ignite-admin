package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite/controllers"
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
