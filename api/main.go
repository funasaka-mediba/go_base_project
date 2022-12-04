package main

import (
	"fmt"
	"go_base_project/packages/env"
	"go_base_project/packages/log"
	"go_base_project/registry"
	"net/http"
	"os"

	"github.com/joho/godotenv"
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

	// 一旦エラー逃れ
	fmt.Println(server)

	// TODO: graceful restart&stop
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
