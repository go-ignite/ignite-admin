package main

import (
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-ignite/ignite-admin/controllers"
	"github.com/go-ignite/ignite-admin/jobs"
	utility "github.com/go-ignite/ignite-admin/utils"
	"github.com/go-ignite/ignite/utils"
	"github.com/go-xorm/xorm"
	"github.com/robfig/cron"
)

var confPath = flag.String("c", "./conf/config.toml", "config file")

func main() {
	utility.InitConf(*confPath)
	db := utils.InitDB(utility.DB_Driver, utility.DB_Connect)

	// For normal mode
	if len(os.Args) == 1 {
		go initRouter(db)
		go initJob(db)
		select {}
	} else if os.Args[1] == "recover" {
		// TODO: Restore docker containers from DB
		jobs.RecoverTask(db)
	}
}

func initRouter(db *xorm.Engine) {
	r := gin.Default()
	r.Static("/static/js", "./static/js")
	r.Static("/static/css", "./static/css")
	r.Static("/static/images", "./static/images")
	r.Static("/static/fonts", "./static/fonts")
	r.LoadHTMLGlob("static/*.html")
	mainRouter := &controllers.MainRouter{}
	mainRouter.Initialize(r, db)
}

func initJob(db *xorm.Engine) {
	jobs.SetDB(db)
	c := cron.New()
	c.AddFunc("0 */5 * * * *", jobs.InstantStats)
	c.AddFunc("0 0 0 * * *", jobs.DailyStats)
	c.AddFunc("0 0 0 1 * *", jobs.MonthlyStats)
	c.Start()
	select {}
}
