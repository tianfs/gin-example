package service

import (
    "gin-example/util/jwt"
)

func init() {

}

type Auth struct {
    UserName string
    PassWord string
}

func (this *Auth) GetToken() (string, error) {
    token, err := jwt.GenerateToken("username", "password")
    if err != nil {
        return "", err
    }
    return token, nil
}
