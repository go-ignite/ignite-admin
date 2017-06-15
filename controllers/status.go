package controllers

import (
	"ignite/models"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router *MainRouter) StatusHandler(c *gin.Context) {
	session := sessions.Default(c)
	v := session.Get("userId")
	var uInfo *models.UserInfo
	if v != nil {
		if uId, ok := v.(int64); ok {
			uInfo = &models.UserInfo{
				Id: uId,
			}
		}
	}

	c.HTML(http.StatusOK, "status.html", gin.H{
		"uInfo": uInfo,
	})
}
