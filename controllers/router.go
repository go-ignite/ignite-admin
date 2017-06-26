package controllers

import (
	"flag"
	"fmt"
	"os"

	"github.com/gin-gonic/contrib/jwt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	toml "github.com/pelletier/go-toml"
)

var (
	conf     = flag.String("c", "./config.toml", "config file")
	authUser string
	authPass string
	secret   string
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
	)

	authUser = config.Get("auth.username").(string)
	authPass = config.Get("auth.password").(string)

	fmt.Printf("auth username:%s \r\n", authUser)
	fmt.Printf("auth password:%s \r\n", authPass)

	secret = config.Get("auth.secret").(string)

	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", user, password, host, dbname)
	engine, _ := xorm.NewEngine("mysql", connString)

	err = engine.Ping()

	if err != nil {
		fmt.Println("Cannot connetc to database:", err.Error())
		os.Exit(1)
	}

	self.db = engine
	self.router = r
	self.router.GET("/", self.PanelIndexHandler)
	self.router.POST("/login", self.PanelLoginHandler)
	self.router.GET("/status", self.PanelStatusHandler)

	pg := self.router.Group("/auth")
	pg.Use(jwt.Auth(secret))
	pg.GET("/status_list", self.PanelStatusListHandler)
	pg.PUT("/:id/reset", self.ResetAccountHandler)
	pg.PUT("/:id/destroy", self.DestroyAccountHandler)

	self.router.Run(":8000")
}
