package env

import "github.com/kelseyhightower/envconfig"

var cfg Config

// envconfigパッケージは環境変数を読みこんで必須縛りをつけることもできるし、
// 環境によって読む変数を変えてくれたり、変数のチェックを行えたりすることが利点。
// 標準パッケージだけでやるならosとかflagを使うことになる。
type Config struct {
	GoEnv                string `envconfig:"GO_ENV" default:"prd"`
	ListenHost           string `envconfig:"LISTEN_HOST" default:"prd"`
	ListenPort           string `envconfig:"LISTEN_PORT" default:"prd"`
	AccessAllowOrigin    string `envconfig:"ACCESS_ALLOW_ORIGIN" default:"prd"`
	AccessAllowOriginWeb string `envconfig:"ACCESS_ALLOW_ORIGIN_WEB" default:"prd"`
	HostURL              string `envconfig:"HOST_URL" default:"prd"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`
}

func SetupEnv() error {
	return envconfig.Process("", &cfg)
}

func Env() Config {
	return cfg
}
