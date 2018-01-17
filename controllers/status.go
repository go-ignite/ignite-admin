package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite/models"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelStatusListHandler(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	users := new([]*models.User)
	router.db.Desc("created").Limit(pageSize, pageSize*(pageIndex-1)).Find(users)
	for _, user := range *users {
		if user.ServiceType == "" {
			if user.ServiceId != "" {
				user.ServiceType = "SS"
			} else {
				user.ServiceType = "N/A"
			}
		}
	}

	user := new(models.User)
	total, _ := router.db.Count(user)

	pd := models.PageData{Total: total, PageSize: pageSize, PageIndex: pageIndex, Data: users}
	resp := models.Response{Success: true, Message: "success", Data: pd}
	c.JSON(http.StatusOK, resp)
}
