package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelIndexHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/status")
}

func (router *MainRouter) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
