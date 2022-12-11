package v1

import (
	"fmt"
	"go_base_project/constant"
	"time"

	"github.com/gin-gonic/gin"
)

type AppHandler interface {
	HogeHandler
}

func CreateErrorResponse(requestId string, err error, errCode constant.ErrorCode, refURL string) gin.H {
	return gin.H{
		// "request_id": requestId, リクエストヘッダーにrequestIDとかがあるなら必要
		"timestamp": time.Now().Format(constant.ZapDateTimelayout),
		"error": gin.H{
			"message": fmt.Sprintf("%v", err),
			"code":    errCode,
			"ref_url": refURL,
		},
	}
}
