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
			Engine: gin.New(),
		},
		container: container.Container{},
	}
}

func (r registry) Registry() routers.Engine {
	// handlerごとに依存レイヤーを注入
	return r.engin
}
