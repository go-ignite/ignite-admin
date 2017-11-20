package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite-admin/conf"
	"github.com/go-ignite/ignite-admin/controllers"
	"github.com/go-ignite/ignite-admin/jobs"
	"github.com/go-ignite/ignite/utils"
	"github.com/go-xorm/xorm"
	"github.com/robfig/cron"
)

func main() {
	conf.InitConf()
	db := utils.InitDB(conf.DB_Driver, conf.DB_Connect)
	go initRouter(db)
	go initJob(db)
	select {}
}

func initRouter(db *xorm.Engine) {
	r := gin.Default()
	r.Static("/static/js", "./static/js")
	r.Static("/static/css", "./static/css")
	r.Static("/static/images", "./static/images")
	r.LoadHTMLGlob("static/*.html")
	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r, db)
}

func initJob(db *xorm.Engine) {
	jobs.SetDB(db)
	c := cron.New()
	c.AddFunc("* */5 * * * *", jobs.InstantStats)
	c.AddFunc("0 0 0 * * *", jobs.DailyStats)
	c.AddFunc("0 0 1 * * *", jobs.MonthlyStats)
	c.Start()
	select {}
}
