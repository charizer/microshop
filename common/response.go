package common

import (
	"github.com/gin-gonic/gin"
	"microshop/infrastructure/logger"
)

type HTTPData struct {
	ErrNo  int         `json:"errno"`
	ErrMsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

func httpSuccess(val interface{}) HTTPData {
	return HTTPData{
		ErrNo:  0,
		ErrMsg: "",
		Data:   val,
	}
}

func HandleError(c *gin.Context, httpCode, errorCode int, logMsg string) {
	logger.GetLogger().Errorf("httpCode:%d, errorCode:%d error:%s", httpCode, errorCode, logMsg)
	c.JSON(httpCode, ErrorResponse(ErrorCode(errorCode), ErrorText(errorCode), logMsg))
}

func HandleSucc(c *gin.Context, httpCode int, logMsg string, data interface{}) {
	if logMsg != "" {
		logger.GetLogger().Infoln(logMsg)
	}
	c.JSON(httpCode, httpSuccess(data))
}
