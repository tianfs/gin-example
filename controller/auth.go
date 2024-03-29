package controller

import (
    "github.com/gin-gonic/gin"
    "gin-example/service"
    "gin-example/util/e"
)

func  GetToken(c *gin.Context) {
    AuthS := service.Auth{}
    token, err := AuthS.GetToken()
    if err != nil {
        e.FailResponse(c, 999, "错误了")
        return
    }
    e.SuccessResponse(c, map[string]string{"token":token})

}

