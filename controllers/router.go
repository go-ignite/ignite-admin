package controllers

import (
	"flag"
	"net/http"
	"strings"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	utility "github.com/go-ignite/ignite-admin/utils"
	"github.com/go-xorm/xorm"
)

type MainRouter struct {
	router *gin.Engine
	db     *xorm.Engine
}

func (self *MainRouter) Initialize(r *gin.Engine, db *xorm.Engine) {
	flag.Parse()

	self.router = r
	self.db = db
	self.router.GET("/", self.PanelIndexHandler)
	self.router.NoRoute(func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.String(), "/api/") {
			c.String(http.StatusNotFound, "not found")
			return
		}
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := self.router.Group("/api")
	{
		api.POST("/login", self.PanelLoginHandler)

		pg := api.Group("/auth")
		pg.Use(jwt.Auth(utility.Auth_Secret))
		{
			//user account related operations
			pg.GET("/status_list", self.PanelStatusListHandler)
			pg.PUT("/:id/reset", self.ResetAccountHandler)
			pg.PUT("/:id/destroy", self.DestroyAccountHandler)
			pg.PUT("/:id/stop", self.StopServiceHandler)
			pg.PUT("/:id/start", self.StartServiceHandler)
			pg.PUT("/:id/renew", self.RenewServiceHandler)

			//invite code related operations
			pg.GET("/code_list", self.InviteCodeListHandler)
			pg.PUT("/:id/remove", self.RemoveInviteCodeHandler)
			pg.POST("/code_generate", self.GenerateInviteCodeHandler)
		}
	}

	self.router.Run(utility.APP_Address)
}
