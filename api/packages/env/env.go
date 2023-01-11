package env

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

var cfg Config

// envconfigパッケージは環境変数を読みこんで必須縛りをつけることもできるし、
// 環境によって読む変数を変えてくれたり、変数のチェックを行えたりすることが利点。
// 標準パッケージだけでやるならosとかflagを使うことになる。
type Config struct {
	GoEnv             string `envconfig:"GO_ENV" default:"prd"`
	ListenHost        string `envconfig:"LISTEN_HOST" default:"prd"`
	ListenPort        string `envconfig:"LISTEN_PORT" default:"prd"`
	AccessAllowOrigin string `envconfig:"ACCESS_ALLOW_ORIGIN" default:"prd"`
	// AccessAllowOriginWeb string `envconfig:"ACCESS_ALLOW_ORIGIN_WEB" default:"prd"`
	HostURL string `envconfig:"HOST_URL" default:"prd"`

	LogLevel string `envconfig:"LOG_LEVEL" default:"info"`

	// MYSQL
	DbMysqlReadHost     string        `envconfig:"DB_HOST_READER"`
	DbMysqlWriteHost    string        `envconfig:"DB_HOST_WRITER"`
	DbMysqlPort         string        `envconfig:"DB_PORT"`
	DbMysqlUser         string        `envconfig:"DB_USERNAME"`
	DbMysqlPassword     string        `envconfig:"DB_PASSWORD"`
	DbMysqlDatabase     string        `envconfig:"DB_DATABASE"`
	DbMysqlMaxIdleConns int           `envconfig:"DB_MYSQL_MAX_IDLE_CONNS"`
	DbMysqlMaxOpenConns int           `envconfig:"DB_MYSQL_MAX_OPEN_CONNS"`
	DbMysqlMaxLifetime  time.Duration `envconfig:"DB_MYSQL_MAX_LIFETIME"`
}

func SetupEnv() error {
	return envconfig.Process("", &cfg)
}

func Env() Config {
	return cfg
}
