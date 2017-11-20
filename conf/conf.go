package conf

import (
	"flag"
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

var (
	confPath = flag.String("c", "./config.toml", "config file")
)

var (
	// for app config
	APP_Address string

	// for db config
	DB_Driver, DB_Connect string

	// for auth config
	Auth_Username, Auth_Password, Auth_Secret string
)

func InitConf() {
	//Check config file
	if _, err := os.Stat(*confPath); os.IsNotExist(err) {
		fmt.Println("Cannot load config.toml, file doesn't exist...")
		os.Exit(1)
	}

	config, err := toml.LoadFile(*confPath)

	if err != nil {
		fmt.Println("Failed to load config file:", *confPath)
		fmt.Println(err.Error())
		os.Exit(1)
	}
	APP_Address = config.Get("app.address").(string)

	DB_Driver = config.Get("db.driver").(string)
	DB_Connect = config.Get("db.connect").(string)

	Auth_Username = config.Get("auth.username").(string)
	Auth_Password = config.Get("auth.password").(string)
	Auth_Secret = config.Get("auth.secret").(string)
}
