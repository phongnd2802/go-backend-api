package initializes

import (
	"fmt"
	"github.com/phongnd2802/go-backend-api/global"
	"go.uber.org/zap"
)

func Run() {
	loadconfig()
	initLogger()
	global.Logger.Info("Config Log OK", zap.String("ok", "success"))
	initMysql()
	initRedis()

	route := initRoute()
	route.Run(fmt.Sprintf(":%v", global.Config.Server.Port))
}
