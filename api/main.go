package main

import (
	"context"
	"fmt"
	"go_base_project/packages/env"
	"go_base_project/packages/log"
	"go_base_project/registry"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

var version string

func main() {
	setupEnv()
	setupLogger()
	addr := fmt.Sprintf("%s:%s", env.Env().ListenHost, env.Env().ListenPort)

	registry := registry.NewRegistry()
	router := registry.Registry().Engine
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	// 以下Ginでgraceful restart&stopをする際の定型文(こちらでかいてもいいよと公式ドキュメントに載ってる -> https://github.com/gin-gonic/examples/blob/master/graceful-shutdown/graceful-shutdown/notify-with-context/server.go)
	// ちなみに、ginを使用した以下の実装と標準パッケージでの実装の比較はこちらがわかりやすかった -> https://sourjp.github.io/posts/go-gin-graceful/
	// ginを使用すると、handler周りの記述が少し楽でいいですね。

	/*
		補足
		graceful restart&stopって普通のrestart&stopとどう違うのか
		普通のgracefulなしの方は、システム停止のシグナルを受け取ったらプロセスが稼働中でも即座に停止するため、ユーザーの処理が途中で中断することもある。
		それに対し、gracefulありの方は、ゆるやかな停止・再起動と言うイメージ。
		今稼働中のプロセスがあったらそれが完了するまでは待機して、もしあらかじめ決めた時間以上に時間が経ったら、プロセスをkillして停止するという挙動になる。
		なので、本番環境などでユーザーが利用するようなサービスに関してはgraceful restart&stopで実装することがおすすめ。
	*/

	go func() {
		log.Logger.Info("server started", zap.String("listener", addr))                // これはただのログなので任意。
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed { // 大事なのはserver.ListenAndServe()
			log.Logger.Fatal("failed to serve server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)                                    // channel作成
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT) // control+C = SIGINT がわかりやすいけど、そう言った停止シグナルを受け取る。
	sg := <-quit                                                       // channelからsignalを受け取る。
	log.Logger.Info("gracefully stopping server...", zap.String("signal", sg.String()))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // 10秒は待ってくれる

	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Logger.Fatal("failed to stop server", zap.Error(err))
	}
	log.Logger.Info("stopped server")
}

// どの環境を利用するかの設定はdocker-compose.ymlやecs-task-def.jsonみたいに外から与えたいので以下の関数を用意する
// (その方がデプロイやビルド時に環境を決めれて便利だから)
func readEnv() {
	if err := godotenv.Load(fmt.Sprintf("./config/%s.env", os.Getenv("GO_ENV"))); err != nil {
		panic(err)
	}
}

func setupEnv() {
	readEnv()
	if err := env.SetupEnv(); err != nil {
		panic(err)
	}
}

func setupLogger() {
	lg, err := log.NewLogger(&log.Config{
		Env:        os.Getenv("APPLICATION_ENV"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: version,
	})
	if err != nil {
		panic(err)
	}
	log.SetLogger(lg)
}
