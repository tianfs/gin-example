package test

import (
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

type base1 struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    []testD `json:"data"`
}

var Errorss = errors.New("错误了a:=123")

func TestError(t *testing.T) {
	aaaa := fmt.Errorf("dsadsad %v \n %w \n %T", Errorss, Errorss, Errorss)
	fmt.Println(aaaa)
	newbase := base1{
		Code:    123,
		Message: "你还",
	}
	fmt.Printf("结构体 %v \n %+v", newbase, newbase)
}
