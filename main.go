package main

import (
	"fmt"
	"grs-server/core"
	"grs-server/global"
	"grs-server/routers"
)

func main() {

	global.Config = core.InitConfig()	//	gin的配置
	global.DB = core.InitMysql()		//	gorm的配置

	fmt.Println(global.Config)

	r := routers.Routers()

	addr := global.Config.System.Addr()

	err := r.Run(addr)
	if err != nil {
		return
	}
}
