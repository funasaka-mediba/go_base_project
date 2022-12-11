package customError

import (
	"fmt"
	"go_base_project/constant"
	"runtime"
)

// CustomError 独自エラー(想定できるエラーケースをあらかじめ定義しておくことで、それに沿った実装をフロント側で行える)
type CustomError struct {
	Err       error  // 標準のエラー
	ErrSource string // メッセージ
	Code      constant.ErrorCode
	RefURL    string
	Level     Level
	Status    int
}

// Level ログ出力する際のエラーレベルを定義
type Level string

const (
	Warn  Level = "warn"
	Error Level = "error"
	Panic Level = "panic"
	Fatal Level = "fatal"
)

// NewErr コンストラクタ
func NewErr(e error, code constant.ErrorCode, level Level, status int, refURL string) *CustomError {
	// この一文の意味は？
	_, file, line, _ := runtime.Caller(1)
	err := &CustomError{
		Err:       e,
		ErrSource: file + ":" + fmt.Sprint(line),
		Code:      code,
		RefURL:    refURL,
		Level:     level,
		Status:    status,
	}
	return err
}

// GetError err.GetError()のように使用する。なくてもいいけど、そうするとerr.Errみたいに使用することになり、ちょっとわかりづらくなるので、Getterを用意しておく。
func (e *CustomError) GetError() error {
	return e.Err
}
