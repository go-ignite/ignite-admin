package controllers

import (
	"flag"

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
	self.router.POST("/login", self.PanelLoginHandler)

	pg := self.router.Group("/auth")
	pg.Use(jwt.Auth(utility.Auth_Secret))

	//user account related operations
	pg.GET("/status_list", self.PanelStatusListHandler)
	pg.PUT("/:id/reset", self.ResetAccountHandler)
	pg.PUT("/:id/destroy", self.DestroyAccountHandler)
	pg.PUT("/:id/stop", self.StopServiceHandler)
	pg.PUT("/:id/start", self.StartServiceHandler)

	//invite code related operations
	pg.GET("/code_list", self.InviteCodeListHandler)
	pg.PUT("/:id/remove", self.RemoveInviteCodeHandler)
	pg.POST("/code_generate", self.GenerateInviteCodeHandler)

	self.router.Run(utility.APP_Address)
}
