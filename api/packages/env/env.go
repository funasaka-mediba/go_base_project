package env

var cfg Config

type Config struct {
	GoEnv      string `envconfig:"GO_ENV" default:"prd"`
	ListenHost string `envconfig:"LISTEN_HOST" default:"prd"`
	ListenPort string `envconfig:"LISTEN_PORT" default:"prd"`
}

func Env() Config {
	return cfg
}
