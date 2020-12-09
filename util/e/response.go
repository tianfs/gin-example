package e

import (
    "github.com/gin-gonic/gin"
)

type Response struct {
    Code    int         `json:"code"`
    Message string      `json:"message"`
    Data    interface{} `json:"data"`
}
//成功返回
func SuccessResponse(c *gin.Context, data interface{}) {
    c.JSON(SUCCESS, Response{
        Code:    SUCCESS_CODE,
        Message: GetMsg(SUCCESS_CODE),
        Data:    data,
    })
}
//非常严重的失败 返回500错误
func ErrorResponse(c *gin.Context, code int, message ...string) {
    currMessage := GetMsg(code)
    if len(message) >= 1 {
        currMessage = message[0]
    }
    c.JSON(ERROR, Response{
        Code:    code,
        Message: currMessage,
        Data:    map[string]interface{}{},
    })
}
//http请求成功，但是程序结果是失败
func FailResponse(c *gin.Context, code int, message ...string){
    currMessage := GetMsg(code)
    if len(message) >= 1 {
        currMessage = message[0]
    }
    c.JSON(SUCCESS, Response{
        Code:    code,
        Message: currMessage,
        Data:    map[string]interface{}{},
    })
}