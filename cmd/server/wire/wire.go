//go:build wireinject
// +build wireinject

package wire

// 使用 Google Wire 定义依赖注入 Provider 集合与注入器
// 本文件受 build tag `wireinject` 控制：仅在执行 wire 代码生成时参与编译；
// 实际运行时使用 wire 生成的 wire_gen.go（build tag `!wireinject`）。
// Provider 集合说明：
// - ServerSet：HTTP 服务器（路由/中间件等）
// - RepositorySet：数据访问层（数据库/缓存等）
// - ServiceSet：业务服务层
// - HandlerSet：路由处理器

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/khy20040121/orbit-layout-basic/internal/handler"
	"github.com/khy20040121/orbit-layout-basic/internal/repository"
	"github.com/khy20040121/orbit-layout-basic/internal/server"
	"github.com/khy20040121/orbit-layout-basic/internal/service"
	"github.com/khy20040121/orbit-layout-basic/pkg/log"
	"github.com/spf13/viper"
)

var ServerSet = wire.NewSet(server.NewServerHTTP) // HTTP Server Provider 集合

var RepositorySet = wire.NewSet( // 仓储层 Provider 集合
	repository.NewDb,
	repository.NewRepository,
	repository.NewUserRepository,
)

var ServiceSet = wire.NewSet( // 服务层 Provider 集合
	service.NewService,
	service.NewUserService,
)

var HandlerSet = wire.NewSet( // 处理器层 Provider 集合
	handler.NewHandler,
	handler.NewUserHandler,
)

// NewWire 是 Wire 的注入器声明：根据上面的 Provider 集合构建 *gin.Engine。
// 仅在 wire 生成阶段（build tag: wireinject）使用，实际运行调用 wire_gen.go 中生成的同名函数。
func NewWire(*viper.Viper, *log.Logger) (*gin.Engine, func(), error) {
	panic(wire.Build(
		ServerSet,     // HTTP Server
		RepositorySet, // 仓储层
		ServiceSet,    // 服务层
		HandlerSet,    // 处理器
	))
}
