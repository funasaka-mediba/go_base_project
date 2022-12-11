package v1

import "github.com/gin-gonic/gin"

type HogeHandler interface {
	Hoge(ctx *gin.Context)
}

type hogeHandler struct {
}

func NewHogeHandler() HogeHandler {
	return hogeHandler{}
}

func (h hogeHandler) Hoge(ctx *gin.Context) {

}
