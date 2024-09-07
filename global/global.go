package global

import (
	"database/sql"
	"github.com/phongnd2802/go-backend-api/pkg/logger"
	"github.com/phongnd2802/go-backend-api/pkg/setting"
	"github.com/redis/go-redis/v9"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *sql.DB
	Rdb    *redis.Client
)
