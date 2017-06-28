package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) InviteCodeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "code.html", nil)
}
