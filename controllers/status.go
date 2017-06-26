package controllers

import (
	"ignite/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) PanelStatusHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "status.html", nil)
}

func (router *MainRouter) PanelStatusListHandler(c *gin.Context) {
	users := new([]*models.User)
	router.db.Find(users)

	resp := models.Response{Success: true, Message: "success", Data: users}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) ResetAccountHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	user := new(models.User)
	user.PackageUsed = 0

	router.db.Id(uid).Cols("package_used").Update(user)
	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) DestroyAccountHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("uid"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	user := new(models.User)
	router.db.Id(uid).Get(user)

	//TODO:
	//1. Destroy user's container
	//2. Delete user's account

	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}
