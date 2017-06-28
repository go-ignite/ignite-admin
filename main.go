package main

import (
	"ignite-admin/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/static", "./static")
	r.LoadHTMLGlob("static/*.html")

	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r)
}
