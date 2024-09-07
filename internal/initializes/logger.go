package initializes

import (
	"github.com/phongnd2802/go-backend-api/global"
	"github.com/phongnd2802/go-backend-api/pkg/logger"
)

func initLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
