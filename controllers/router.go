package controllers

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	toml "github.com/pelletier/go-toml"
)

var (
	conf = flag.String("c", "./config.toml", "config file")
)

type MainRouter struct {
	router *gin.Engine
	db     *xorm.Engine
}

func (self *MainRouter) Initialize(r *gin.Engine) {
	flag.Parse()
	//Check config file
	if _, err := os.Stat(*conf); os.IsNotExist(err) {
		fmt.Println("Cannot load config.toml, file doesn't exist...")
		os.Exit(1)
	}

	config, err := toml.LoadFile(*conf)

	if err != nil {
		fmt.Println("Failed to load config file:", *conf)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Init DB connection
	var (
		user     = config.Get("mysql.user").(string)
		password = config.Get("mysql.password").(string)
		host     = config.Get("mysql.host").(string)
		dbname   = config.Get("mysql.dbname").(string)

		authUser = config.Get("auth.username").(string)
		authPass = config.Get("auth.password").(string)
	)

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, dbname)
	engine, _ := xorm.NewEngine("mysql", connString)

	err = engine.Ping()

	if err != nil {
		fmt.Println("Cannot connetc to database:", err.Error())
		os.Exit(1)
	}

	self.db = engine
	self.router = r

	pg := self.router.Group("/", gin.BasicAuth(gin.Accounts{
		authUser: authPass,
	}))

	pg.GET("/", self.PanelIndexHandler)
	pg.GET("/status", self.PanelStatusHandler)
	pg.GET("/logout", self.LogoutHandler)

	self.router.Run(":8000")
}
