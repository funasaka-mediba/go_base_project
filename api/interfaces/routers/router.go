package routers

import (
	v1 "go_base_project/interfaces/handler/v1"
	"go_base_project/interfaces/middleware"
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
		"/hoges",
		v1.GetHoges,
	)
	apiv1.GET(
		"/hoge",
		middleware.BindGetHogeRequestHeader(),
		v1.GetHoge,
	)
	apiv1.POST(
		"/hoge",
		middleware.BindPostHogeRequestHeader(),
		v1.PostHoge,
	)
	apiv1.DELETE(
		"/hoge",
		middleware.BindDeleteHogeRequestHeader(),
		v1.DeleteHoge,
	)
}

// SetCORS cors information
func (e *Engine) SetCORS() {
	e.Engine.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			env.Env().AccessAllowOrigin,
			// env.Env().AccessAllowOriginWeb,
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"DELETE",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"*",
		},
		// cookieなどの情報を必要とするかどうか
		// AllowCredentials: true,
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	}))
}

// HealthCheck .
func HealthCheck(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"health_check": "OK",
	})
}
