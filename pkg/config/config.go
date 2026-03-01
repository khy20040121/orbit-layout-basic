package config

// Package config 负责加载配置文件（位于项目 config/ 目录）
// 支持通过 APP_CONF 环境变量或 -conf 参数指定，默认使用 config/local.yml。
// 常见环境文件：
// - config/local.yml：本地开发环境
// - config/prod.yml：生产环境

import (
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

// NewConfig 加载配置文件并返回 viper 实例
func NewConfig() *viper.Viper {
	// 优先从环境变量读取配置路径：APP_CONF
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		// 未设置 APP_CONF 时，支持命令行参数 -conf 指定配置文件
		flag.StringVar(&envConf, "conf", "config/local.yml", "config path, eg: -conf config/local.yml")
		flag.Parse()
	}
	if envConf == "" {
		// 兜底默认值：使用本地开发配置
		envConf = "config/local.yml"
	}
	// 打印加载的配置文件路径，便于排查
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)
}

// getConfig 使用 viper 读取指定的 yml 配置文件
func getConfig(path string) *viper.Viper {
	conf := viper.New()
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return conf
}
