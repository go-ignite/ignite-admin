package controllers

import (
	"ignite/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelStatusListHandler(c *gin.Context) {
	users := new([]*models.User)
	router.db.Desc("created").Find(users)

	resp := models.Response{Success: true, Message: "success", Data: users}
	c.JSON(http.StatusOK, resp)
}
