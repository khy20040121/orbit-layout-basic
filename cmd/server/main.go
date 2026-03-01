package main

import (
	"fmt"

	"github.com/khy20040121/orbit-layout-basic/cmd/server/wire"
	"github.com/khy20040121/orbit-layout-basic/pkg/config"
	"github.com/khy20040121/orbit-layout-basic/pkg/http"
	"github.com/khy20040121/orbit-layout-basic/pkg/log"
	"go.uber.org/zap"
)

func main() {
	// 加载配置：默认读取 config/local.yml。
	// 可通过 APP_CONF 或 -conf 指定，比如：
	//   APP_CONF=config/prod.yml go run ./cmd/server
	conf := config.NewConfig()

	// 初始化日志：日志级别/格式/输出由配置（log.*）控制
	logger := log.NewLog(conf)

	// 打印启动地址：端口来自配置 http.port
	logger.Info("server start", zap.String("host", "http://127.0.0.1:"+conf.GetString("http.port")))

	// 依赖注入，构建应用
	app, cleanup, err := wire.NewWire(conf, logger)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// 启动 HTTP 服务
	http.Run(app, fmt.Sprintf(":%d", conf.GetInt("http.port")))
}
