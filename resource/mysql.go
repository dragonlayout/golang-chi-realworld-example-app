package resource

import (
	"database/sql"
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/config"
	"github.com/gin-gonic/gin"
)

func InitMysqlPool(conf *config.MysqlConfig) (dbConn *sql.DB, err error) {
	dbConn, err = sql.Open("mysql", concatDSN(conf.Address, conf.Protocol, conf.User, conf.Password, conf.Port, conf.Database))
	if err != nil {
		_, _ = fmt.Fprintf(gin.DefaultWriter, "database connect fail: %v\n", err)
		return
	}
	dbConn.SetConnMaxLifetime()
	dbConn.SetConnMaxIdleTime()
	dbConn.SetMaxIdleConns()
	dbConn.SetMaxOpenConns()

	if err = dbConn.Ping(); err != nil {
		_, _ = fmt.Fprintf(gin.DefaultWriter, "database ping fail: %v\n", err)
		return
	}
	return
}

func concatDSN(address string, protocol string, user string, password string, port int, database string) string {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", address, protocol, user, password, port, database)
	return dsn
}