package controller

import (
    "github.com/gin-gonic/gin"
    "gin-example/service"
    "gin-example/util/e"
)

type Auth struct {
}
func NewAuth() Auth {
    return Auth{};
}
func (this *Auth) GetToken(c *gin.Context) {
    AuthS := service.Auth{}
    token, err := AuthS.GetToken()
    if err != nil {
        e.FailResponse(c, 999, "错误了")
        return
    }
    e.SuccessResponse(c, map[string]string{"token":token})

}

