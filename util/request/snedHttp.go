package request

import (
    "encoding/json"
    "errors"
    "gin-example/util/e"
    "io/ioutil"
    "net/http"
    "net/url"
)

type httpResponse struct {
    Code    int                    `json:"errorCode"`
    Message string                 `json:"errorMsg"`
    Data    map[string]interface{} `json:"data"`
}


type httpRequest struct {
    Body *[]byte
}
//get请求
func NewGet(url string, params *map[string]interface{}) (*httpRequest,error) {
    request := new(httpRequest)
    res, error := http.Get(url)

    if error != nil {
        return nil,error
    }
    defer res.Body.Close()
    body, error := ioutil.ReadAll(res.Body)
    if error != nil {
        return nil,error
    }

    if res.StatusCode != e.HTTP_SUCCESS {
        return nil,errors.New("网络请求错误")
    }
    request.Body = &body
    return request,nil
}
func (this *httpRequest) ToJson(dataStruct interface{}) error {
    json.Unmarshal(*this.Body, dataStruct)
    return nil
}
//post请求
func NewPost(url string, params *url.Values) (*httpRequest,error) {
    request := new(httpRequest)
    client := &http.Client{}
    res, error := client.PostForm(url,*params)
    if error != nil {
        return nil,error
    }
    defer res.Body.Close()
    body, error := ioutil.ReadAll(res.Body)
    if error != nil {
        return nil,error
    }

    if res.StatusCode != e.HTTP_SUCCESS {
        return nil,errors.New("网络请求错误")
    }
    request.Body = &body
    return request,nil
}




