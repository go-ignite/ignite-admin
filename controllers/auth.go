package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite/models"
	"github.com/go-ignite/ignite/ss"
)

func (router *MainRouter) ResetAccountHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))

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
	uid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	user := new(models.User)
	router.db.Id(uid).Get(user)

	//1. Destroy user's container
	if user.ServiceId != "" {
		err = ss.RemoveContainer(user.ServiceId)

		if err != nil {
			resp := models.Response{Success: false, Message: "终止用户容器失败!"}
			c.JSON(http.StatusOK, resp)
			return
		}
	}

	//2. Delete user's account
	_, err = router.db.Id(uid).Delete(new(models.User))
	if err != nil {
		resp := models.Response{Success: false, Message: "删除用户失败!"}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) StopServiceHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	user := new(models.User)
	router.db.Id(uid).Get(user)

	//1. Stop user's container
	if ss.IsContainerRunning(user.ServiceId) {
		err = ss.StopContainer(user.ServiceId)

		if err != nil {
			resp := models.Response{Success: false, Message: "停止服务失败"}
			c.JSON(http.StatusOK, resp)
			return
		}

		//2. Update service status
		user.Status = 2
		router.db.Id(uid).Cols("status").Update(user)
	}

	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) StartServiceHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	user := new(models.User)
	router.db.Id(uid).Get(user)

	//1. Start user's container
	if !ss.IsContainerRunning(user.ServiceId) {
		err = ss.StartContainer(user.ServiceId)

		if err != nil {
			resp := models.Response{Success: false, Message: "启动服务失败"}
			c.JSON(http.StatusOK, resp)
			return
		}

		//2. Update service status
		user.Status = 1
		router.db.Id(uid).Cols("status").Update(user)
	}

	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) RenewServiceHandler(c *gin.Context) {
	uid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := models.Response{Success: false, Message: "用户ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	renewEntity := struct {
		Month int `json:"month"`
	}{}

	if err := c.BindJSON(&renewEntity); err != nil || renewEntity.Month == 0 {
		resp := models.Response{Success: false, Message: "续费参数不正确!"}
		c.JSON(http.StatusOK, &resp)
		return
	}

	user := new(models.User)
	router.db.Id(uid).Get(user)
	if user.Id == 0 {
		resp := models.Response{Success: false, Message: "获取用户失败!"}
		c.JSON(http.StatusOK, &resp)
		return
	}

	user.Expired = user.Expired.AddDate(0, renewEntity.Month, 0)
	if user.Expired.Before(user.Created) {
		resp := models.Response{Success: false, Message: "服务到期时间不能早于创建时间!"}
		c.JSON(http.StatusOK, &resp)
		return
	}

	if user.ServiceId != "" {
		var err error
		running := ss.IsContainerRunning(user.ServiceId)
		now := time.Now()
		if running && user.Expired.Before(now) {
			err = ss.StopContainer(user.ServiceId)
			if err == nil {
				user.Status = 2
			}
		} else if !running && user.Expired.After(now) {
			err = ss.StartContainer(user.ServiceId)
			if err == nil {
				user.Status = 1
			}
		}
		if err != nil {
			resp := models.Response{Success: false, Message: "启动/停止服务失败!"}
			c.JSON(http.StatusOK, &resp)
			return
		}
	}

	if _, err := router.db.Id(uid).Cols("expired", "status").Update(user); err != nil {
		resp := models.Response{Success: false, Message: "更新过期时间失败!"}
		c.JSON(http.StatusOK, &resp)
		return
	}

	updateMap := map[string]interface{}{
		"id":      user.Id,
		"expired": user.Expired.Format("2006-01-02"),
		"status":  user.Status,
	}
	resp := models.Response{Success: true, Message: "success", Data: updateMap}
	c.JSON(http.StatusOK, resp)
}
