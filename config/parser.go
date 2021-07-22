package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pelletier/go-toml"
	"os"
)

type Config struct {
	ServerEnv map[string]*ServerConfig `toml:"server"`
	MysqlEnv  map[string]*MysqlConfig  `toml:"mysql"`
}

type ServerConfig struct {
	ListenAddress string `toml:"listen-address"`
	ListenPort    int    `toml:"listen-port"`
}

type MysqlConfig struct {
	Address  string `toml:"address"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	Protocol string `toml:"protocol"`
	Database string `toml:"database"`
}

func Parse(filepath string) *Config {
	var fs []byte
	var err error
	if fs, err = os.ReadFile(filepath); err != nil {
		_, _ = fmt.Fprintf(gin.DefaultWriter, "config file with TOML format not found: %v\n", err)
		os.Exit(1)
	}
	//if err != nil {
	//	_, _ = fmt.Fprintf(gin.DefaultWriter, "config file with TOML format not found: %v\n", err)
	//	os.Exit(1)
	//}
	config := new(Config)
	if err = toml.Unmarshal(fs, config); err != nil {
		_, _ = fmt.Fprintf(gin.DefaultWriter, "config file with TOML format parse fail: %v\n", err)
		os.Exit(1)
	}
	return config
}
