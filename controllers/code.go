package controllers

import (
	"ignite/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (router *MainRouter) InviteCodeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "code.html", nil)
}

func (router *MainRouter) InviteCodeListHandler(c *gin.Context) {
	pageIndex, _ := strconv.Atoi(c.Query("pageIndex"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))

	codes := new([]*models.InviteCode)
	router.db.Desc("created").Where("available=1").Limit(pageSize, pageSize*(pageIndex-1)).Find(codes)

	code := new(models.InviteCode)
	total, _ := router.db.Count(code)

	pd := models.PageData{Total: total, PageSize: pageSize, PageIndex: pageIndex, Data: codes}
	resp := models.Response{Success: true, Message: "success", Data: pd}
	c.JSON(http.StatusOK, resp)
}

func (router *MainRouter) RemoveInviteCodeHandler(c *gin.Context) {
	cid, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		resp := models.Response{Success: false, Message: "邀请码ID参数不正确"}
		c.JSON(http.StatusOK, resp)
		return
	}

	code := new(models.InviteCode)
	_, err = router.db.Id(cid).Delete(code)

	if err != nil {
		resp := models.Response{Success: false, Message: "邀请码删除失败"}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp := models.Response{Success: true, Message: "success"}
	c.JSON(http.StatusOK, resp)
}
