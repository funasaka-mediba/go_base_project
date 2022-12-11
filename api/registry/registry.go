package registry

import (
	"go_base_project/interfaces/routers"
	"go_base_project/registry/container"

	"github.com/gin-gonic/gin"
)

type Registry interface {
	Registry() routers.Engine
}

type registry struct {
	engin     routers.Engine
	container container.Container
}

func NewRegistry() Registry {
	return &registry{
		engin: routers.Engine{
			// gin.Default()使うのがginの使い方でよく書かれているけど、
			// デフォルトでLoggerとRecoveryのミドルウェアを使ってるから、zapとか使いたい時はNew()にする。
			Engine: gin.New(),
		},
		container: container.Container{},
	}
}

func (r registry) Registry() routers.Engine {
	// handlerごとに依存レイヤーを注入
	r.engin.SetBase()
	r.engin.SetCORS()
	// r.engin.SetRouter(r.container.GetAppHandler())
	return r.engin
}
