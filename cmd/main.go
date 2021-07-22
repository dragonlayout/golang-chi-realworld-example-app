package main

import (
	"flag"
	"fmt"
	"github.com/dragonlayout/golang-chi-realworld-example-app/assembly"
	"github.com/dragonlayout/golang-chi-realworld-example-app/config"
	"github.com/gin-gonic/gin"
	"os"
)

var env = flag.String("env", "development", "enviroment to use\nValues: \"development\", \"production\"")
var configFile = flag.String("config-file", "", "path of TOML format config file")

func main() {
	flag.Parse()

	var conf *config.Config

	// TODO: 1. 配置文件加载
	{
		if *configFile == "" {
			conf = config.Parse("./config/config-example.toml")
		} else {
			conf = config.Parse(*configFile)
		}
		// TODO: 判断启动参数
		if *env != "production" && *env != "development" {
			os.Exit(1)
		}
		if _, ok := conf.ServerEnv[*env]; !ok {
			fmt.Fprintf(gin.DefaultWriter, "config file invalid! \"%s\" enviroment not found", *env)
			os.Exit(1)
		}
	}

	// TODO: 2. 组件装配
	route := assembly.Build(conf, *env)

	// TODO: 3. 启动 http server
	route.Run(fmt.Sprintf("%s:%d", conf.ServerEnv[*env].ListenAddress, conf.ServerEnv[*env].ListenPort))

}
