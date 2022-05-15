package resp

import (
	"gin-template/lib/errx"
	"github.com/gin-gonic/gin"
)

func FormatRespSelfMsg(httpState int, errorCode int, msg string, returnData interface{}) (int, gin.H) {
	return httpState, gin.H{
		"data": returnData,
		"msg":  msg,
		"code": errorCode,
	}
}

func FormatRespAutoMsg(httpState int, errorCode int, returnData interface{}) (int, gin.H) {
	return httpState, gin.H{
		"data": returnData,
		"msg":  errx.GetErrMsg(errorCode),
		"code": errorCode,
	}
}
