package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	ErrCode int         `json:"errcode"`
	ErrMsg  interface{} `json:"errmsg"`
	Data    interface{} `json:"data"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		ErrCode: 0,
		ErrMsg:  "ok",
		Data:    data,
	})
}

func Fail(c *gin.Context, errorCode int, msg interface{}, data interface{}) {
	c.Abort()
	c.JSON(http.StatusOK, Response{
		ErrCode: errorCode,
		ErrMsg:  msg,
		Data:    data,
	})
}

func App(c *gin.Context, statusCode int, msg string) {
	c.JSON(statusCode, Response{
		ErrCode: statusCode,
		ErrMsg:  msg,
		Data:    nil,
	})
}
