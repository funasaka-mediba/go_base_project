package test

import (
	"go_base_project/packages/env"
	"go_base_project/packages/log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func InitTest(path string) {
	setupEnv(path)
	os.Chdir(strings.Replace(path, "/config/test.env", "", 1))
	setupLogger()
}

func readEnv(path string) {
	err := godotenv.Load(path)
	if err != nil {
		panic(err)
	}
}

func setupEnv(path string) {
	readEnv(path)
	err := env.SetupEnv()
	if err != nil {
		panic(err)
	}
}

func setupLogger() {
	lg, err := log.NewLogger(&log.Config{
		Env:        os.Getenv("APPLICATION_ENV"),
		LogLevel:   os.Getenv("LOG_LEVEL"),
		AppName:    os.Getenv("APP_NAME"),
		AppVersion: "",
	})
	if err != nil {
		panic(err)
	}
	log.SetLogger(lg)
}
