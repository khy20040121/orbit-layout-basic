package handler

// 基础 Handler：承载 HTTP 层共享依赖
// 说明：logger 是在程序启动时构造的同一个实例，通过 Wire 注入到所有 Handler/Service/Repository。
// 并不是“每个文件夹一个 logger”。如果需要按业务打标签，可在具体 Handler 中基于同一实例派生子 logger（例如 logger.With(...)).
import (
	"github.com/khy20040121/orbit-layout-basic/pkg/log"
)

type Handler struct {
	// 所有业务 Handler 复用的日志器
	logger *log.Logger
}

// NewHandler 由依赖注入传入统一的 logger 实例
func NewHandler(logger *log.Logger) *Handler {
	return &Handler{
		logger: logger,
	}
}
