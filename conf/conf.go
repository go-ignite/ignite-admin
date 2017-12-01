package conf

import (
	"flag"
	"fmt"
	"os"

	toml "github.com/pelletier/go-toml"
)

var (
	confPath = flag.String("c", "./data/config.toml", "config file")
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
	if _, err := os.Stat(*confPath); !os.IsNotExist(err) {
		if config, err := toml.LoadFile(*confPath); err == nil {
			APP_Address = config.Get("app.address").(string)

			DB_Driver = config.Get("db.driver").(string)
			DB_Connect = config.Get("db.connect").(string)

			Auth_Username = config.Get("auth.username").(string)
			Auth_Password = config.Get("auth.password").(string)
			Auth_Secret = config.Get("auth.secret").(string)
		}
	}
	if driver := os.Getenv("DB_DRIVER"); driver != "" {
		DB_Driver = driver
	}
	if connect := os.Getenv("DB_CONNECT"); connect != "" {
		DB_Connect = connect
	}
	if username := os.Getenv("AUTH_USERNAME"); username != "" {
		Auth_Username = username
	}
	if password := os.Getenv("AUTH_PASSWORD"); password != "" {
		Auth_Password = password
	}
	if secret := os.Getenv("Auth_SECRET"); secret != "" {
		Auth_Secret = secret
	}
	fmt.Println("config: ", map[string]interface{}{
		"address":       APP_Address,
		"db_driver":     DB_Driver,
		"db_connect":    DB_Connect,
		"auth_username": Auth_Username,
		"auth_password": Auth_Password,
		"auth_secret":   Auth_Secret,
	})
}
