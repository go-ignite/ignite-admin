package controllers

import (
	"fmt"
	"net/http"

	"ignite/models"
	"ignite/ss"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (router *MainRouter) PanelIndexHandler(c *gin.Context) {
	userID, exists := c.Get("userId")

	if !exists {
		c.HTML(http.StatusOK, "panel.html", nil)
		return
	}

	user := new(models.User)
	router.db.Id(userID).Get(user)

	uInfo := &models.UserInfo{
		Id:           user.Id,
		Host:         ss.Host,
		Username:     user.Username,
		Status:       user.Status,
		PackageUsed:  fmt.Sprintf("%.2f", user.PackageUsed),
		PackageLimit: user.PackageLimit,
		PackageLeft:  fmt.Sprintf("%.2f", float32(user.PackageLimit)-user.PackageUsed),
		ServicePort:  user.ServicePort,
		ServicePwd:   user.ServicePwd,
		Expired:      user.Expired.Format("2006-01-02"),
	}
	if user.PackageLimit == 0 {
		uInfo.PackageLeftPercent = "0"
	} else {
		uInfo.PackageLeftPercent = fmt.Sprintf("%.2f", (float32(user.PackageLimit)-user.PackageUsed)/float32(user.PackageLimit)*100)
	}

	c.HTML(http.StatusOK, "panel.html", gin.H{
		"uInfo": uInfo,
	})
}

func (router *MainRouter) LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("userId")
	session.Save()

	c.Redirect(http.StatusFound, "/")
}
