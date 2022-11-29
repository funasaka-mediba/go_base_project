package main

import (
	"fmt"
	"go_base_project/packages/env"
	"go_base_project/packages/log"
	"go_base_project/registry"
	"net/http"
	"os"
)

var version string

func main() {
	// setupEnv()
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
