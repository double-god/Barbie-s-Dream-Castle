// 统一响应，存放可复用的工具函数。JWT 生成/解析、统一响应 `Respond`、密码加密等。
package util

import (
	"net/http"
	"practice_usermanagement/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// Respond 统一响应格式
func Respond(c *gin.Context, httpStatus int, appCode int, data interface{}) {
	c.JSON(httpStatus, Response{
		Code: appCode,
		Msg:  e.GetMsg(appCode),
		Data: data,
	})
}
func RespondSuccess(c *gin.Context, data interface{}) {
	Respond(c, http.StatusOK, e.SUCCESS, data)
}
func RespondError(c *gin.Context, httpCode, appCode int) {
	Respond(c, httpCode, appCode, map[string]interface{}{})
}
