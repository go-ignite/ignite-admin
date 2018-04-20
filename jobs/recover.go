package jobs

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"github.com/go-ignite/ignite/models"
	"github.com/go-ignite/ignite/ss"
	"github.com/go-xorm/xorm"
)

var (
	checkSymbol = fmt.Sprintf("\x1b[0;32m%s\x1b[0m", "\u2714")
	crossSymbol = fmt.Sprintf("\x1b[0;31m%s\x1b[0m", "\u2716")
)

func RecoverTask(db *xorm.Engine) {
	// Check for the existence of user's container
	users := []*models.User{}
	if err := db.Where("service_id!=''").Find(&users); err != nil {
		log.Println("Query users error:", err.Error())
		os.Exit(1)
	}
	var wg sync.WaitGroup
	wg.Add(len(users))
	for _, user := range users {
		go recoverUser(db, &wg, user)
	}
	wg.Wait()
}

func recoverUser(db *xorm.Engine, wg *sync.WaitGroup, user *models.User) {
	defer wg.Add(-1)
	exist, err := ss.ContainerExist(user.ServiceId)
	if err != nil {
		fmt.Printf("%s [%s] inspect error:%v\n", crossSymbol, user.Username, err)
		return
	}
	if !exist {
		// Check port available
		ln, err := net.Listen("tcp", fmt.Sprintf(":%d", user.ServicePort))
		if err != nil {
			fmt.Printf("%s [%s] port[%d] is unavailable\n", crossSymbol, user.Username, user.ServicePort)
			return
		}
		ln.Close()

		if user.ServiceType == "" {
			user.ServiceType = "SS"
		}
		if user.ServiceMethod == "" {
			user.ServiceMethod = "aes-256-cfb"
		}
		// Create container
		result, err := ss.CreateAndStartContainer(user.ServiceType, strings.ToLower(user.Username), user.ServiceMethod, user.ServicePwd, user.ServicePort)
		if err != nil {
			fmt.Printf("%s [%s] create container error:%v\n", crossSymbol, user.Username, err)
			return
		}
		user.ServiceId = result.ID
		_, err = db.Id(user.Id).Cols("service_id").Update(user)
		if err != nil {
			fmt.Printf("%s [%s] update service id error:%v\n", crossSymbol, user.Username, err)
			return
		}
		if user.Status == 2 {
			// Start container
			if err := ss.StopContainer(user.ServiceId); err != nil {
				fmt.Printf("%s [%s] stop container error:%v\n", crossSymbol, user.Username, err)
				return
			}
		}
		fmt.Printf("%s [%s] container recovery succeeded\n", checkSymbol, user.Username)
		return
	}
	fmt.Printf("%s [%s] container is running\n", checkSymbol, user.Username)
}
