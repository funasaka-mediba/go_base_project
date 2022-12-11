package routers

import (
	v1 "go_base_project/interfaces/handler/v1"
	"go_base_project/packages/env"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Engin *gin.Engine

type Engine struct {
	Engine *gin.Engine
}

// SetBase Init initialize gin
func (e *Engine) SetBase() {
	e.Engine.Use(gin.Recovery())

	e.Engine.GET("v1/status", HealthCheck)
}

func (e *Engine) SetRouter(v1 v1.AppHandler) {
	// v1
	apiv1 := e.Engine.Group("/v1")

	apiv1.GET(
		"/hoge",
		v1.Hoge,
	)
}

// SetCORS cors information
func (e *Engine) SetCORS() {
	e.Engine.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			env.Env().AccessAllowOrigin,
			env.Env().AccessAllowOriginWeb,
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"OPTIONS",
		},
		AllowHeaders: []string{
			"*",
		},
		MaxAge: 24 * time.Hour, // プリフライトのキャッシュ時間
	}))
}

// HealthCheck .
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"health_check": "OK",
	})
}
