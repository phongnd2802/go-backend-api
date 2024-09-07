package initializes

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/phongnd2802/go-backend-api/global"
	"go.uber.org/zap"
	"time"
)

func checkErrorPanic(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func initMysql() {
	m := global.Config.Mysql

	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.User, m.Pass, m.Host, m.Port, m.DB)
	db, err := sql.Open("mysql", s)
	checkErrorPanic(err, "sql.Open failed")
	global.Logger.Info("sql.Open success")
	global.Mdb = db

	setPool()
}

func setPool() {
	m := global.Config.Mysql

	global.Mdb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns) * time.Second)
	global.Mdb.SetMaxOpenConns(m.MaxOpenConns)
	global.Mdb.SetConnMaxLifetime(time.Duration(m.ConnMaxLife) * time.Second)
}
