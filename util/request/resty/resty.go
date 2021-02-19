package resty

import (
    "fmt"
    "github.com/go-resty/resty/v2"
)


func Get(url string, params map[string]string)   {
    client := resty.New()
    resp, err := client.R().
        SetQueryParams(params).
        Get(url)
    fmt.Println("resp", resp)
    fmt.Println("errs", err)
}
// post请求
func PostJson(url string, params map[string]string)  {
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/json").
        SetBody(params).
        Post(url)
    fmt.Println("resp", resp)
    fmt.Println("errs", err)

}

// postForm请求
func PostForm(url string, params map[string]string) {
    client := resty.New()
    resp, err := client.R().
        SetHeader("Content-Type", "application/x-www-form-urlencoded").
        SetBody(params).
        Post(url)
    fmt.Println("resp", resp)
    fmt.Println("errs", err)

}