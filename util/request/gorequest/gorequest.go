package gorequest

import (
    "fmt"
    "github.com/parnurzeal/gorequest"
)

// post请求
func PostJson(url string, params *map[string]string) (gorequest.Response, []byte, []error) {
    request := gorequest.New()
    resp, body, errs := request.Post(url).
        Set("Content-Type", "application/json").
        Send(params).
        EndBytes()
    fmt.Println("resp", resp)
    fmt.Println("body", body)
    fmt.Println("errs", errs)
    return resp, body, errs

}

// postForm请求
func PostForm(url string, params *map[string]string) (gorequest.Response, []byte, []error)  {
    request := gorequest.New()

    resp, body, errs := request.Post(url).
        Set("Content-Type", "application/x-www-form-urlencoded").
        Send(params).
        EndBytes()
    fmt.Println("resp", resp)
    fmt.Println("body", body)
    fmt.Println("errs", errs)
    return resp, body, errs

}
func Get(url string, params *map[string]string) (gorequest.Response, []byte, []error)  {
    request := gorequest.New()
    resp, body, errs := request.Get(url).
        Send(params).
        EndBytes()
    fmt.Println("resp", resp)
    fmt.Println("body", body)
    fmt.Println("errs", errs)
    return resp, body, errs
}
