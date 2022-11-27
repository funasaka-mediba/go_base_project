package main

import (
	"fmt"
	"go_base_project/packages/env"
	"go_base_project/registry"
	"net/http"
)

func main() {
	// setupEnv()
	// setupLogger()
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
