package assembly

import (
	"database/sql"
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/config"
	"github.com/dragonlayout/golang-chi-realworld-example-app/route"
	"github.com/dragonlayout/golang-chi-realworld-example-app/service"
	"github.com/gin-gonic/gin"
)

// Build 日志, 数据库等中间件, 业务服务, handler, 路由 初始化和依赖注入
// 返回值 一般为 web 组件, 由 main.go 来调用运行
func Build(conf *config.Config) *gin.Engine {

	var db *sql.DB
	// TODO: 数据库初始化
	fmt.Println(conf)

	// service
	{
		_ = service.New(db)
	}

	// handler
	// TODO: handler 装载

	r := route.New()
	return r
}
