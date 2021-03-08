package config

import (
	"appcheck/internal/util"
	"fmt"

	"github.com/BurntSushi/toml"
)

var globalConfig Config

type Config struct {
	// 数据库配置
	Mysql struct {
		Server   string `toml:"server"`
		Port     int    `toml:"port"`
		User     string `toml:"user"`
		Password string `toml:"password"`
		Name     string `toml:"name"`
	}
	// redis配置
	Redis struct {
		Server string `toml:"server"`
		Port   int    `toml:"port"`
	}
	// http配置
	Http struct {
		Port  int    `toml:"port"`
		Host  string `toml:"host"`
		Debug bool   `toml:"debug"`
	}
}

// 通过文件加载配置
func Load(fileName string) error {

	if util.FileExist(fileName) {
		_, err := toml.DecodeFile(fileName, &globalConfig)
		return err
	}

	return nil
}

// 获取数据库源
func GetMysqlDSN() string {

	// mysql dsn
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		globalConfig.Mysql.User, globalConfig.Mysql.Password,
		globalConfig.Mysql.Server, globalConfig.Mysql.Port, globalConfig.Mysql.Name)

	return dsn
}

// 获取redis地址
func GetRedisAddr() string {
	return fmt.Sprintf("%s:%d", globalConfig.Redis.Server, globalConfig.Redis.Port)
}

// 获取Http地址
func GetHttpAddress() string {
	var (
		host = globalConfig.Http.Host
		port = globalConfig.Http.Port
	)
	if host == "" {
		host = "0.0.0.0"
	}
	if port <= 0 {
		port = 8080
	}
	return fmt.Sprintf("%s:%d", host, port)
}

func EnableHttpDebug() bool {
	return globalConfig.Http.Debug
}
