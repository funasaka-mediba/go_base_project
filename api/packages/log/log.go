package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger
var defaultLogLevel = "info"

// 環境に合わせて環境変数やログレベルを変更できるようなConfig(zapでは本来なくてもよい)
type Config struct {
	Env        string
	LogLevel   string
	AppName    string
	AppVersion string
}

// 参考 https://qiita.com/emonuh/items/28dbee9bf2fe51d28153
// ここってなんでdefaultZapLoggerConfig := zap.Configという書き方じゃダメなんだろう？ -> 変数の宣言はいいけど、:= は関数の中じゃないとだめよ。
var defaultZapLoggerConfig = zap.Config{
	Level:       zap.NewAtomicLevelAt(ConvertToZapLevel(defaultLogLevel)),
	Development: false, // true: DEVモード、false: PRDモード
	// サンプリングの設定を定義する。同じログレベルかつ同じメッセージのログを1秒間に何回出すか制限できる。
	Sampling: &zap.SamplingConfig{
		Initial:    100, // 最初のタイミングで100回まで同時に出力できる
		Thereafter: 100, // それ以降のタイミングで100回まで同時に出力できる
	},
	Encoding: "json", // console or json
	EncoderConfig: zapcore.EncoderConfig{ // jsonの中身にどう表示するかの設定
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	},
	OutputPaths:      []string{"stderr"}, // ログの出力先
	ErrorOutputPaths: []string{"stderr"}, // zap内部エラー出力先
}

// 本来なら、NewLogger関数がなくてもdefaultZapLoggerConfig.Build()でいいけど、
// Envとか入れたり、LogLevelを変更できるようにするために、以下のようなコンストラクタを用意してる
func NewLogger(cfg *Config) (*zap.Logger, error) {
	zcfg := defaultZapLoggerConfig
	zcfg.Level = zap.NewAtomicLevelAt(ConvertToZapLevel(cfg.LogLevel))
	l, err := zcfg.Build() // 引数にoptionをとれるが割愛
	if err != nil {
		return nil, err
	}
	return l.With(
		zap.String("env", cfg.Env),
		zap.String("app_name", cfg.AppName),
		zap.String("app_version", cfg.AppVersion),
	), nil
}

func SetLogger(l *zap.Logger) {
	Logger = l
}

func ConvertToZapLevel(lvl string) zapcore.Level {
	switch lvl {
	case "debug":
		return zap.DebugLevel
	case "info":
		return zap.InfoLevel
	case "warn":
		return zap.WarnLevel
	case "error":
		return zap.ErrorLevel
	case "dpanic":
		return zap.DPanicLevel
	case "panic":
		return zap.PanicLevel
	case "fatal":
		return zap.FatalLevel
	}
	return zap.InfoLevel
}
